package prost

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"protocol-state-cacher/config"
	"protocol-state-cacher/pkgs"
	"protocol-state-cacher/pkgs/redis"
	"protocol-state-cacher/pkgs/reporting"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/cenkalti/backoff"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	log "github.com/sirupsen/logrus"
)

var lastProcessedBlock int64

type EventProcessor struct {
	workerLimit chan struct{}
	wg          sync.WaitGroup
}

func NewEventProcessor(maxWorkers int) *EventProcessor {
	return &EventProcessor{
		workerLimit: make(chan struct{}, maxWorkers),
	}
}

// MonitorEvents continuously monitors blockchain events for updates
func (ep *EventProcessor) MonitorEvents() {
	log.Println("Monitoring blockchain events for updates...")

	ticker := time.NewTicker(time.Duration(float64(time.Second) * config.SettingsObj.BlockInterval))
	defer ticker.Stop()

	for range ticker.C {
		latestBlock, err := fetchBlock(nil)
		if err != nil {
			log.Errorf("Error fetching latest block: %s", err.Error())
			continue
		}

		latestBlockNumber := latestBlock.Number().Int64()
		targetBlockNumber := latestBlockNumber - int64(config.SettingsObj.BlockOffset)

		if targetBlockNumber < 0 {
			log.Warn("Target block number is below zero. Skipping iteration...")
			continue
		}

		if lastProcessedBlock == 0 {
			lastProcessedBlock = targetBlockNumber
		}

		// Process new blocks and backtrack for error correction
		for blockNum := lastProcessedBlock + 1; blockNum <= targetBlockNumber; blockNum++ {
			block, err := fetchBlock(big.NewInt(blockNum))
			if err != nil {
				log.Printf("Error fetching block %d: %v", blockNum, err)
				continue
			}

			if block == nil {
				log.Errorf("Received nil block for number: %d", blockNum)
				continue
			}

			// Limit concurrent goroutines
			ep.workerLimit <- struct{}{}
			ep.wg.Add(1)
			go func(b *types.Block) {
				defer ep.wg.Done()
				defer func() { <-ep.workerLimit }()
				ProcessProtocolStateEvents(b)
			}(block) // Pass block as parameter to closure

			lastProcessedBlock = blockNum
		}
	}
}

// fetchBlock retrieves a block from the client using retry logic
func fetchBlock(blockNum *big.Int) (*types.Block, error) {
	var block *types.Block
	operation := func() error {
		var err error
		block, err = Client.BlockByNumber(context.Background(), blockNum) // Pass blockNum (nil for the latest block)
		if err != nil {
			log.Errorf("Failed to fetch block %v: %v", blockNum, err)
			return err // Return the error to trigger a retry
		}
		return nil // Block successfully fetched, return nil to stop retries
	}

	// Retry fetching the block with a backoff strategy
	if err := backoff.Retry(operation, backoff.WithMaxRetries(backoff.NewConstantBackOff(200*time.Millisecond), 3)); err != nil {
		errMsg := fmt.Sprintf("Failed to fetch block %v after retries: %s", blockNum, err.Error())
		reporting.SendFailureNotification(pkgs.MonitorEvents, errMsg, time.Now().String(), "High")
		log.Error(errMsg)
		return nil, err
	}

	return block, nil
}

