package prost

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"math/big"
	"net/http"
	"protocol-state-cacher/config"
	"protocol-state-cacher/pkgs"
	listenerCommon "protocol-state-cacher/pkgs/common"
	"protocol-state-cacher/pkgs/contract"
	"protocol-state-cacher/pkgs/redis"
	"strconv"
	"sync"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	log "github.com/sirupsen/logrus"
)

var (
	Instances    map[common.Address]*contract.Contract
	Client       *ethclient.Client
	CurrentBlock *types.Block
)

const BlockTime = 1

func init() {
	Instances = make(map[common.Address]*contract.Contract)
}

func ConfigureClient() {
	rpcClient, err := rpc.DialOptions(
		context.Background(),
		config.SettingsObj.ClientUrl,
		rpc.WithHTTPClient(
			&http.Client{Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}},
		),
	)
	if err != nil {
		log.Fatal(err)
	}
	Client = ethclient.NewClient(rpcClient)
}

func ConfigureContractInstance() {
	for _, dataMarketAddr := range config.SettingsObj.DataMarketAddresses {
		dataMarket := common.HexToAddress(dataMarketAddr)
		instance, err := contract.NewContract(dataMarket, Client)
		if err != nil {
			log.Errorf("Failed to create contract instance for data market %s: %v", dataMarketAddr, err)
			continue
		}
		Instances[dataMarket] = instance
	}
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

func ColdSyncMappings() {
	for {
		coldSyncAllSlots()
		time.Sleep(60 * time.Minute)
	}

}

func ColdSyncValues() {
	go func() {
		for {
			PopulateStateVars()
			time.Sleep((time.Second * 10) * BlockTime)
		}
	}()
}

func coldSyncAllSlots() {
	for dataMarket, instance := range Instances {
		if slotCount, err := instance.SlotCounter(&bind.CallOpts{}); slotCount != nil && err == nil {
			PersistState(context.Background(), redis.ContractStateVariable(pkgs.SlotCounter), slotCount.String())

			var allSlots []string
			var mu sync.Mutex

			for i := int64(0); i <= 6000; i += 20 {
				var wg sync.WaitGroup

				for j := i; j < i+20 && j <= 6000; j++ {
					wg.Add(1)
					go func(slotIndex int64) {
						defer wg.Done()
						slot, err := instance.GetSlotInfo(&bind.CallOpts{}, dataMarket, big.NewInt(slotIndex))

						if slot == (contract.PowerloomDataMarketSlotInfo{}) || err != nil {
							log.Debugln("Error getting slot info: ", slotIndex)
							return
						}

						slotMarshalled, err := json.Marshal(slot)
						if err != nil {
							log.Debugln("Error marshalling slot info: ", slotIndex)
							return
						}

						PersistState(
							context.Background(),
							redis.SlotInfo(slot.SlotId.String()),
							string(slotMarshalled),
						)

						mu.Lock()
						allSlots = append(allSlots, redis.SlotInfo(slot.SlotId.String()))
						mu.Unlock()

						log.Debugln("Slot info: ", slotIndex, string(slotMarshalled), slot.SlotId.String())
					}(j)
				}

				wg.Wait()

				if len(allSlots) > 0 {
					mu.Lock()
					err := redis.AddToSet(context.Background(), "AllSlotsInfo", allSlots...)
					mu.Unlock()
					if err != nil {
						log.Errorln("Error adding slots to set: ", err)
						listenerCommon.SendFailureNotification("ColdSync", err.Error(), time.Now().String(), "ERROR")
					}
					allSlots = nil // reset batch
				}
			}
		} else {
			log.Errorln("Error getting slot counter for data market", dataMarket.String(), ":", err)
			listenerCommon.SendFailureNotification("ColdSync", err.Error(), time.Now().String(), "ERROR")
		}
		break // Exit after processing first instance
	}
}

func PopulateStateVars() {
	for {
		if block, err := Client.BlockByNumber(context.Background(), nil); err == nil {
			CurrentBlock = block
			break
		} else {
			log.Debugln("Encountered error while fetching current block: ", err.Error())
		}
	}

	for dataMarket, instance := range Instances {
		if output, err := instance.CurrentEpoch(&bind.CallOpts{}, dataMarket); output.EpochId != nil && err == nil {
			key := redis.ContractStateVariableWithDataMarket(dataMarket.String(), pkgs.CurrentEpoch)
			PersistState(context.Background(), key, output.EpochId.String())
		}

		if output, err := instance.EpochsInADay(&bind.CallOpts{}, dataMarket); output != nil && err == nil {
			key := redis.ContractStateVariableWithDataMarket(dataMarket.String(), pkgs.EpochsInADay)
			PersistState(context.Background(), key, output.String())
		}

		if output, err := instance.EPOCHSIZE(&bind.CallOpts{}, dataMarket); output != 0 && err == nil {
			key := redis.ContractStateVariableWithDataMarket(dataMarket.String(), pkgs.EPOCH_SIZE)
			PersistState(context.Background(), key, strconv.Itoa(int(output)))
		}

		if output, err := instance.SOURCECHAINBLOCKTIME(&bind.CallOpts{}, dataMarket); output != nil && err == nil {
			key := redis.ContractStateVariableWithDataMarket(dataMarket.String(), pkgs.SOURCE_CHAIN_BLOCK_TIME)
			PersistState(context.Background(), key, strconv.Itoa(int(output.Int64())))
		}
	}
}

func PersistState(ctx context.Context, key string, val string) {
	var err error
	if err = redis.Set(ctx, key, val, 0); err != nil {
		log.Errorln("Error setting state variable: ", key, val)
		listenerCommon.SendFailureNotification("PersistState", err.Error(), time.Now().String(), "ERROR")
	}
	if err = redis.PersistKey(ctx, key); err != nil {
		log.Errorln("Error persisting state variable: ", key)
		listenerCommon.SendFailureNotification("PersistState", err.Error(), time.Now().String(), "ERROR")
		return
	}
}
