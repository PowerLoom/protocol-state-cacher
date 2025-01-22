package prost

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"protocol-state-cacher/config"
	"protocol-state-cacher/pkgs"
	"protocol-state-cacher/pkgs/contract"
	"protocol-state-cacher/pkgs/redis"
	"protocol-state-cacher/pkgs/reporting"
	"protocol-state-cacher/pkgs/snapshotterStateContract"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	log "github.com/sirupsen/logrus"
)

var (
	allSlots                    []string
	mu                          sync.Mutex
	Client                      *ethclient.Client
	Instance                    *contract.Contract
	ContractABI                 abi.ABI
	SnapshotterStateContractABI abi.ABI
	SnapshotterStateInstances   = make(map[string]*snapshotterStateContract.SnapshotterStateContract)
)

func Initialize() {
	// Initialize RPC client
	ConfigureClient()

	// Initialize contract instances
	ConfigureContractInstance()
	ConfigureSnapshotterStateContractInstance()

	// Initialize ABI instances
	ConfigureABI()
	ConfigureSnapshotterStateABI()
}

func ConfigureClient() {
	// Initialize RPC client
	rpcClient, err := rpc.DialOptions(context.Background(), config.SettingsObj.ClientUrl, rpc.WithHTTPClient(&http.Client{Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}}))
	if err != nil {
		log.Errorf("Failed to connect to client: %s", err)
		log.Fatal(err)
	}

	Client = ethclient.NewClient(rpcClient)
}

func ConfigureContractInstance() {
	// Initialize single contract instance using the protocol state contract address
	protocolStateAddress := common.HexToAddress(config.SettingsObj.ContractAddress)
	instance, err := contract.NewContract(protocolStateAddress, Client)
	if err != nil {
		log.Fatalf("Failed to create protocol state contract instance: %v", err)
	}

	Instance = instance
}

func ConfigureSnapshotterStateContractInstance() {
	// Initialize snapshotter state contract instance
	for _, dataMarketContractAddress := range config.SettingsObj.DataMarketContractAddresses {
		snapshotterStateInstance, _ := snapshotterStateContract.NewSnapshotterStateContract(dataMarketContractAddress, Client)
		SnapshotterStateInstances[dataMarketContractAddress.Hex()] = snapshotterStateInstance
	}
}

func ConfigureABI() {
	// Initialize contract ABI
	contractABI, err := abi.JSON(strings.NewReader(contract.ContractMetaData.ABI))
	if err != nil {
		log.Errorf("Failed to configure protocol state contract ABI: %s", err)
		log.Fatal(err)
	}

	ContractABI = contractABI
}

func ConfigureSnapshotterStateABI() {
	// Initialize snapshotter state contract ABI
	snapshotterStateABI, err := abi.JSON(strings.NewReader(snapshotterStateContract.SnapshotterStateContractMetaData.ABI))
	if err != nil {
		log.Errorf("Failed to configure snapshotter state contract ABI: %s", err)
		log.Fatal(err)
	}

	SnapshotterStateContractABI = snapshotterStateABI
}

func MustQuery[K any](ctx context.Context, call func(opts *bind.CallOpts) (val K, err error)) (K, error) {
	expBackOff := backoff.NewExponentialBackOff()
	expBackOff.MaxElapsedTime = 1 * time.Minute
	var val K
	operation := func() error {
		var err error
		val, err = call(&bind.CallOpts{})
		return err
	}
	// Use the retry package to execute the operation with backoff
	err := backoff.Retry(operation, backoff.WithContext(expBackOff, ctx))
	if err != nil {
		return *new(K), err
	}
	return val, err
}

// FetchAllSlots fetches all slot information once during startup.
func FetchAllSlots() error {
	log.Println("Fetching all slot information at startup...")

	// Fetch total node count
	nodeCount, err := Instance.GetTotalNodeCount(&bind.CallOpts{Context: context.Background()})
	if err != nil {
		log.Errorf("Error fetching total node count: %s", err.Error())
		return fmt.Errorf("failed to fetch total node count: %s", err.Error())
	}

	// Loop through all data market addresses in the configuration
	for _, dataMarketAddress := range config.SettingsObj.DataMarketContractAddresses {
		log.Printf("Fetching slots for data market address: %s", dataMarketAddress.Hex())

		// Process slots in batches of 20 for parallelism
		for i := int64(0); i <= nodeCount.Int64(); i += 20 {
			var wg sync.WaitGroup

			// Fetch each slot in the current batch
			for j := i; j < i+20 && j <= nodeCount.Int64(); j++ {
				wg.Add(1)

				go func(slotIndex int64) {
					defer wg.Done()
					addSlotInfo(dataMarketAddress.Hex(), slotIndex)
				}(j)
			}

			wg.Wait()

			// Add batch of slots to Redis
			if len(allSlots) > 0 {
				if err := redis.AddToSet(context.Background(), redis.AllSlotInfo(), allSlots...); err != nil {
					errMsg := fmt.Sprintf("Error adding slots to Redis set for data market %s: %v", dataMarketAddress.Hex(), err)
					reporting.SendFailureNotification(pkgs.FetchAllSlots, errMsg, time.Now().String(), "High")
					log.Error(errMsg)
				}
			}
		}
	}

	log.Println("✅ Completed fetching all slot information at startup.")
	return nil
}

