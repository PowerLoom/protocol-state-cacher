package prost

import (
	"context"
	"encoding/json"
	"github.com/cenkalti/backoff/v4"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
	"math/big"
	"protocol-state-cacher/config"
	"protocol-state-cacher/pkgs"
	listenerCommon "protocol-state-cacher/pkgs/common"
	"protocol-state-cacher/pkgs/contract"
	"protocol-state-cacher/pkgs/redis"
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
	var err error
	Client, err = ethclient.Dial(config.SettingsObj.ClientUrl)
	if err != nil {
		log.Fatal(err)
	}
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

func ColdSync() {
	for {
		PopulateStateVars()
		coldSyncAllSlots()
		time.Sleep((time.Millisecond * 500) * BlockTime)
	}
}

func coldSyncAllSlots() {
	var allSlots []string

	if slotCount, err := Instance.SlotCounter(&bind.CallOpts{}); slotCount != nil && err == nil {
		PersistState(context.Background(), redis.ContractStateVariable(pkgs.SlotCounter), slotCount.String())

		for i := int64(0); i <= slotCount.Int64(); i++ {
			slot, err := Instance.GetSlotInfo(&bind.CallOpts{}, DataMarket, big.NewInt(i))

			if slot == (contract.PowerloomDataMarketSlotInfo{}) || err != nil {
				log.Debugln("Error getting slot info: ", i)
				continue
			}

			slotMarshalled, err := json.Marshal(slot)

			if err != nil {
				log.Debugln("Error marshalling slot info: ", i)
				continue
			}

			PersistState(
				context.Background(),
				redis.SlotInfo(slot.SlotId.String()),
				string(slotMarshalled),
			)

			allSlots = append(
				allSlots,
				redis.SlotInfo(slot.SlotId.String()),
			)

			log.Debugln("Slot info: ", i, string(slotMarshalled), slot.SlotId.String())
		}

		err := redis.AddToSet(context.Background(), "AllSlotsInfo", allSlots...)
		if err != nil {
			log.Errorln("Error adding slots to set: ", err)
			listenerCommon.SendFailureNotification("ColdSync", err.Error(), time.Now().String(), "ERROR")
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