// ProcessProtocolStateEvents processes protocol state events
func ProcessProtocolStateEvents(block *types.Block) {
	var logs []types.Log
	var err error

	hash := block.Hash()
	blockNum := block.Number().Int64()

	// Create a filter query to fetch logs for the block
	filterQuery := ethereum.FilterQuery{
		BlockHash: &hash,
		Addresses: []common.Address{common.HexToAddress(config.SettingsObj.ContractAddress)},
	}

	operation := func() error {
		logs, err = Client.FilterLogs(context.Background(), filterQuery)
		return err
	}

	if err = backoff.Retry(operation, backoff.WithMaxRetries(backoff.NewConstantBackOff(200*time.Millisecond), 3)); err != nil {
		errorMsg := fmt.Sprintf("Error fetching logs for block number %d: %s", blockNum, err.Error())
		reporting.SendFailureNotification(pkgs.ProcessProtocolStateEvents, errorMsg, time.Now().String(), "High")
		log.Error(errorMsg)
		return
	}

	log.Infof("Processing %d logs for block number %d", len(logs), blockNum)

	for _, vLog := range logs {
		// Check the event signature and handle the `EpochReleased` event
		switch vLog.Topics[0].Hex() {
		case ContractABI.Events["EpochReleased"].ID.Hex():
			log.Debugf("EpochReleased event detected in block %d", block.Number().Int64())

			// Parse the `EpochReleased` event from the log
			releasedEvent, err := Instance.ParseEpochReleased(vLog)
			if err != nil {
				errorMsg := fmt.Sprintf("Epoch release parse error for block %d: %s", block.Number().Int64(), err.Error())
				reporting.SendFailureNotification(pkgs.ProcessProtocolStateEvents, errorMsg, time.Now().String(), "High")
				log.Error(errorMsg)
				continue
			}

			// Check if the DataMarketAddress in the event matches any address in the DataMarketAddress array
			if isValidDataMarketAddress(releasedEvent.DataMarketAddress.Hex()) {
				// Get data market address and epoch ID from the event
				dataMarketAddress := releasedEvent.DataMarketAddress
				epochID := releasedEvent.EpochId.String()

				// Persist the current epoch
				currentEpochKey := redis.CurrentEpochID(strings.ToLower(dataMarketAddress.Hex()))
				PersistState(context.Background(), currentEpochKey, epochID)
				log.Infof("Current epoch set for data market %s to %s", strings.ToLower(dataMarketAddress.Hex()), epochID)
			}
		}
	}
}

// Create a package-level SlotManager instance with a reasonable batch size
var slotManager = NewSlotManager(100)

func addSlotInfo(slotID int64) {
	// Fetch the slot info from the contract
	slot, err := SnapshotterStateInstance.NodeInfo(&bind.CallOpts{}, big.NewInt(slotID))
	if err != nil {
		log.Printf("Error fetching slot %d: %v", slotID, err)
		return
	}

	// Check if the slot info is empty
	if slot == (struct {
		SnapshotterAddress common.Address
		NodePrice          *big.Int
		AmountSentOnL1     *big.Int
		MintedOn           *big.Int
		BurnedOn           *big.Int
		LastUpdated        *big.Int
		IsLegacy           bool
		ClaimedTokens      bool
		Active             bool
		IsKyced            bool
	}{}) {
		log.Printf("No data for slot %d", slotID)
		return
	}

	// Marshal the slot info to JSON
	slotMarshalled, err := json.Marshal(slot)
	if err != nil {
		log.Printf("Error marshalling returned information for slot %d: %v", slotID, err)
		return
	}

	// Add slot data to the SlotManager, which will handle persistence
	if slotManager.AddSlot(slotID, string(slotMarshalled)) {
		log.Printf("Batch size reached, slots have been flushed to Redis")
	}

	log.Printf("Added slot %d to batch", slotID)
}

// SlotManager handles the storage and batching of slot information
type SlotManager struct {
	mu        sync.RWMutex
	slots     map[int64]string // slot ID --> slot information data unmarshalled
	batchSize int
}

// NewSlotManager creates a new SlotManager with specified batch size
func NewSlotManager(batchSize int) *SlotManager {
	return &SlotManager{
		slots:     make(map[int64]string),
		batchSize: batchSize,
	}
}

// AddSlot adds a slot and returns true if batch size was reached and flush occurred
func (sm *SlotManager) AddSlot(slotID int64, slotData string) bool {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	sm.slots[slotID] = slotData

	if len(sm.slots) >= sm.batchSize {
		sm.flushSlots()
		return true
	}
	return false
}

// flushSlots persists all current slots to Redis and clears the internal storage
// Note: This method assumes the caller holds the lock
func (sm *SlotManager) flushSlots() {
	slots := make([]any, 0, len(sm.slots))
	for slotID, slotData := range sm.slots {
		slotKey := redis.SlotInfo(strconv.FormatInt(slotID, 10))
		PersistState(context.Background(), slotKey, slotData)
		slots = append(slots, strconv.FormatInt(slotID, 10))
	}
	allSlotsKey := redis.AllSlotInfo()
	redis.RedisClient.SAdd(context.Background(), allSlotsKey, slots...)
	log.Printf("Flushed batch of %d slots to Redis. Range: %v - %v", len(sm.slots), slots[0], slots[len(slots)-1])
	sm.slots = make(map[int64]string)
}

// ForceFlush forces a flush of current slots even if batch size hasn't been reached
func (sm *SlotManager) ForceFlush() {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	if len(sm.slots) > 0 {
		sm.flushSlots()
	}
}

// Cleanup ensures all pending slots are flushed before program termination
func Cleanup() {
	slotManager.ForceFlush()
}