func StartPeriodicStateSync() {
	go func() {
		for {
			// Poll dynamic state variables
			DynamicStateVariables()

			// Poll static variables if PollingStaticStateVariables is true
			if config.SettingsObj.PollingStaticStateVariables {
				StaticStateVariables()
			}

			time.Sleep(time.Second * time.Duration(config.SettingsObj.SlotSyncInterval))
		}
	}()
}

func StaticStateVariables() {
	// Iterate over all data markets and set static state variables
	for _, dataMarketAddress := range config.SettingsObj.DataMarketContractAddresses {
		log.Infof("Setting static state variables for data market: %s", dataMarketAddress)

		// Set epochs in a day
		if output, err := Instance.EpochsInADay(&bind.CallOpts{}, dataMarketAddress); output != nil && err == nil {
			epochsInADayKey := redis.ContractStateVariableWithDataMarket(dataMarketAddress.Hex(), pkgs.EpochsInADay)
			PersistState(context.Background(), epochsInADayKey, output.String())
			log.Infof("Epochs in a day set for data market %s to %s", strings.ToLower(dataMarketAddress.Hex()), output.String())
		}

		// Set epoch size
		if output, err := Instance.EPOCHSIZE(&bind.CallOpts{}, dataMarketAddress); output != 0 && err == nil {
			epochSizeKey := redis.ContractStateVariableWithDataMarket(dataMarketAddress.Hex(), pkgs.EPOCH_SIZE)
			PersistState(context.Background(), epochSizeKey, strconv.Itoa(int(output)))
			log.Infof("Epoch size set for data market %s to %s", strings.ToLower(dataMarketAddress.Hex()), strconv.Itoa(int(output)))
		}

		// Set source chain block time
		if output, err := Instance.SOURCECHAINBLOCKTIME(&bind.CallOpts{}, dataMarketAddress); output != nil && err == nil {
			sourceChainBlockTimeKey := redis.ContractStateVariableWithDataMarket(dataMarketAddress.Hex(), pkgs.SOURCE_CHAIN_BLOCK_TIME)
			PersistState(context.Background(), sourceChainBlockTimeKey, strconv.Itoa(int(output.Int64())))
			log.Infof("Source chain block time set for data market %s to %s", strings.ToLower(dataMarketAddress.Hex()), strconv.Itoa(int(output.Int64())))
		}
	}
}

func DynamicStateVariables() {
	// Set total nodes count
	if output, err := Instance.GetTotalNodeCount(&bind.CallOpts{Context: context.Background()}); output != nil && err == nil {
		totalNodesCountKey := redis.TotalNodesCountKey()
		PersistState(context.Background(), totalNodesCountKey, strconv.Itoa(int(output.Int64())))
		log.Infof("Total nodes count set to %s", strconv.Itoa(int(output.Int64())))
	}

	// Iterate over all data markets and set dynamic state variables
	for _, dataMarketAddress := range config.SettingsObj.DataMarketContractAddresses {
		log.Infof("Setting dynamic state variables for data market: %s", dataMarketAddress)

		// Set current epoch
		if output, err := Instance.CurrentEpoch(&bind.CallOpts{}, dataMarketAddress); output.EpochId != nil && err == nil {
			currentEpochKey := redis.CurrentEpochID(strings.ToLower(dataMarketAddress.Hex()))
			PersistState(context.Background(), currentEpochKey, output.EpochId.String())
			log.Infof("Current epoch set for data market %s to %s", strings.ToLower(dataMarketAddress.Hex()), output.EpochId.String())
		}

		// Set day counter
		if output, err := Instance.DayCounter(&bind.CallOpts{}, dataMarketAddress); output != nil && err == nil {
			dayCounterKey := redis.DataMarketCurrentDay(dataMarketAddress.Hex())
			PersistState(context.Background(), dayCounterKey, strconv.Itoa(int(output.Int64())))
			log.Infof("Day counter set for data market %s to %s", strings.ToLower(dataMarketAddress.Hex()), strconv.Itoa(int(output.Int64())))
		}
	}
}

func PersistState(ctx context.Context, key, value string) {
	// Set the state variable in Redis
	if err := redis.Set(ctx, key, value, 0); err != nil {
		errMsg := fmt.Sprintf("Error setting state variable %s in Redis: %s", key, err.Error())
		reporting.SendFailureNotification(pkgs.PersistState, errMsg, time.Now().String(), "High")
		log.Error(errMsg)
	}

	// Persist the state variable in Redis
	if err := redis.PersistKey(ctx, key); err != nil {
		errMsg := fmt.Sprintf("Error persisting state variable %s in Redis: %s", key, err.Error())
		reporting.SendFailureNotification(pkgs.PersistState, errMsg, time.Now().String(), "High")
		log.Error(errMsg)
	}
}
