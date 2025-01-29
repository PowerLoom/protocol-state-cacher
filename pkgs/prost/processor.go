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
	"time"

	"github.com/cenkalti/backoff"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	log "github.com/sirupsen/logrus"
)

var lastProcessedBlock int64

var emptySlotInfo = struct {
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
}{}

// MonitorEvents continuously monitors blockchain events for updates
func MonitorEvents() {
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
				log.Errorf("Error fetching block %d: %v", blockNum, err)
				continue
			}

			if block == nil {
				log.Errorf("Received nil block for number: %d", blockNum)
				continue
			}

			// go ProcessSnapshotterStateEvents(block)
			go ProcessProtocolStateEvents(block)

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

// ProcessSnapshotterStateEvents processes snapshotter state events
func ProcessSnapshotterStateEvents(block *types.Block) {
	var logs []types.Log
	var err error

	hash := block.Hash()
	blockNum := block.Number().Int64()

	// Create a filter query to fetch logs for the block
	filterQuery := ethereum.FilterQuery{
		BlockHash: &hash,
		Addresses: []common.Address{SnapshotterStateAddress},
	}

	operation := func() error {
		logs, err = Client.FilterLogs(context.Background(), filterQuery)
		return err
	}

	if err = backoff.Retry(operation, backoff.WithMaxRetries(backoff.NewConstantBackOff(200*time.Millisecond), 3)); err != nil {
		errorMsg := fmt.Sprintf("Error fetching logs for block number %d: %s", blockNum, err.Error())
		reporting.SendFailureNotification(pkgs.ProcessSnapshotterStateEvents, errorMsg, time.Now().String(), "High")
		log.Error(errorMsg)
		return
	}

	log.Infof("Processing %d logs for block number %d", len(logs), blockNum)

	// Process the logs for the current block
	for _, vLog := range logs {
		// Check the event signature
		switch vLog.Topics[0].Hex() {
		case SnapshotterStateContractABI.Events["allSnapshottersUpdated"].ID.Hex():
			log.Debugf("allSnapshottersUpdated event detected in block %d", blockNum)

			// Parse the `allSnapshottersUpdated` event from the log
			releasedEvent, err := SnapshotterStateInstance.ParseAllSnapshottersUpdated(vLog)
			if err != nil {
				errorMsg := fmt.Sprintf("Failed to parse `allSnapshottersUpdated` event for block %d: %s", blockNum, err.Error())
				reporting.SendFailureNotification(pkgs.ProcessSnapshotterStateEvents, errorMsg, time.Now().String(), "High")
				log.Error(errorMsg)
				continue
			}

			// Fetch the transaction details using the transaction hash
			txHash := releasedEvent.Raw.TxHash
			tx, _, err := Client.TransactionByHash(context.Background(), txHash)
			if err != nil {
				errorMsg := fmt.Sprintf("Failed to fetch transaction details for hash %s in block %d: %s", txHash.Hex(), blockNum, err.Error())
				reporting.SendFailureNotification(pkgs.ProcessSnapshotterStateEvents, errorMsg, time.Now().String(), "High")
				log.Error(errorMsg)
				continue
			}

			// Decode the transaction input to get the node ID and snapshotter address
			nodeID, snapshotterAddress, err := decodeTransactionInput(tx.Data())
			if err != nil {
				errMsg := fmt.Sprintf("Failed to decode transaction input for hash %s in block %d: %s", txHash.Hex(), blockNum, err.Error())
				reporting.SendFailureNotification(pkgs.ProcessSnapshotterStateEvents, errMsg, time.Now().String(), "High")
				log.Error(errMsg)
				continue
			}

			// Process the node ID and snapshotter address
			log.Infof("🚀 Node ID: %d, Snapshotter Address: %s", nodeID, snapshotterAddress.Hex())
			addSlotInfo(nodeID)
		}
	}
}

func addSlotInfo(slotID int64) {
	// Fetch the slot info from the snapshotter state contract
	slotInfo, err := SnapshotterStateInstance.NodeInfo(&bind.CallOpts{}, big.NewInt(slotID))
	if err != nil {
		log.Errorf("Error fetching slot info for slot %d: %v", slotID, err)
		return
	}

	// Check if the slot info is empty
	if slotInfo == emptySlotInfo {
		log.Errorf("No data for slot %d", slotID)
		return
	}

	// Marshal the slot info to JSON
	slotInfoMarshalled, err := json.Marshal(slotInfo)
	if err != nil {
		log.Errorf("Error marshalling slot info for slot %d: %v", slotID, err)
		return
	}

	// Persist slot information
	slotInfoKey := redis.SlotInfo(strconv.FormatInt(slotID, 10))
	PersistState(context.Background(), slotInfoKey, string(slotInfoMarshalled))

	// Add slot key to the batch
	mu.Lock()
	allSlots = append(allSlots, slotInfoKey)
	mu.Unlock()

	log.Printf("Fetched and persisted slot info for slot %d", slotID)
}

func decodeTransactionInput(inputData []byte) (int64, common.Address, error) {
	// Ensure input data is non-empty
	if len(inputData) == 0 {
		return 0, common.Address{}, fmt.Errorf("input data is empty, nothing to decode")
	}

	// Extract the method ID (first 4 bytes)
	methodID := inputData[:4]

	// Find the method by its ID
	method, err := SnapshotterStateContractABI.MethodById(methodID)
	if err != nil {
		return 0, common.Address{}, fmt.Errorf("failed to find method by ID: %v", err)
	}

	if method.Name != "assignSnapshotterToNode" {
		return 0, common.Address{}, fmt.Errorf("unexpected method: %s", method.Name)
	}

	// Decode the parameters of the method
	params := map[string]interface{}{}
	err = method.Inputs.UnpackIntoMap(params, inputData[4:])
	if err != nil {
		return 0, common.Address{}, fmt.Errorf("failed to unpack input data: %w", err)
	}

	nodeID, ok := params["nodeId"].(*big.Int)
	if !ok {
		return 0, common.Address{}, fmt.Errorf("invalid type for nodeId")
	}

	snapshotterAddress, ok := params["snapshotterAddress"].(common.Address)
	if !ok {
		return 0, common.Address{}, fmt.Errorf("invalid type for snapshotterAddress")
	}

	return nodeID.Int64(), snapshotterAddress, nil
}
