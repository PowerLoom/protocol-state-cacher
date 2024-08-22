package prost

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"github.com/cenkalti/backoff/v4"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	log "github.com/sirupsen/logrus"
	"math/big"
	"net/http"
	"protocol-state-cacher/config"
	"protocol-state-cacher/pkgs"
	listenerCommon "protocol-state-cacher/pkgs/common"
	"protocol-state-cacher/pkgs/contract"
	"protocol-state-cacher/pkgs/redis"
	"sync"
	"time"
)

var Instance *contract.Contract

var (
	Client       *ethclient.Client
	CurrentBlock *types.Block
	DataMarket   common.Address
)

const BlockTime = 1

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
	DataMarket = common.HexToAddress(config.SettingsObj.DataMarketAddress)
	Instance, _ = contract.NewContract(common.HexToAddress(config.SettingsObj.ContractAddress), Client)
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
		time.Sleep(1 * time.Minute)
	}

}

func ColdSyncValues() {
	go func() {
		for {
			PopulateStateVars()
			time.Sleep((time.Millisecond * 500) * BlockTime)
		}
	}()
}

func coldSyncAllSlots() {
	if slotCount, err := Instance.SlotCounter(&bind.CallOpts{}); slotCount != nil && err == nil {
		PersistState(context.Background(), redis.ContractStateVariable(pkgs.SlotCounter), slotCount.String())

		var allSlots []string
		var mu sync.Mutex

		// TODO: MAKE THIS COUNT TO SLOT NUMBER
		for i := int64(0); i <= 6000; i += 20 {
			var wg sync.WaitGroup

			for j := i; j < i+20 && j <= 6000; j++ {
				wg.Add(1)
				go func(slotIndex int64) {
					defer wg.Done()
					slot, err := Instance.GetSlotInfo(&bind.CallOpts{}, DataMarket, big.NewInt(slotIndex))

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
		log.Errorln("Error getting slot counter: ", err)
		listenerCommon.SendFailureNotification("ColdSync", err.Error(), time.Now().String(), "ERROR")
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

	if output, err := Instance.CurrentEpoch(&bind.CallOpts{}, DataMarket); output.EpochId != nil && err == nil {
		key := redis.ContractStateVariable(pkgs.CurrentEpoch)
		PersistState(context.Background(), key, output.EpochId.String())
	}

	if output, err := Instance.CurrentBatchId(&bind.CallOpts{}, DataMarket); output != nil && err == nil {
		key := redis.ContractStateVariable(pkgs.CurrentBatchId)
		PersistState(context.Background(), key, output.String())
	}

	if output, err := Instance.EpochsInADay(&bind.CallOpts{}, DataMarket); output != nil && err == nil {
		key := redis.ContractStateVariable(pkgs.EpochsInADay)
		PersistState(context.Background(), key, output.String())
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
