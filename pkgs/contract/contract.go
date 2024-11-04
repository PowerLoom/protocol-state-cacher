// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// PowerloomDataMarketSlotInfo is an auto generated low-level Go binding around an user-defined struct.
type PowerloomDataMarketSlotInfo struct {
	SlotId                  *big.Int
	SnapshotterAddress      common.Address
	RewardPoints            *big.Int
	CurrentDaySnapshotCount *big.Int
}

// ContractMetaData contains all meta data concerning the Contract contract.
var ContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"dataMarketAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"adminAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"AdminsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"dataMarketAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"BatchSubmissionsCompleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"dataMarketAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"snapshotterAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"slotId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"dayId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewardPoints\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"DailyTaskCompletedEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"ownerAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"epochSize\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sourceChainBlockTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"useBlockNumberAsEpochId\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"protocolState\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"dataMarketAddress\",\"type\":\"address\"}],\"name\":\"DataMarketCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"dataMarketAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"dayId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"DayStartedEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"dataMarketAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"batchCid\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validatorAddr\",\"type\":\"address\"}],\"name\":\"DelayedAttestationSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"dataMarketAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"batchCid\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"DelayedBatchSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"dataMarketAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"snapshotterAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"slotId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"snapshotCid\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"projectId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"DelayedSnapshotSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"EmergencyWithdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"dataMarketAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"begin\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"EpochReleased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"RewardsClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"dataMarketAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sequencerAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"SequencersUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"dataMarketAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"batchCid\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validatorAddr\",\"type\":\"address\"}],\"name\":\"SnapshotBatchAttestationSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"dataMarketAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"batchCid\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"SnapshotBatchFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"dataMarketAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"batchCid\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"SnapshotBatchSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"dataMarketAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epochEnd\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"projectId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"snapshotCid\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"SnapshotFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"dataMarketAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"batchCid\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"TriggerBatchResubmission\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"dataMarketAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"batchCid\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"ValidatorAttestationsInvalidated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"dataMarketAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"validatorAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"ValidatorsUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"contractPowerloomDataMarket\",\"name\":\"dataMarket\",\"type\":\"address\"}],\"name\":\"DAY_SIZE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractPowerloomDataMarket\",\"name\":\"dataMarket\",\"type\":\"address\"}],\"name\":\"EPOCH_SIZE\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractPowerloomDataMarket\",\"name\":\"dataMarket\",\"type\":\"address\"}],\"name\":\"SOURCE_CHAIN_BLOCK_TIME\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractPowerloomDataMarket\",\"name\":\"dataMarket\",\"type\":\"address\"}],\"name\":\"SOURCE_CHAIN_ID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractPowerloomDataMarket\",\"name\":\"dataMarket\",\"type\":\"address\"}],\"name\":\"USE_BLOCK_NUMBER_AS_EPOCH_ID\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"allSnapshotters\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractPowerloomDataMarket\",\"name\":\"dataMarket\",\"type\":\"address\"}],\"name\":\"attestationSubmissionWindow\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractPowerloomDataMarket\",\"name\":\"dataMarket\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"batchCid\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"attestationsReceived\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractPowerloomDataMarket\",\"name\":\"dataMarket\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"batchCid\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"finalizedCidsRootHash\",\"type\":\"bytes32\"}],\"name\":\"attestationsReceivedCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractPowerloomDataMarket\",\"name\":\"dataMarket\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"batchCid\",\"type\":\"string\"}],\"name\":\"batchCidAttestationStatus\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractPowerloomDataMarket\",\"name\":\"dataMarket\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"batchCid\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"idx\",\"type\":\"uint256\"}],\"name\":\"batchCidDivergentValidators\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractPowerloomDataMarket\",\"name\":\"dataMarket\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"batchCid\",\"type\":\"string\"}],\"name\":\"batchCidSequencerAttestation\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractPowerloomDataMarket\",\"name\":\"dataMarket\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"batchCid\",\"type\":\"string\"}],\"name\":\"batchCidToProjects\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractPowerloomDataMarket\",\"name\":\"dataMarket\",\"type\":\"address\"}],\"name\":\"batchSubmissionWindow\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractPowerloomDataMarket\",\"name\":\"dataMarket\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"batchCid\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"}],\"name\":\"checkDynamicConsensusAttestations\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractPowerloomDataMarket\",\"name\":\"dataMarket\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"slotId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"day\",\"type\":\"uint256\"}],\"name\":\"checkSlotTaskStatusForDay\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"claimRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ownerAddress\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"epochSize\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sourceChainBlockTime\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"useBlockNumberAsEpochId\",\"type\":\"bool\"}],\"name\":\"createDataMarket\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractPowerloomDataMarket\",\"name\":\"dataMarket\",\"type\":\"address\"}],\"name\":\"currentEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"begin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractPowerloomDataMarket\",\"name\":\"dataMarket\",\"type\":\"address\"}],\"name\":\"dailySnapshotQuota\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dataMarketCount\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"dataMarketAddress\",\"type\":\"address\"}],\"name\":\"dataMarketEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"dataMarketFactory\",\"outputs\":[{\"internalType\":\"contractDataMarketFactory\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"dataMarketId\",\"type\":\"uint8\"}],\"name\":\"dataMarketIdToAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"dataMarketAddress\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"dataMarkets\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"ownerAddress\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"epochSize\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"sourceChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sourceChainBlockTime\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"useBlockNumberAsEpochId\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"dataMarketAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractPowerloomDataMarket\",\"name\":\"dataMarket\",\"type\":\"address\"}],\"name\":\"dayCounter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractPowerloomDataMarket\",\"name\":\"dataMarket\",\"type\":\"address\"}],\"name\":\"deploymentBlockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"emergencyWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"enabledNodeCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractPowerloomDataMarket\",\"name\":\"dataMarket\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"}],\"name\":\"endBatchSubmissions\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractPowerloomDataMarket\",\"name\":\"dataMarket\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"}],\"name\":\"epochIdToBatchCids\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractPowerloomDataMarket\",\"name\":\"dataMarket\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"}],\"name\":\"epochInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blocknumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochEnd\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractPowerloomDataMarket\",\"name\":\"dataMarket\",\"type\":\"address\"}],\"name\":\"epochManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractPowerloomDataMarket\",\"name\":\"dataMarket\",\"type\":\"address\"}],\"name\":\"epochsInADay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractPowerloomDataMarket\",\"name\":\"dataMarket\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"batchCid\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"}],\"name\":\"forceCompleteConsensusAttestations\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractPowerloomDataMarket\",\"name\":\"dataMarket\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"begin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"}],\"name\":\"forceSkipEpoch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractPowerloomDataMarket\",\"name\":\"dataMarket\",\"type\":\"address\"}],\"name\":\"getEpochManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractPowerloomDataMarket\",\"name\":\"dataMarket\",\"type\":\"address\"}],\"name\":\"getSequencerId\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractPowerloomDataMarket\",\"name\":\"dataMarket\",\"type\":\"address\"}],\"name\":\"getSequencers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractPowerloomDataMarket\",\"name\":\"dataMarket\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"slotId\",\"type\":\"uint256\"}],\"name\":\"getSlotInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"slotId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"snapshotterAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"rewardPoints\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentDaySnapshotCount\",\"type\":\"uint256\"}],\"internalType\":\"structPowerloomDataMarket.SlotInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"slotId\",\"type\":\"uint256\"}],\"name\":\"getSlotRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"rewards\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractPowerloomDataMarket\",\"name\":\"dataMarket\",\"type\":\"address\"}],\"name\":\"getTotalSequencersCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalSnapshotterCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractPowerloomDataMarket\",\"name\":\"dataMarket\",\"type\":\"address\"}],\"name\":\"getTotalValidatorsCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractPowerloomDataMarket\",\"name\":\"dataMarket\",\"type\":\"address\"}],\"name\":\"getValidators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"initialOwner\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractPowerloomDataMarket\",\"name\":\"dataMarket\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"projectId\",\"type\":\"string\"}],\"name\":\"lastFinalizedSnapshot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractPowerloomDataMarket\",\"name\":\"dataMarket\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"projectId\",\"type\":\"string\"}],\"name\":\"lastSequencerFinalizedSnapshot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractPowerloomDataMarket\",\"name\":\"dataMarket\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_dayCounter\",\"type\":\"uint256\"}],\"name\":\"loadCurrentDay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractPowerloomDataMarket\",\"name\":\"dataMarket\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"slotId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dayId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"snapshotCount\",\"type\":\"uint256\"}],\"name\":\"loadSlotSubmissions\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractPowerloomDataMarket\",\"name\":\"dataMarket\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"batchCid\",\"type\":\"string\"}],\"name\":\"maxAttestationFinalizedRootHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractPowerloomDataMarket\",\"name\":\"dataMarket\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"batchCid\",\"type\":\"string\"}],\"name\":\"maxAttestationsCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractPowerloomDataMarket\",\"name\":\"dataMarket\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"projectId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"}],\"name\":\"maxSnapshotsCid\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractPowerloomDataMarket\",\"name\":\"dataMarket\",\"type\":\"address\"}],\"name\":\"minAttestationsForConsensus\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractPowerloomDataMarket\",\"name\":\"dataMarket\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"projectId\",\"type\":\"string\"}],\"name\":\"projectFirstEpochId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractPowerloomDataMarket\",\"name\":\"dataMarket\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"begin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"}],\"name\":\"releaseEpoch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractPowerloomDataMarket\",\"name\":\"dataMarket\",\"type\":\"address\"}],\"name\":\"rewardPoolSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractPowerloomDataMarket\",\"name\":\"dataMarket\",\"type\":\"address\"}],\"name\":\"rewardsEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractPowerloomDataMarket\",\"name\":\"dataMarket\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_sequencerId\",\"type\":\"string\"}],\"name\":\"setSequencerId\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractPowerloomDataMarket\",\"name\":\"dataMarket\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"slotId\",\"type\":\"uint256\"}],\"name\":\"slotRewardPoints\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"slotRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"slotId\",\"type\":\"uint256\"}],\"name\":\"slotSnapshotterMapping\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractPowerloomDataMarket\",\"name\":\"dataMarket\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"slotId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dayId\",\"type\":\"uint256\"}],\"name\":\"slotSubmissionCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractPowerloomDataMarket\",\"name\":\"dataMarket\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"projectId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"}],\"name\":\"snapshotStatus\",\"outputs\":[{\"internalType\":\"enumPowerloomDataMarket.SnapshotStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"snapshotCid\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractPowerloomDataMarket\",\"name\":\"dataMarket\",\"type\":\"address\"}],\"name\":\"snapshotSubmissionWindow\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"snapshotterState\",\"outputs\":[{\"internalType\":\"contractPowerloomNodes\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractPowerloomDataMarket\",\"name\":\"dataMarket\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"batchCid\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"finalizedCidsRootHash\",\"type\":\"bytes32\"}],\"name\":\"submitBatchAttestation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractPowerloomDataMarket\",\"name\":\"dataMarket\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"batchCid\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"epochId\",\"type\":\"uint256\"},{\"internalType\":\"string[]\",\"name\":\"projectIds\",\"type\":\"string[]\"},{\"internalType\":\"string[]\",\"name\":\"snapshotCids\",\"type\":\"string[]\"},{\"internalType\":\"bytes32\",\"name\":\"finalizedCidsRootHash\",\"type\":\"bytes32\"}],\"name\":\"submitSubmissionBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"dataMarketAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"toggleDataMarket\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractPowerloomDataMarket\",\"name\":\"dataMarket\",\"type\":\"address\"}],\"name\":\"toggleRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractPowerloomDataMarket\",\"name\":\"dataMarket\",\"type\":\"address\"},{\"internalType\":\"enumPowerloomDataMarket.Role\",\"name\":\"role\",\"type\":\"uint8\"},{\"internalType\":\"address[]\",\"name\":\"_addresses\",\"type\":\"address[]\"},{\"internalType\":\"bool[]\",\"name\":\"_status\",\"type\":\"bool[]\"}],\"name\":\"updateAddresses\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractPowerloomDataMarket\",\"name\":\"dataMarket\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"newattestationSubmissionWindow\",\"type\":\"uint256\"}],\"name\":\"updateAttestationSubmissionWindow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractPowerloomDataMarket\",\"name\":\"dataMarket\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"newbatchSubmissionWindow\",\"type\":\"uint256\"}],\"name\":\"updateBatchSubmissionWindow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractPowerloomDataMarket\",\"name\":\"dataMarket\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_dailySnapshotQuota\",\"type\":\"uint256\"}],\"name\":\"updateDailySnapshotQuota\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"updateDataMarketFactory\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractPowerloomDataMarket\",\"name\":\"dataMarket\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"newDaySize\",\"type\":\"uint256\"}],\"name\":\"updateDaySize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractPowerloomDataMarket\",\"name\":\"dataMarket\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"updateEpochManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractPowerloomDataMarket\",\"name\":\"dataMarket\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_minAttestationsForConsensus\",\"type\":\"uint256\"}],\"name\":\"updateMinAttestationsForConsensus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractPowerloomDataMarket\",\"name\":\"dataMarket\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"newRewardPoolSize\",\"type\":\"uint256\"}],\"name\":\"updateRewardPoolSize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractPowerloomDataMarket\",\"name\":\"dataMarket\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"slotIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"submissionsList\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"day\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"eligibleNodes\",\"type\":\"uint256\"}],\"name\":\"updateRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractPowerloomDataMarket\",\"name\":\"dataMarket\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"newsnapshotSubmissionWindow\",\"type\":\"uint256\"}],\"name\":\"updateSnapshotSubmissionWindow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"updateSnapshotterState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalRewards\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalClaimed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastClaimed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastUpdated\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ContractABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractMetaData.ABI instead.
var ContractABI = ContractMetaData.ABI

// Contract is an auto generated Go binding around an Ethereum contract.
type Contract struct {
	ContractCaller     // Read-only binding to the contract
	ContractTransactor // Write-only binding to the contract
	ContractFilterer   // Log filterer for contract events
}

// ContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractSession struct {
	Contract     *Contract         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractCallerSession struct {
	Contract *ContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractTransactorSession struct {
	Contract     *ContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractRaw struct {
	Contract *Contract // Generic contract binding to access the raw methods on
}

// ContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractCallerRaw struct {
	Contract *ContractCaller // Generic read-only contract binding to access the raw methods on
}

// ContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractTransactorRaw struct {
	Contract *ContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContract creates a new instance of Contract, bound to a specific deployed contract.
func NewContract(address common.Address, backend bind.ContractBackend) (*Contract, error) {
	contract, err := bindContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// NewContractCaller creates a new read-only instance of Contract, bound to a specific deployed contract.
func NewContractCaller(address common.Address, caller bind.ContractCaller) (*ContractCaller, error) {
	contract, err := bindContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractCaller{contract: contract}, nil
}

// NewContractTransactor creates a new write-only instance of Contract, bound to a specific deployed contract.
func NewContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractTransactor, error) {
	contract, err := bindContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTransactor{contract: contract}, nil
}

// NewContractFilterer creates a new log filterer instance of Contract, bound to a specific deployed contract.
func NewContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractFilterer, error) {
	contract, err := bindContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractFilterer{contract: contract}, nil
}

// bindContract binds a generic wrapper to an already deployed contract.
func bindContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.ContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transact(opts, method, params...)
}

// DAYSIZE is a free data retrieval call binding the contract method 0x04a0a5bb.
//
// Solidity: function DAY_SIZE(address dataMarket) view returns(uint256)
func (_Contract *ContractCaller) DAYSIZE(opts *bind.CallOpts, dataMarket common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "DAY_SIZE", dataMarket)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DAYSIZE is a free data retrieval call binding the contract method 0x04a0a5bb.
//
// Solidity: function DAY_SIZE(address dataMarket) view returns(uint256)
func (_Contract *ContractSession) DAYSIZE(dataMarket common.Address) (*big.Int, error) {
	return _Contract.Contract.DAYSIZE(&_Contract.CallOpts, dataMarket)
}

// DAYSIZE is a free data retrieval call binding the contract method 0x04a0a5bb.
//
// Solidity: function DAY_SIZE(address dataMarket) view returns(uint256)
func (_Contract *ContractCallerSession) DAYSIZE(dataMarket common.Address) (*big.Int, error) {
	return _Contract.Contract.DAYSIZE(&_Contract.CallOpts, dataMarket)
}

// EPOCHSIZE is a free data retrieval call binding the contract method 0xc12c2aa9.
//
// Solidity: function EPOCH_SIZE(address dataMarket) view returns(uint8)
func (_Contract *ContractCaller) EPOCHSIZE(opts *bind.CallOpts, dataMarket common.Address) (uint8, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "EPOCH_SIZE", dataMarket)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// EPOCHSIZE is a free data retrieval call binding the contract method 0xc12c2aa9.
//
// Solidity: function EPOCH_SIZE(address dataMarket) view returns(uint8)
func (_Contract *ContractSession) EPOCHSIZE(dataMarket common.Address) (uint8, error) {
	return _Contract.Contract.EPOCHSIZE(&_Contract.CallOpts, dataMarket)
}

// EPOCHSIZE is a free data retrieval call binding the contract method 0xc12c2aa9.
//
// Solidity: function EPOCH_SIZE(address dataMarket) view returns(uint8)
func (_Contract *ContractCallerSession) EPOCHSIZE(dataMarket common.Address) (uint8, error) {
	return _Contract.Contract.EPOCHSIZE(&_Contract.CallOpts, dataMarket)
}

// SOURCECHAINBLOCKTIME is a free data retrieval call binding the contract method 0x0f34e6a9.
//
// Solidity: function SOURCE_CHAIN_BLOCK_TIME(address dataMarket) view returns(uint256)
func (_Contract *ContractCaller) SOURCECHAINBLOCKTIME(opts *bind.CallOpts, dataMarket common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "SOURCE_CHAIN_BLOCK_TIME", dataMarket)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SOURCECHAINBLOCKTIME is a free data retrieval call binding the contract method 0x0f34e6a9.
//
// Solidity: function SOURCE_CHAIN_BLOCK_TIME(address dataMarket) view returns(uint256)
func (_Contract *ContractSession) SOURCECHAINBLOCKTIME(dataMarket common.Address) (*big.Int, error) {
	return _Contract.Contract.SOURCECHAINBLOCKTIME(&_Contract.CallOpts, dataMarket)
}

// SOURCECHAINBLOCKTIME is a free data retrieval call binding the contract method 0x0f34e6a9.
//
// Solidity: function SOURCE_CHAIN_BLOCK_TIME(address dataMarket) view returns(uint256)
func (_Contract *ContractCallerSession) SOURCECHAINBLOCKTIME(dataMarket common.Address) (*big.Int, error) {
	return _Contract.Contract.SOURCECHAINBLOCKTIME(&_Contract.CallOpts, dataMarket)
}

// SOURCECHAINID is a free data retrieval call binding the contract method 0x40e29707.
//
// Solidity: function SOURCE_CHAIN_ID(address dataMarket) view returns(uint256)
func (_Contract *ContractCaller) SOURCECHAINID(opts *bind.CallOpts, dataMarket common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "SOURCE_CHAIN_ID", dataMarket)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SOURCECHAINID is a free data retrieval call binding the contract method 0x40e29707.
//
// Solidity: function SOURCE_CHAIN_ID(address dataMarket) view returns(uint256)
func (_Contract *ContractSession) SOURCECHAINID(dataMarket common.Address) (*big.Int, error) {
	return _Contract.Contract.SOURCECHAINID(&_Contract.CallOpts, dataMarket)
}

// SOURCECHAINID is a free data retrieval call binding the contract method 0x40e29707.
//
// Solidity: function SOURCE_CHAIN_ID(address dataMarket) view returns(uint256)
func (_Contract *ContractCallerSession) SOURCECHAINID(dataMarket common.Address) (*big.Int, error) {
	return _Contract.Contract.SOURCECHAINID(&_Contract.CallOpts, dataMarket)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Contract *ContractCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Contract *ContractSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _Contract.Contract.UPGRADEINTERFACEVERSION(&_Contract.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Contract *ContractCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _Contract.Contract.UPGRADEINTERFACEVERSION(&_Contract.CallOpts)
}

// USEBLOCKNUMBERASEPOCHID is a free data retrieval call binding the contract method 0x865fb4eb.
//
// Solidity: function USE_BLOCK_NUMBER_AS_EPOCH_ID(address dataMarket) view returns(bool)
func (_Contract *ContractCaller) USEBLOCKNUMBERASEPOCHID(opts *bind.CallOpts, dataMarket common.Address) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "USE_BLOCK_NUMBER_AS_EPOCH_ID", dataMarket)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// USEBLOCKNUMBERASEPOCHID is a free data retrieval call binding the contract method 0x865fb4eb.
//
// Solidity: function USE_BLOCK_NUMBER_AS_EPOCH_ID(address dataMarket) view returns(bool)
func (_Contract *ContractSession) USEBLOCKNUMBERASEPOCHID(dataMarket common.Address) (bool, error) {
	return _Contract.Contract.USEBLOCKNUMBERASEPOCHID(&_Contract.CallOpts, dataMarket)
}

// USEBLOCKNUMBERASEPOCHID is a free data retrieval call binding the contract method 0x865fb4eb.
//
// Solidity: function USE_BLOCK_NUMBER_AS_EPOCH_ID(address dataMarket) view returns(bool)
func (_Contract *ContractCallerSession) USEBLOCKNUMBERASEPOCHID(dataMarket common.Address) (bool, error) {
	return _Contract.Contract.USEBLOCKNUMBERASEPOCHID(&_Contract.CallOpts, dataMarket)
}

// AllSnapshotters is a free data retrieval call binding the contract method 0x3d15d0f4.
//
// Solidity: function allSnapshotters(address addr) view returns(bool)
func (_Contract *ContractCaller) AllSnapshotters(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "allSnapshotters", addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllSnapshotters is a free data retrieval call binding the contract method 0x3d15d0f4.
//
// Solidity: function allSnapshotters(address addr) view returns(bool)
func (_Contract *ContractSession) AllSnapshotters(addr common.Address) (bool, error) {
	return _Contract.Contract.AllSnapshotters(&_Contract.CallOpts, addr)
}

// AllSnapshotters is a free data retrieval call binding the contract method 0x3d15d0f4.
//
// Solidity: function allSnapshotters(address addr) view returns(bool)
func (_Contract *ContractCallerSession) AllSnapshotters(addr common.Address) (bool, error) {
	return _Contract.Contract.AllSnapshotters(&_Contract.CallOpts, addr)
}

// AttestationSubmissionWindow is a free data retrieval call binding the contract method 0xe1d5fbce.
//
// Solidity: function attestationSubmissionWindow(address dataMarket) view returns(uint256)
func (_Contract *ContractCaller) AttestationSubmissionWindow(opts *bind.CallOpts, dataMarket common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "attestationSubmissionWindow", dataMarket)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AttestationSubmissionWindow is a free data retrieval call binding the contract method 0xe1d5fbce.
//
// Solidity: function attestationSubmissionWindow(address dataMarket) view returns(uint256)
func (_Contract *ContractSession) AttestationSubmissionWindow(dataMarket common.Address) (*big.Int, error) {
	return _Contract.Contract.AttestationSubmissionWindow(&_Contract.CallOpts, dataMarket)
}

// AttestationSubmissionWindow is a free data retrieval call binding the contract method 0xe1d5fbce.
//
// Solidity: function attestationSubmissionWindow(address dataMarket) view returns(uint256)
func (_Contract *ContractCallerSession) AttestationSubmissionWindow(dataMarket common.Address) (*big.Int, error) {
	return _Contract.Contract.AttestationSubmissionWindow(&_Contract.CallOpts, dataMarket)
}

// AttestationsReceived is a free data retrieval call binding the contract method 0x53a5a874.
//
// Solidity: function attestationsReceived(address dataMarket, string batchCid, address validator) view returns(bool)
func (_Contract *ContractCaller) AttestationsReceived(opts *bind.CallOpts, dataMarket common.Address, batchCid string, validator common.Address) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "attestationsReceived", dataMarket, batchCid, validator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AttestationsReceived is a free data retrieval call binding the contract method 0x53a5a874.
//
// Solidity: function attestationsReceived(address dataMarket, string batchCid, address validator) view returns(bool)
func (_Contract *ContractSession) AttestationsReceived(dataMarket common.Address, batchCid string, validator common.Address) (bool, error) {
	return _Contract.Contract.AttestationsReceived(&_Contract.CallOpts, dataMarket, batchCid, validator)
}

// AttestationsReceived is a free data retrieval call binding the contract method 0x53a5a874.
//
// Solidity: function attestationsReceived(address dataMarket, string batchCid, address validator) view returns(bool)
func (_Contract *ContractCallerSession) AttestationsReceived(dataMarket common.Address, batchCid string, validator common.Address) (bool, error) {
	return _Contract.Contract.AttestationsReceived(&_Contract.CallOpts, dataMarket, batchCid, validator)
}

// AttestationsReceivedCount is a free data retrieval call binding the contract method 0x751b8eeb.
//
// Solidity: function attestationsReceivedCount(address dataMarket, string batchCid, bytes32 finalizedCidsRootHash) view returns(uint256)
func (_Contract *ContractCaller) AttestationsReceivedCount(opts *bind.CallOpts, dataMarket common.Address, batchCid string, finalizedCidsRootHash [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "attestationsReceivedCount", dataMarket, batchCid, finalizedCidsRootHash)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AttestationsReceivedCount is a free data retrieval call binding the contract method 0x751b8eeb.
//
// Solidity: function attestationsReceivedCount(address dataMarket, string batchCid, bytes32 finalizedCidsRootHash) view returns(uint256)
func (_Contract *ContractSession) AttestationsReceivedCount(dataMarket common.Address, batchCid string, finalizedCidsRootHash [32]byte) (*big.Int, error) {
	return _Contract.Contract.AttestationsReceivedCount(&_Contract.CallOpts, dataMarket, batchCid, finalizedCidsRootHash)
}

// AttestationsReceivedCount is a free data retrieval call binding the contract method 0x751b8eeb.
//
// Solidity: function attestationsReceivedCount(address dataMarket, string batchCid, bytes32 finalizedCidsRootHash) view returns(uint256)
func (_Contract *ContractCallerSession) AttestationsReceivedCount(dataMarket common.Address, batchCid string, finalizedCidsRootHash [32]byte) (*big.Int, error) {
	return _Contract.Contract.AttestationsReceivedCount(&_Contract.CallOpts, dataMarket, batchCid, finalizedCidsRootHash)
}

// BatchCidAttestationStatus is a free data retrieval call binding the contract method 0xb646154f.
//
// Solidity: function batchCidAttestationStatus(address dataMarket, string batchCid) view returns(bool)
func (_Contract *ContractCaller) BatchCidAttestationStatus(opts *bind.CallOpts, dataMarket common.Address, batchCid string) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "batchCidAttestationStatus", dataMarket, batchCid)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BatchCidAttestationStatus is a free data retrieval call binding the contract method 0xb646154f.
//
// Solidity: function batchCidAttestationStatus(address dataMarket, string batchCid) view returns(bool)
func (_Contract *ContractSession) BatchCidAttestationStatus(dataMarket common.Address, batchCid string) (bool, error) {
	return _Contract.Contract.BatchCidAttestationStatus(&_Contract.CallOpts, dataMarket, batchCid)
}

// BatchCidAttestationStatus is a free data retrieval call binding the contract method 0xb646154f.
//
// Solidity: function batchCidAttestationStatus(address dataMarket, string batchCid) view returns(bool)
func (_Contract *ContractCallerSession) BatchCidAttestationStatus(dataMarket common.Address, batchCid string) (bool, error) {
	return _Contract.Contract.BatchCidAttestationStatus(&_Contract.CallOpts, dataMarket, batchCid)
}

// BatchCidDivergentValidators is a free data retrieval call binding the contract method 0x4ecac85a.
//
// Solidity: function batchCidDivergentValidators(address dataMarket, string batchCid, uint256 idx) view returns(address)
func (_Contract *ContractCaller) BatchCidDivergentValidators(opts *bind.CallOpts, dataMarket common.Address, batchCid string, idx *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "batchCidDivergentValidators", dataMarket, batchCid, idx)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BatchCidDivergentValidators is a free data retrieval call binding the contract method 0x4ecac85a.
//
// Solidity: function batchCidDivergentValidators(address dataMarket, string batchCid, uint256 idx) view returns(address)
func (_Contract *ContractSession) BatchCidDivergentValidators(dataMarket common.Address, batchCid string, idx *big.Int) (common.Address, error) {
	return _Contract.Contract.BatchCidDivergentValidators(&_Contract.CallOpts, dataMarket, batchCid, idx)
}

// BatchCidDivergentValidators is a free data retrieval call binding the contract method 0x4ecac85a.
//
// Solidity: function batchCidDivergentValidators(address dataMarket, string batchCid, uint256 idx) view returns(address)
func (_Contract *ContractCallerSession) BatchCidDivergentValidators(dataMarket common.Address, batchCid string, idx *big.Int) (common.Address, error) {
	return _Contract.Contract.BatchCidDivergentValidators(&_Contract.CallOpts, dataMarket, batchCid, idx)
}

// BatchCidSequencerAttestation is a free data retrieval call binding the contract method 0xc1045a5f.
//
// Solidity: function batchCidSequencerAttestation(address dataMarket, string batchCid) view returns(bytes32)
func (_Contract *ContractCaller) BatchCidSequencerAttestation(opts *bind.CallOpts, dataMarket common.Address, batchCid string) ([32]byte, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "batchCidSequencerAttestation", dataMarket, batchCid)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BatchCidSequencerAttestation is a free data retrieval call binding the contract method 0xc1045a5f.
//
// Solidity: function batchCidSequencerAttestation(address dataMarket, string batchCid) view returns(bytes32)
func (_Contract *ContractSession) BatchCidSequencerAttestation(dataMarket common.Address, batchCid string) ([32]byte, error) {
	return _Contract.Contract.BatchCidSequencerAttestation(&_Contract.CallOpts, dataMarket, batchCid)
}

// BatchCidSequencerAttestation is a free data retrieval call binding the contract method 0xc1045a5f.
//
// Solidity: function batchCidSequencerAttestation(address dataMarket, string batchCid) view returns(bytes32)
func (_Contract *ContractCallerSession) BatchCidSequencerAttestation(dataMarket common.Address, batchCid string) ([32]byte, error) {
	return _Contract.Contract.BatchCidSequencerAttestation(&_Contract.CallOpts, dataMarket, batchCid)
}

// BatchCidToProjects is a free data retrieval call binding the contract method 0xe06a9ecb.
//
// Solidity: function batchCidToProjects(address dataMarket, string batchCid) view returns(string[])
func (_Contract *ContractCaller) BatchCidToProjects(opts *bind.CallOpts, dataMarket common.Address, batchCid string) ([]string, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "batchCidToProjects", dataMarket, batchCid)

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// BatchCidToProjects is a free data retrieval call binding the contract method 0xe06a9ecb.
//
// Solidity: function batchCidToProjects(address dataMarket, string batchCid) view returns(string[])
func (_Contract *ContractSession) BatchCidToProjects(dataMarket common.Address, batchCid string) ([]string, error) {
	return _Contract.Contract.BatchCidToProjects(&_Contract.CallOpts, dataMarket, batchCid)
}

// BatchCidToProjects is a free data retrieval call binding the contract method 0xe06a9ecb.
//
// Solidity: function batchCidToProjects(address dataMarket, string batchCid) view returns(string[])
func (_Contract *ContractCallerSession) BatchCidToProjects(dataMarket common.Address, batchCid string) ([]string, error) {
	return _Contract.Contract.BatchCidToProjects(&_Contract.CallOpts, dataMarket, batchCid)
}

// BatchSubmissionWindow is a free data retrieval call binding the contract method 0x4d9c25d4.
//
// Solidity: function batchSubmissionWindow(address dataMarket) view returns(uint256)
func (_Contract *ContractCaller) BatchSubmissionWindow(opts *bind.CallOpts, dataMarket common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "batchSubmissionWindow", dataMarket)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BatchSubmissionWindow is a free data retrieval call binding the contract method 0x4d9c25d4.
//
// Solidity: function batchSubmissionWindow(address dataMarket) view returns(uint256)
func (_Contract *ContractSession) BatchSubmissionWindow(dataMarket common.Address) (*big.Int, error) {
	return _Contract.Contract.BatchSubmissionWindow(&_Contract.CallOpts, dataMarket)
}

// BatchSubmissionWindow is a free data retrieval call binding the contract method 0x4d9c25d4.
//
// Solidity: function batchSubmissionWindow(address dataMarket) view returns(uint256)
func (_Contract *ContractCallerSession) BatchSubmissionWindow(dataMarket common.Address) (*big.Int, error) {
	return _Contract.Contract.BatchSubmissionWindow(&_Contract.CallOpts, dataMarket)
}

// CheckDynamicConsensusAttestations is a free data retrieval call binding the contract method 0x687d738d.
//
// Solidity: function checkDynamicConsensusAttestations(address dataMarket, string batchCid, uint256 epochId) view returns(bool)
func (_Contract *ContractCaller) CheckDynamicConsensusAttestations(opts *bind.CallOpts, dataMarket common.Address, batchCid string, epochId *big.Int) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "checkDynamicConsensusAttestations", dataMarket, batchCid, epochId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckDynamicConsensusAttestations is a free data retrieval call binding the contract method 0x687d738d.
//
// Solidity: function checkDynamicConsensusAttestations(address dataMarket, string batchCid, uint256 epochId) view returns(bool)
func (_Contract *ContractSession) CheckDynamicConsensusAttestations(dataMarket common.Address, batchCid string, epochId *big.Int) (bool, error) {
	return _Contract.Contract.CheckDynamicConsensusAttestations(&_Contract.CallOpts, dataMarket, batchCid, epochId)
}

// CheckDynamicConsensusAttestations is a free data retrieval call binding the contract method 0x687d738d.
//
// Solidity: function checkDynamicConsensusAttestations(address dataMarket, string batchCid, uint256 epochId) view returns(bool)
func (_Contract *ContractCallerSession) CheckDynamicConsensusAttestations(dataMarket common.Address, batchCid string, epochId *big.Int) (bool, error) {
	return _Contract.Contract.CheckDynamicConsensusAttestations(&_Contract.CallOpts, dataMarket, batchCid, epochId)
}

// CheckSlotTaskStatusForDay is a free data retrieval call binding the contract method 0xc00d0f9c.
//
// Solidity: function checkSlotTaskStatusForDay(address dataMarket, uint256 slotId, uint256 day) view returns(bool)
func (_Contract *ContractCaller) CheckSlotTaskStatusForDay(opts *bind.CallOpts, dataMarket common.Address, slotId *big.Int, day *big.Int) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "checkSlotTaskStatusForDay", dataMarket, slotId, day)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckSlotTaskStatusForDay is a free data retrieval call binding the contract method 0xc00d0f9c.
//
// Solidity: function checkSlotTaskStatusForDay(address dataMarket, uint256 slotId, uint256 day) view returns(bool)
func (_Contract *ContractSession) CheckSlotTaskStatusForDay(dataMarket common.Address, slotId *big.Int, day *big.Int) (bool, error) {
	return _Contract.Contract.CheckSlotTaskStatusForDay(&_Contract.CallOpts, dataMarket, slotId, day)
}

// CheckSlotTaskStatusForDay is a free data retrieval call binding the contract method 0xc00d0f9c.
//
// Solidity: function checkSlotTaskStatusForDay(address dataMarket, uint256 slotId, uint256 day) view returns(bool)
func (_Contract *ContractCallerSession) CheckSlotTaskStatusForDay(dataMarket common.Address, slotId *big.Int, day *big.Int) (bool, error) {
	return _Contract.Contract.CheckSlotTaskStatusForDay(&_Contract.CallOpts, dataMarket, slotId, day)
}

// CurrentEpoch is a free data retrieval call binding the contract method 0x0736e19f.
//
// Solidity: function currentEpoch(address dataMarket) view returns(uint256 begin, uint256 end, uint256 epochId)
func (_Contract *ContractCaller) CurrentEpoch(opts *bind.CallOpts, dataMarket common.Address) (struct {
	Begin   *big.Int
	End     *big.Int
	EpochId *big.Int
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "currentEpoch", dataMarket)

	outstruct := new(struct {
		Begin   *big.Int
		End     *big.Int
		EpochId *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Begin = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.End = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.EpochId = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// CurrentEpoch is a free data retrieval call binding the contract method 0x0736e19f.
//
// Solidity: function currentEpoch(address dataMarket) view returns(uint256 begin, uint256 end, uint256 epochId)
func (_Contract *ContractSession) CurrentEpoch(dataMarket common.Address) (struct {
	Begin   *big.Int
	End     *big.Int
	EpochId *big.Int
}, error) {
	return _Contract.Contract.CurrentEpoch(&_Contract.CallOpts, dataMarket)
}

// CurrentEpoch is a free data retrieval call binding the contract method 0x0736e19f.
//
// Solidity: function currentEpoch(address dataMarket) view returns(uint256 begin, uint256 end, uint256 epochId)
func (_Contract *ContractCallerSession) CurrentEpoch(dataMarket common.Address) (struct {
	Begin   *big.Int
	End     *big.Int
	EpochId *big.Int
}, error) {
	return _Contract.Contract.CurrentEpoch(&_Contract.CallOpts, dataMarket)
}

// DailySnapshotQuota is a free data retrieval call binding the contract method 0x095cb210.
//
// Solidity: function dailySnapshotQuota(address dataMarket) view returns(uint256)
func (_Contract *ContractCaller) DailySnapshotQuota(opts *bind.CallOpts, dataMarket common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "dailySnapshotQuota", dataMarket)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DailySnapshotQuota is a free data retrieval call binding the contract method 0x095cb210.
//
// Solidity: function dailySnapshotQuota(address dataMarket) view returns(uint256)
func (_Contract *ContractSession) DailySnapshotQuota(dataMarket common.Address) (*big.Int, error) {
	return _Contract.Contract.DailySnapshotQuota(&_Contract.CallOpts, dataMarket)
}

// DailySnapshotQuota is a free data retrieval call binding the contract method 0x095cb210.
//
// Solidity: function dailySnapshotQuota(address dataMarket) view returns(uint256)
func (_Contract *ContractCallerSession) DailySnapshotQuota(dataMarket common.Address) (*big.Int, error) {
	return _Contract.Contract.DailySnapshotQuota(&_Contract.CallOpts, dataMarket)
}

// DataMarketCount is a free data retrieval call binding the contract method 0xef829a3d.
//
// Solidity: function dataMarketCount() view returns(uint8)
func (_Contract *ContractCaller) DataMarketCount(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "dataMarketCount")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// DataMarketCount is a free data retrieval call binding the contract method 0xef829a3d.
//
// Solidity: function dataMarketCount() view returns(uint8)
func (_Contract *ContractSession) DataMarketCount() (uint8, error) {
	return _Contract.Contract.DataMarketCount(&_Contract.CallOpts)
}

// DataMarketCount is a free data retrieval call binding the contract method 0xef829a3d.
//
// Solidity: function dataMarketCount() view returns(uint8)
func (_Contract *ContractCallerSession) DataMarketCount() (uint8, error) {
	return _Contract.Contract.DataMarketCount(&_Contract.CallOpts)
}

// DataMarketEnabled is a free data retrieval call binding the contract method 0x75fd5c7c.
//
// Solidity: function dataMarketEnabled(address dataMarketAddress) view returns(bool)
func (_Contract *ContractCaller) DataMarketEnabled(opts *bind.CallOpts, dataMarketAddress common.Address) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "dataMarketEnabled", dataMarketAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// DataMarketEnabled is a free data retrieval call binding the contract method 0x75fd5c7c.
//
// Solidity: function dataMarketEnabled(address dataMarketAddress) view returns(bool)
func (_Contract *ContractSession) DataMarketEnabled(dataMarketAddress common.Address) (bool, error) {
	return _Contract.Contract.DataMarketEnabled(&_Contract.CallOpts, dataMarketAddress)
}

// DataMarketEnabled is a free data retrieval call binding the contract method 0x75fd5c7c.
//
// Solidity: function dataMarketEnabled(address dataMarketAddress) view returns(bool)
func (_Contract *ContractCallerSession) DataMarketEnabled(dataMarketAddress common.Address) (bool, error) {
	return _Contract.Contract.DataMarketEnabled(&_Contract.CallOpts, dataMarketAddress)
}

// DataMarketFactory is a free data retrieval call binding the contract method 0x5a5c908b.
//
// Solidity: function dataMarketFactory() view returns(address)
func (_Contract *ContractCaller) DataMarketFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "dataMarketFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DataMarketFactory is a free data retrieval call binding the contract method 0x5a5c908b.
//
// Solidity: function dataMarketFactory() view returns(address)
func (_Contract *ContractSession) DataMarketFactory() (common.Address, error) {
	return _Contract.Contract.DataMarketFactory(&_Contract.CallOpts)
}

// DataMarketFactory is a free data retrieval call binding the contract method 0x5a5c908b.
//
// Solidity: function dataMarketFactory() view returns(address)
func (_Contract *ContractCallerSession) DataMarketFactory() (common.Address, error) {
	return _Contract.Contract.DataMarketFactory(&_Contract.CallOpts)
}

// DataMarketIdToAddress is a free data retrieval call binding the contract method 0x0857b13f.
//
// Solidity: function dataMarketIdToAddress(uint8 dataMarketId) view returns(address dataMarketAddress)
func (_Contract *ContractCaller) DataMarketIdToAddress(opts *bind.CallOpts, dataMarketId uint8) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "dataMarketIdToAddress", dataMarketId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DataMarketIdToAddress is a free data retrieval call binding the contract method 0x0857b13f.
//
// Solidity: function dataMarketIdToAddress(uint8 dataMarketId) view returns(address dataMarketAddress)
func (_Contract *ContractSession) DataMarketIdToAddress(dataMarketId uint8) (common.Address, error) {
	return _Contract.Contract.DataMarketIdToAddress(&_Contract.CallOpts, dataMarketId)
}

// DataMarketIdToAddress is a free data retrieval call binding the contract method 0x0857b13f.
//
// Solidity: function dataMarketIdToAddress(uint8 dataMarketId) view returns(address dataMarketAddress)
func (_Contract *ContractCallerSession) DataMarketIdToAddress(dataMarketId uint8) (common.Address, error) {
	return _Contract.Contract.DataMarketIdToAddress(&_Contract.CallOpts, dataMarketId)
}

// DataMarkets is a free data retrieval call binding the contract method 0x5f9abf70.
//
// Solidity: function dataMarkets(address ) view returns(address ownerAddress, uint8 epochSize, uint256 sourceChainId, uint256 sourceChainBlockTime, bool useBlockNumberAsEpochId, bool enabled, address dataMarketAddress, uint256 createdAt)
func (_Contract *ContractCaller) DataMarkets(opts *bind.CallOpts, arg0 common.Address) (struct {
	OwnerAddress            common.Address
	EpochSize               uint8
	SourceChainId           *big.Int
	SourceChainBlockTime    *big.Int
	UseBlockNumberAsEpochId bool
	Enabled                 bool
	DataMarketAddress       common.Address
	CreatedAt               *big.Int
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "dataMarkets", arg0)

	outstruct := new(struct {
		OwnerAddress            common.Address
		EpochSize               uint8
		SourceChainId           *big.Int
		SourceChainBlockTime    *big.Int
		UseBlockNumberAsEpochId bool
		Enabled                 bool
		DataMarketAddress       common.Address
		CreatedAt               *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.OwnerAddress = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.EpochSize = *abi.ConvertType(out[1], new(uint8)).(*uint8)
	outstruct.SourceChainId = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.SourceChainBlockTime = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.UseBlockNumberAsEpochId = *abi.ConvertType(out[4], new(bool)).(*bool)
	outstruct.Enabled = *abi.ConvertType(out[5], new(bool)).(*bool)
	outstruct.DataMarketAddress = *abi.ConvertType(out[6], new(common.Address)).(*common.Address)
	outstruct.CreatedAt = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// DataMarkets is a free data retrieval call binding the contract method 0x5f9abf70.
//
// Solidity: function dataMarkets(address ) view returns(address ownerAddress, uint8 epochSize, uint256 sourceChainId, uint256 sourceChainBlockTime, bool useBlockNumberAsEpochId, bool enabled, address dataMarketAddress, uint256 createdAt)
func (_Contract *ContractSession) DataMarkets(arg0 common.Address) (struct {
	OwnerAddress            common.Address
	EpochSize               uint8
	SourceChainId           *big.Int
	SourceChainBlockTime    *big.Int
	UseBlockNumberAsEpochId bool
	Enabled                 bool
	DataMarketAddress       common.Address
	CreatedAt               *big.Int
}, error) {
	return _Contract.Contract.DataMarkets(&_Contract.CallOpts, arg0)
}

// DataMarkets is a free data retrieval call binding the contract method 0x5f9abf70.
//
// Solidity: function dataMarkets(address ) view returns(address ownerAddress, uint8 epochSize, uint256 sourceChainId, uint256 sourceChainBlockTime, bool useBlockNumberAsEpochId, bool enabled, address dataMarketAddress, uint256 createdAt)
func (_Contract *ContractCallerSession) DataMarkets(arg0 common.Address) (struct {
	OwnerAddress            common.Address
	EpochSize               uint8
	SourceChainId           *big.Int
	SourceChainBlockTime    *big.Int
	UseBlockNumberAsEpochId bool
	Enabled                 bool
	DataMarketAddress       common.Address
	CreatedAt               *big.Int
}, error) {
	return _Contract.Contract.DataMarkets(&_Contract.CallOpts, arg0)
}

// DayCounter is a free data retrieval call binding the contract method 0x02143405.
//
// Solidity: function dayCounter(address dataMarket) view returns(uint256)
func (_Contract *ContractCaller) DayCounter(opts *bind.CallOpts, dataMarket common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "dayCounter", dataMarket)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DayCounter is a free data retrieval call binding the contract method 0x02143405.
//
// Solidity: function dayCounter(address dataMarket) view returns(uint256)
func (_Contract *ContractSession) DayCounter(dataMarket common.Address) (*big.Int, error) {
	return _Contract.Contract.DayCounter(&_Contract.CallOpts, dataMarket)
}

// DayCounter is a free data retrieval call binding the contract method 0x02143405.
//
// Solidity: function dayCounter(address dataMarket) view returns(uint256)
func (_Contract *ContractCallerSession) DayCounter(dataMarket common.Address) (*big.Int, error) {
	return _Contract.Contract.DayCounter(&_Contract.CallOpts, dataMarket)
}

// DeploymentBlockNumber is a free data retrieval call binding the contract method 0x06b8d494.
//
// Solidity: function deploymentBlockNumber(address dataMarket) view returns(uint256)
func (_Contract *ContractCaller) DeploymentBlockNumber(opts *bind.CallOpts, dataMarket common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "deploymentBlockNumber", dataMarket)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DeploymentBlockNumber is a free data retrieval call binding the contract method 0x06b8d494.
//
// Solidity: function deploymentBlockNumber(address dataMarket) view returns(uint256)
func (_Contract *ContractSession) DeploymentBlockNumber(dataMarket common.Address) (*big.Int, error) {
	return _Contract.Contract.DeploymentBlockNumber(&_Contract.CallOpts, dataMarket)
}

// DeploymentBlockNumber is a free data retrieval call binding the contract method 0x06b8d494.
//
// Solidity: function deploymentBlockNumber(address dataMarket) view returns(uint256)
func (_Contract *ContractCallerSession) DeploymentBlockNumber(dataMarket common.Address) (*big.Int, error) {
	return _Contract.Contract.DeploymentBlockNumber(&_Contract.CallOpts, dataMarket)
}

// EnabledNodeCount is a free data retrieval call binding the contract method 0xce7b6afb.
//
// Solidity: function enabledNodeCount() view returns(uint256)
func (_Contract *ContractCaller) EnabledNodeCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "enabledNodeCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EnabledNodeCount is a free data retrieval call binding the contract method 0xce7b6afb.
//
// Solidity: function enabledNodeCount() view returns(uint256)
func (_Contract *ContractSession) EnabledNodeCount() (*big.Int, error) {
	return _Contract.Contract.EnabledNodeCount(&_Contract.CallOpts)
}

// EnabledNodeCount is a free data retrieval call binding the contract method 0xce7b6afb.
//
// Solidity: function enabledNodeCount() view returns(uint256)
func (_Contract *ContractCallerSession) EnabledNodeCount() (*big.Int, error) {
	return _Contract.Contract.EnabledNodeCount(&_Contract.CallOpts)
}

// EpochIdToBatchCids is a free data retrieval call binding the contract method 0x2edb1d00.
//
// Solidity: function epochIdToBatchCids(address dataMarket, uint256 epochId) view returns(string[])
func (_Contract *ContractCaller) EpochIdToBatchCids(opts *bind.CallOpts, dataMarket common.Address, epochId *big.Int) ([]string, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "epochIdToBatchCids", dataMarket, epochId)

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// EpochIdToBatchCids is a free data retrieval call binding the contract method 0x2edb1d00.
//
// Solidity: function epochIdToBatchCids(address dataMarket, uint256 epochId) view returns(string[])
func (_Contract *ContractSession) EpochIdToBatchCids(dataMarket common.Address, epochId *big.Int) ([]string, error) {
	return _Contract.Contract.EpochIdToBatchCids(&_Contract.CallOpts, dataMarket, epochId)
}

// EpochIdToBatchCids is a free data retrieval call binding the contract method 0x2edb1d00.
//
// Solidity: function epochIdToBatchCids(address dataMarket, uint256 epochId) view returns(string[])
func (_Contract *ContractCallerSession) EpochIdToBatchCids(dataMarket common.Address, epochId *big.Int) ([]string, error) {
	return _Contract.Contract.EpochIdToBatchCids(&_Contract.CallOpts, dataMarket, epochId)
}

// EpochInfo is a free data retrieval call binding the contract method 0xc9ab0c83.
//
// Solidity: function epochInfo(address dataMarket, uint256 epochId) view returns(uint256 timestamp, uint256 blocknumber, uint256 epochEnd)
func (_Contract *ContractCaller) EpochInfo(opts *bind.CallOpts, dataMarket common.Address, epochId *big.Int) (struct {
	Timestamp   *big.Int
	Blocknumber *big.Int
	EpochEnd    *big.Int
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "epochInfo", dataMarket, epochId)

	outstruct := new(struct {
		Timestamp   *big.Int
		Blocknumber *big.Int
		EpochEnd    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Timestamp = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Blocknumber = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.EpochEnd = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// EpochInfo is a free data retrieval call binding the contract method 0xc9ab0c83.
//
// Solidity: function epochInfo(address dataMarket, uint256 epochId) view returns(uint256 timestamp, uint256 blocknumber, uint256 epochEnd)
func (_Contract *ContractSession) EpochInfo(dataMarket common.Address, epochId *big.Int) (struct {
	Timestamp   *big.Int
	Blocknumber *big.Int
	EpochEnd    *big.Int
}, error) {
	return _Contract.Contract.EpochInfo(&_Contract.CallOpts, dataMarket, epochId)
}

// EpochInfo is a free data retrieval call binding the contract method 0xc9ab0c83.
//
// Solidity: function epochInfo(address dataMarket, uint256 epochId) view returns(uint256 timestamp, uint256 blocknumber, uint256 epochEnd)
func (_Contract *ContractCallerSession) EpochInfo(dataMarket common.Address, epochId *big.Int) (struct {
	Timestamp   *big.Int
	Blocknumber *big.Int
	EpochEnd    *big.Int
}, error) {
	return _Contract.Contract.EpochInfo(&_Contract.CallOpts, dataMarket, epochId)
}

// EpochManager is a free data retrieval call binding the contract method 0x0d5a7a52.
//
// Solidity: function epochManager(address dataMarket) view returns(address)
func (_Contract *ContractCaller) EpochManager(opts *bind.CallOpts, dataMarket common.Address) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "epochManager", dataMarket)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EpochManager is a free data retrieval call binding the contract method 0x0d5a7a52.
//
// Solidity: function epochManager(address dataMarket) view returns(address)
func (_Contract *ContractSession) EpochManager(dataMarket common.Address) (common.Address, error) {
	return _Contract.Contract.EpochManager(&_Contract.CallOpts, dataMarket)
}

// EpochManager is a free data retrieval call binding the contract method 0x0d5a7a52.
//
// Solidity: function epochManager(address dataMarket) view returns(address)
func (_Contract *ContractCallerSession) EpochManager(dataMarket common.Address) (common.Address, error) {
	return _Contract.Contract.EpochManager(&_Contract.CallOpts, dataMarket)
}

// EpochsInADay is a free data retrieval call binding the contract method 0x0a1b7227.
//
// Solidity: function epochsInADay(address dataMarket) view returns(uint256)
func (_Contract *ContractCaller) EpochsInADay(opts *bind.CallOpts, dataMarket common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "epochsInADay", dataMarket)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EpochsInADay is a free data retrieval call binding the contract method 0x0a1b7227.
//
// Solidity: function epochsInADay(address dataMarket) view returns(uint256)
func (_Contract *ContractSession) EpochsInADay(dataMarket common.Address) (*big.Int, error) {
	return _Contract.Contract.EpochsInADay(&_Contract.CallOpts, dataMarket)
}

// EpochsInADay is a free data retrieval call binding the contract method 0x0a1b7227.
//
// Solidity: function epochsInADay(address dataMarket) view returns(uint256)
func (_Contract *ContractCallerSession) EpochsInADay(dataMarket common.Address) (*big.Int, error) {
	return _Contract.Contract.EpochsInADay(&_Contract.CallOpts, dataMarket)
}

// GetEpochManager is a free data retrieval call binding the contract method 0x6ce4f26c.
//
// Solidity: function getEpochManager(address dataMarket) view returns(address)
func (_Contract *ContractCaller) GetEpochManager(opts *bind.CallOpts, dataMarket common.Address) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getEpochManager", dataMarket)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetEpochManager is a free data retrieval call binding the contract method 0x6ce4f26c.
//
// Solidity: function getEpochManager(address dataMarket) view returns(address)
func (_Contract *ContractSession) GetEpochManager(dataMarket common.Address) (common.Address, error) {
	return _Contract.Contract.GetEpochManager(&_Contract.CallOpts, dataMarket)
}

// GetEpochManager is a free data retrieval call binding the contract method 0x6ce4f26c.
//
// Solidity: function getEpochManager(address dataMarket) view returns(address)
func (_Contract *ContractCallerSession) GetEpochManager(dataMarket common.Address) (common.Address, error) {
	return _Contract.Contract.GetEpochManager(&_Contract.CallOpts, dataMarket)
}

// GetSequencerId is a free data retrieval call binding the contract method 0x1c7d13a6.
//
// Solidity: function getSequencerId(address dataMarket) view returns(string)
func (_Contract *ContractCaller) GetSequencerId(opts *bind.CallOpts, dataMarket common.Address) (string, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getSequencerId", dataMarket)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetSequencerId is a free data retrieval call binding the contract method 0x1c7d13a6.
//
// Solidity: function getSequencerId(address dataMarket) view returns(string)
func (_Contract *ContractSession) GetSequencerId(dataMarket common.Address) (string, error) {
	return _Contract.Contract.GetSequencerId(&_Contract.CallOpts, dataMarket)
}

// GetSequencerId is a free data retrieval call binding the contract method 0x1c7d13a6.
//
// Solidity: function getSequencerId(address dataMarket) view returns(string)
func (_Contract *ContractCallerSession) GetSequencerId(dataMarket common.Address) (string, error) {
	return _Contract.Contract.GetSequencerId(&_Contract.CallOpts, dataMarket)
}

// GetSequencers is a free data retrieval call binding the contract method 0x6b0ad9ac.
//
// Solidity: function getSequencers(address dataMarket) view returns(address[])
func (_Contract *ContractCaller) GetSequencers(opts *bind.CallOpts, dataMarket common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getSequencers", dataMarket)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetSequencers is a free data retrieval call binding the contract method 0x6b0ad9ac.
//
// Solidity: function getSequencers(address dataMarket) view returns(address[])
func (_Contract *ContractSession) GetSequencers(dataMarket common.Address) ([]common.Address, error) {
	return _Contract.Contract.GetSequencers(&_Contract.CallOpts, dataMarket)
}

// GetSequencers is a free data retrieval call binding the contract method 0x6b0ad9ac.
//
// Solidity: function getSequencers(address dataMarket) view returns(address[])
func (_Contract *ContractCallerSession) GetSequencers(dataMarket common.Address) ([]common.Address, error) {
	return _Contract.Contract.GetSequencers(&_Contract.CallOpts, dataMarket)
}

// GetSlotInfo is a free data retrieval call binding the contract method 0xc367e244.
//
// Solidity: function getSlotInfo(address dataMarket, uint256 slotId) view returns((uint256,address,uint256,uint256))
func (_Contract *ContractCaller) GetSlotInfo(opts *bind.CallOpts, dataMarket common.Address, slotId *big.Int) (PowerloomDataMarketSlotInfo, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getSlotInfo", dataMarket, slotId)

	if err != nil {
		return *new(PowerloomDataMarketSlotInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(PowerloomDataMarketSlotInfo)).(*PowerloomDataMarketSlotInfo)

	return out0, err

}

// GetSlotInfo is a free data retrieval call binding the contract method 0xc367e244.
//
// Solidity: function getSlotInfo(address dataMarket, uint256 slotId) view returns((uint256,address,uint256,uint256))
func (_Contract *ContractSession) GetSlotInfo(dataMarket common.Address, slotId *big.Int) (PowerloomDataMarketSlotInfo, error) {
	return _Contract.Contract.GetSlotInfo(&_Contract.CallOpts, dataMarket, slotId)
}

// GetSlotInfo is a free data retrieval call binding the contract method 0xc367e244.
//
// Solidity: function getSlotInfo(address dataMarket, uint256 slotId) view returns((uint256,address,uint256,uint256))
func (_Contract *ContractCallerSession) GetSlotInfo(dataMarket common.Address, slotId *big.Int) (PowerloomDataMarketSlotInfo, error) {
	return _Contract.Contract.GetSlotInfo(&_Contract.CallOpts, dataMarket, slotId)
}

// GetSlotRewards is a free data retrieval call binding the contract method 0x9ab1013d.
//
// Solidity: function getSlotRewards(uint256 slotId) view returns(uint256 rewards)
func (_Contract *ContractCaller) GetSlotRewards(opts *bind.CallOpts, slotId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getSlotRewards", slotId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSlotRewards is a free data retrieval call binding the contract method 0x9ab1013d.
//
// Solidity: function getSlotRewards(uint256 slotId) view returns(uint256 rewards)
func (_Contract *ContractSession) GetSlotRewards(slotId *big.Int) (*big.Int, error) {
	return _Contract.Contract.GetSlotRewards(&_Contract.CallOpts, slotId)
}

// GetSlotRewards is a free data retrieval call binding the contract method 0x9ab1013d.
//
// Solidity: function getSlotRewards(uint256 slotId) view returns(uint256 rewards)
func (_Contract *ContractCallerSession) GetSlotRewards(slotId *big.Int) (*big.Int, error) {
	return _Contract.Contract.GetSlotRewards(&_Contract.CallOpts, slotId)
}

// GetTotalSequencersCount is a free data retrieval call binding the contract method 0x665ebe8c.
//
// Solidity: function getTotalSequencersCount(address dataMarket) view returns(uint256)
func (_Contract *ContractCaller) GetTotalSequencersCount(opts *bind.CallOpts, dataMarket common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getTotalSequencersCount", dataMarket)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalSequencersCount is a free data retrieval call binding the contract method 0x665ebe8c.
//
// Solidity: function getTotalSequencersCount(address dataMarket) view returns(uint256)
func (_Contract *ContractSession) GetTotalSequencersCount(dataMarket common.Address) (*big.Int, error) {
	return _Contract.Contract.GetTotalSequencersCount(&_Contract.CallOpts, dataMarket)
}

// GetTotalSequencersCount is a free data retrieval call binding the contract method 0x665ebe8c.
//
// Solidity: function getTotalSequencersCount(address dataMarket) view returns(uint256)
func (_Contract *ContractCallerSession) GetTotalSequencersCount(dataMarket common.Address) (*big.Int, error) {
	return _Contract.Contract.GetTotalSequencersCount(&_Contract.CallOpts, dataMarket)
}

// GetTotalSnapshotterCount is a free data retrieval call binding the contract method 0x92ae6f66.
//
// Solidity: function getTotalSnapshotterCount() view returns(uint256)
func (_Contract *ContractCaller) GetTotalSnapshotterCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getTotalSnapshotterCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalSnapshotterCount is a free data retrieval call binding the contract method 0x92ae6f66.
//
// Solidity: function getTotalSnapshotterCount() view returns(uint256)
func (_Contract *ContractSession) GetTotalSnapshotterCount() (*big.Int, error) {
	return _Contract.Contract.GetTotalSnapshotterCount(&_Contract.CallOpts)
}

// GetTotalSnapshotterCount is a free data retrieval call binding the contract method 0x92ae6f66.
//
// Solidity: function getTotalSnapshotterCount() view returns(uint256)
func (_Contract *ContractCallerSession) GetTotalSnapshotterCount() (*big.Int, error) {
	return _Contract.Contract.GetTotalSnapshotterCount(&_Contract.CallOpts)
}

// GetTotalValidatorsCount is a free data retrieval call binding the contract method 0xc2acc6a3.
//
// Solidity: function getTotalValidatorsCount(address dataMarket) view returns(uint256)
func (_Contract *ContractCaller) GetTotalValidatorsCount(opts *bind.CallOpts, dataMarket common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getTotalValidatorsCount", dataMarket)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalValidatorsCount is a free data retrieval call binding the contract method 0xc2acc6a3.
//
// Solidity: function getTotalValidatorsCount(address dataMarket) view returns(uint256)
func (_Contract *ContractSession) GetTotalValidatorsCount(dataMarket common.Address) (*big.Int, error) {
	return _Contract.Contract.GetTotalValidatorsCount(&_Contract.CallOpts, dataMarket)
}

// GetTotalValidatorsCount is a free data retrieval call binding the contract method 0xc2acc6a3.
//
// Solidity: function getTotalValidatorsCount(address dataMarket) view returns(uint256)
func (_Contract *ContractCallerSession) GetTotalValidatorsCount(dataMarket common.Address) (*big.Int, error) {
	return _Contract.Contract.GetTotalValidatorsCount(&_Contract.CallOpts, dataMarket)
}

// GetValidators is a free data retrieval call binding the contract method 0xff8744a6.
//
// Solidity: function getValidators(address dataMarket) view returns(address[])
func (_Contract *ContractCaller) GetValidators(opts *bind.CallOpts, dataMarket common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getValidators", dataMarket)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetValidators is a free data retrieval call binding the contract method 0xff8744a6.
//
// Solidity: function getValidators(address dataMarket) view returns(address[])
func (_Contract *ContractSession) GetValidators(dataMarket common.Address) ([]common.Address, error) {
	return _Contract.Contract.GetValidators(&_Contract.CallOpts, dataMarket)
}

// GetValidators is a free data retrieval call binding the contract method 0xff8744a6.
//
// Solidity: function getValidators(address dataMarket) view returns(address[])
func (_Contract *ContractCallerSession) GetValidators(dataMarket common.Address) ([]common.Address, error) {
	return _Contract.Contract.GetValidators(&_Contract.CallOpts, dataMarket)
}

// LastFinalizedSnapshot is a free data retrieval call binding the contract method 0xded2465b.
//
// Solidity: function lastFinalizedSnapshot(address dataMarket, string projectId) view returns(uint256)
func (_Contract *ContractCaller) LastFinalizedSnapshot(opts *bind.CallOpts, dataMarket common.Address, projectId string) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "lastFinalizedSnapshot", dataMarket, projectId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastFinalizedSnapshot is a free data retrieval call binding the contract method 0xded2465b.
//
// Solidity: function lastFinalizedSnapshot(address dataMarket, string projectId) view returns(uint256)
func (_Contract *ContractSession) LastFinalizedSnapshot(dataMarket common.Address, projectId string) (*big.Int, error) {
	return _Contract.Contract.LastFinalizedSnapshot(&_Contract.CallOpts, dataMarket, projectId)
}

// LastFinalizedSnapshot is a free data retrieval call binding the contract method 0xded2465b.
//
// Solidity: function lastFinalizedSnapshot(address dataMarket, string projectId) view returns(uint256)
func (_Contract *ContractCallerSession) LastFinalizedSnapshot(dataMarket common.Address, projectId string) (*big.Int, error) {
	return _Contract.Contract.LastFinalizedSnapshot(&_Contract.CallOpts, dataMarket, projectId)
}

// LastSequencerFinalizedSnapshot is a free data retrieval call binding the contract method 0x13be7391.
//
// Solidity: function lastSequencerFinalizedSnapshot(address dataMarket, string projectId) view returns(uint256)
func (_Contract *ContractCaller) LastSequencerFinalizedSnapshot(opts *bind.CallOpts, dataMarket common.Address, projectId string) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "lastSequencerFinalizedSnapshot", dataMarket, projectId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastSequencerFinalizedSnapshot is a free data retrieval call binding the contract method 0x13be7391.
//
// Solidity: function lastSequencerFinalizedSnapshot(address dataMarket, string projectId) view returns(uint256)
func (_Contract *ContractSession) LastSequencerFinalizedSnapshot(dataMarket common.Address, projectId string) (*big.Int, error) {
	return _Contract.Contract.LastSequencerFinalizedSnapshot(&_Contract.CallOpts, dataMarket, projectId)
}

// LastSequencerFinalizedSnapshot is a free data retrieval call binding the contract method 0x13be7391.
//
// Solidity: function lastSequencerFinalizedSnapshot(address dataMarket, string projectId) view returns(uint256)
func (_Contract *ContractCallerSession) LastSequencerFinalizedSnapshot(dataMarket common.Address, projectId string) (*big.Int, error) {
	return _Contract.Contract.LastSequencerFinalizedSnapshot(&_Contract.CallOpts, dataMarket, projectId)
}

// MaxAttestationFinalizedRootHash is a free data retrieval call binding the contract method 0xb1b65bd9.
//
// Solidity: function maxAttestationFinalizedRootHash(address dataMarket, string batchCid) view returns(bytes32)
func (_Contract *ContractCaller) MaxAttestationFinalizedRootHash(opts *bind.CallOpts, dataMarket common.Address, batchCid string) ([32]byte, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "maxAttestationFinalizedRootHash", dataMarket, batchCid)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MaxAttestationFinalizedRootHash is a free data retrieval call binding the contract method 0xb1b65bd9.
//
// Solidity: function maxAttestationFinalizedRootHash(address dataMarket, string batchCid) view returns(bytes32)
func (_Contract *ContractSession) MaxAttestationFinalizedRootHash(dataMarket common.Address, batchCid string) ([32]byte, error) {
	return _Contract.Contract.MaxAttestationFinalizedRootHash(&_Contract.CallOpts, dataMarket, batchCid)
}

// MaxAttestationFinalizedRootHash is a free data retrieval call binding the contract method 0xb1b65bd9.
//
// Solidity: function maxAttestationFinalizedRootHash(address dataMarket, string batchCid) view returns(bytes32)
func (_Contract *ContractCallerSession) MaxAttestationFinalizedRootHash(dataMarket common.Address, batchCid string) ([32]byte, error) {
	return _Contract.Contract.MaxAttestationFinalizedRootHash(&_Contract.CallOpts, dataMarket, batchCid)
}

// MaxAttestationsCount is a free data retrieval call binding the contract method 0xc9f1c8fe.
//
// Solidity: function maxAttestationsCount(address dataMarket, string batchCid) view returns(uint256)
func (_Contract *ContractCaller) MaxAttestationsCount(opts *bind.CallOpts, dataMarket common.Address, batchCid string) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "maxAttestationsCount", dataMarket, batchCid)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxAttestationsCount is a free data retrieval call binding the contract method 0xc9f1c8fe.
//
// Solidity: function maxAttestationsCount(address dataMarket, string batchCid) view returns(uint256)
func (_Contract *ContractSession) MaxAttestationsCount(dataMarket common.Address, batchCid string) (*big.Int, error) {
	return _Contract.Contract.MaxAttestationsCount(&_Contract.CallOpts, dataMarket, batchCid)
}

// MaxAttestationsCount is a free data retrieval call binding the contract method 0xc9f1c8fe.
//
// Solidity: function maxAttestationsCount(address dataMarket, string batchCid) view returns(uint256)
func (_Contract *ContractCallerSession) MaxAttestationsCount(dataMarket common.Address, batchCid string) (*big.Int, error) {
	return _Contract.Contract.MaxAttestationsCount(&_Contract.CallOpts, dataMarket, batchCid)
}

// MaxSnapshotsCid is a free data retrieval call binding the contract method 0x7e9ce892.
//
// Solidity: function maxSnapshotsCid(address dataMarket, string projectId, uint256 epochId) view returns(string)
func (_Contract *ContractCaller) MaxSnapshotsCid(opts *bind.CallOpts, dataMarket common.Address, projectId string, epochId *big.Int) (string, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "maxSnapshotsCid", dataMarket, projectId, epochId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// MaxSnapshotsCid is a free data retrieval call binding the contract method 0x7e9ce892.
//
// Solidity: function maxSnapshotsCid(address dataMarket, string projectId, uint256 epochId) view returns(string)
func (_Contract *ContractSession) MaxSnapshotsCid(dataMarket common.Address, projectId string, epochId *big.Int) (string, error) {
	return _Contract.Contract.MaxSnapshotsCid(&_Contract.CallOpts, dataMarket, projectId, epochId)
}

// MaxSnapshotsCid is a free data retrieval call binding the contract method 0x7e9ce892.
//
// Solidity: function maxSnapshotsCid(address dataMarket, string projectId, uint256 epochId) view returns(string)
func (_Contract *ContractCallerSession) MaxSnapshotsCid(dataMarket common.Address, projectId string, epochId *big.Int) (string, error) {
	return _Contract.Contract.MaxSnapshotsCid(&_Contract.CallOpts, dataMarket, projectId, epochId)
}

// MinAttestationsForConsensus is a free data retrieval call binding the contract method 0xb3d95efa.
//
// Solidity: function minAttestationsForConsensus(address dataMarket) view returns(uint256)
func (_Contract *ContractCaller) MinAttestationsForConsensus(opts *bind.CallOpts, dataMarket common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "minAttestationsForConsensus", dataMarket)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinAttestationsForConsensus is a free data retrieval call binding the contract method 0xb3d95efa.
//
// Solidity: function minAttestationsForConsensus(address dataMarket) view returns(uint256)
func (_Contract *ContractSession) MinAttestationsForConsensus(dataMarket common.Address) (*big.Int, error) {
	return _Contract.Contract.MinAttestationsForConsensus(&_Contract.CallOpts, dataMarket)
}

// MinAttestationsForConsensus is a free data retrieval call binding the contract method 0xb3d95efa.
//
// Solidity: function minAttestationsForConsensus(address dataMarket) view returns(uint256)
func (_Contract *ContractCallerSession) MinAttestationsForConsensus(dataMarket common.Address) (*big.Int, error) {
	return _Contract.Contract.MinAttestationsForConsensus(&_Contract.CallOpts, dataMarket)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractSession) Owner() (common.Address, error) {
	return _Contract.Contract.Owner(&_Contract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contract *ContractCallerSession) Owner() (common.Address, error) {
	return _Contract.Contract.Owner(&_Contract.CallOpts)
}

// ProjectFirstEpochId is a free data retrieval call binding the contract method 0x4242614c.
//
// Solidity: function projectFirstEpochId(address dataMarket, string projectId) view returns(uint256)
func (_Contract *ContractCaller) ProjectFirstEpochId(opts *bind.CallOpts, dataMarket common.Address, projectId string) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "projectFirstEpochId", dataMarket, projectId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProjectFirstEpochId is a free data retrieval call binding the contract method 0x4242614c.
//
// Solidity: function projectFirstEpochId(address dataMarket, string projectId) view returns(uint256)
func (_Contract *ContractSession) ProjectFirstEpochId(dataMarket common.Address, projectId string) (*big.Int, error) {
	return _Contract.Contract.ProjectFirstEpochId(&_Contract.CallOpts, dataMarket, projectId)
}

// ProjectFirstEpochId is a free data retrieval call binding the contract method 0x4242614c.
//
// Solidity: function projectFirstEpochId(address dataMarket, string projectId) view returns(uint256)
func (_Contract *ContractCallerSession) ProjectFirstEpochId(dataMarket common.Address, projectId string) (*big.Int, error) {
	return _Contract.Contract.ProjectFirstEpochId(&_Contract.CallOpts, dataMarket, projectId)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Contract *ContractCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Contract *ContractSession) ProxiableUUID() ([32]byte, error) {
	return _Contract.Contract.ProxiableUUID(&_Contract.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Contract *ContractCallerSession) ProxiableUUID() ([32]byte, error) {
	return _Contract.Contract.ProxiableUUID(&_Contract.CallOpts)
}

// RewardPoolSize is a free data retrieval call binding the contract method 0x5350fa81.
//
// Solidity: function rewardPoolSize(address dataMarket) view returns(uint256)
func (_Contract *ContractCaller) RewardPoolSize(opts *bind.CallOpts, dataMarket common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "rewardPoolSize", dataMarket)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardPoolSize is a free data retrieval call binding the contract method 0x5350fa81.
//
// Solidity: function rewardPoolSize(address dataMarket) view returns(uint256)
func (_Contract *ContractSession) RewardPoolSize(dataMarket common.Address) (*big.Int, error) {
	return _Contract.Contract.RewardPoolSize(&_Contract.CallOpts, dataMarket)
}

// RewardPoolSize is a free data retrieval call binding the contract method 0x5350fa81.
//
// Solidity: function rewardPoolSize(address dataMarket) view returns(uint256)
func (_Contract *ContractCallerSession) RewardPoolSize(dataMarket common.Address) (*big.Int, error) {
	return _Contract.Contract.RewardPoolSize(&_Contract.CallOpts, dataMarket)
}

// RewardsEnabled is a free data retrieval call binding the contract method 0x83450d26.
//
// Solidity: function rewardsEnabled(address dataMarket) view returns(bool)
func (_Contract *ContractCaller) RewardsEnabled(opts *bind.CallOpts, dataMarket common.Address) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "rewardsEnabled", dataMarket)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// RewardsEnabled is a free data retrieval call binding the contract method 0x83450d26.
//
// Solidity: function rewardsEnabled(address dataMarket) view returns(bool)
func (_Contract *ContractSession) RewardsEnabled(dataMarket common.Address) (bool, error) {
	return _Contract.Contract.RewardsEnabled(&_Contract.CallOpts, dataMarket)
}

// RewardsEnabled is a free data retrieval call binding the contract method 0x83450d26.
//
// Solidity: function rewardsEnabled(address dataMarket) view returns(bool)
func (_Contract *ContractCallerSession) RewardsEnabled(dataMarket common.Address) (bool, error) {
	return _Contract.Contract.RewardsEnabled(&_Contract.CallOpts, dataMarket)
}

// SlotRewardPoints is a free data retrieval call binding the contract method 0x9a2458a6.
//
// Solidity: function slotRewardPoints(address dataMarket, uint256 slotId) view returns(uint256)
func (_Contract *ContractCaller) SlotRewardPoints(opts *bind.CallOpts, dataMarket common.Address, slotId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "slotRewardPoints", dataMarket, slotId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SlotRewardPoints is a free data retrieval call binding the contract method 0x9a2458a6.
//
// Solidity: function slotRewardPoints(address dataMarket, uint256 slotId) view returns(uint256)
func (_Contract *ContractSession) SlotRewardPoints(dataMarket common.Address, slotId *big.Int) (*big.Int, error) {
	return _Contract.Contract.SlotRewardPoints(&_Contract.CallOpts, dataMarket, slotId)
}

// SlotRewardPoints is a free data retrieval call binding the contract method 0x9a2458a6.
//
// Solidity: function slotRewardPoints(address dataMarket, uint256 slotId) view returns(uint256)
func (_Contract *ContractCallerSession) SlotRewardPoints(dataMarket common.Address, slotId *big.Int) (*big.Int, error) {
	return _Contract.Contract.SlotRewardPoints(&_Contract.CallOpts, dataMarket, slotId)
}

// SlotRewards is a free data retrieval call binding the contract method 0x53e28269.
//
// Solidity: function slotRewards(uint256 ) view returns(uint256)
func (_Contract *ContractCaller) SlotRewards(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "slotRewards", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SlotRewards is a free data retrieval call binding the contract method 0x53e28269.
//
// Solidity: function slotRewards(uint256 ) view returns(uint256)
func (_Contract *ContractSession) SlotRewards(arg0 *big.Int) (*big.Int, error) {
	return _Contract.Contract.SlotRewards(&_Contract.CallOpts, arg0)
}

// SlotRewards is a free data retrieval call binding the contract method 0x53e28269.
//
// Solidity: function slotRewards(uint256 ) view returns(uint256)
func (_Contract *ContractCallerSession) SlotRewards(arg0 *big.Int) (*big.Int, error) {
	return _Contract.Contract.SlotRewards(&_Contract.CallOpts, arg0)
}

// SlotSnapshotterMapping is a free data retrieval call binding the contract method 0x948a463e.
//
// Solidity: function slotSnapshotterMapping(uint256 slotId) view returns(address)
func (_Contract *ContractCaller) SlotSnapshotterMapping(opts *bind.CallOpts, slotId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "slotSnapshotterMapping", slotId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SlotSnapshotterMapping is a free data retrieval call binding the contract method 0x948a463e.
//
// Solidity: function slotSnapshotterMapping(uint256 slotId) view returns(address)
func (_Contract *ContractSession) SlotSnapshotterMapping(slotId *big.Int) (common.Address, error) {
	return _Contract.Contract.SlotSnapshotterMapping(&_Contract.CallOpts, slotId)
}

// SlotSnapshotterMapping is a free data retrieval call binding the contract method 0x948a463e.
//
// Solidity: function slotSnapshotterMapping(uint256 slotId) view returns(address)
func (_Contract *ContractCallerSession) SlotSnapshotterMapping(slotId *big.Int) (common.Address, error) {
	return _Contract.Contract.SlotSnapshotterMapping(&_Contract.CallOpts, slotId)
}

// SlotSubmissionCount is a free data retrieval call binding the contract method 0x7f9ee950.
//
// Solidity: function slotSubmissionCount(address dataMarket, uint256 slotId, uint256 dayId) view returns(uint256)
func (_Contract *ContractCaller) SlotSubmissionCount(opts *bind.CallOpts, dataMarket common.Address, slotId *big.Int, dayId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "slotSubmissionCount", dataMarket, slotId, dayId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SlotSubmissionCount is a free data retrieval call binding the contract method 0x7f9ee950.
//
// Solidity: function slotSubmissionCount(address dataMarket, uint256 slotId, uint256 dayId) view returns(uint256)
func (_Contract *ContractSession) SlotSubmissionCount(dataMarket common.Address, slotId *big.Int, dayId *big.Int) (*big.Int, error) {
	return _Contract.Contract.SlotSubmissionCount(&_Contract.CallOpts, dataMarket, slotId, dayId)
}

// SlotSubmissionCount is a free data retrieval call binding the contract method 0x7f9ee950.
//
// Solidity: function slotSubmissionCount(address dataMarket, uint256 slotId, uint256 dayId) view returns(uint256)
func (_Contract *ContractCallerSession) SlotSubmissionCount(dataMarket common.Address, slotId *big.Int, dayId *big.Int) (*big.Int, error) {
	return _Contract.Contract.SlotSubmissionCount(&_Contract.CallOpts, dataMarket, slotId, dayId)
}

// SnapshotStatus is a free data retrieval call binding the contract method 0x1f588588.
//
// Solidity: function snapshotStatus(address dataMarket, string projectId, uint256 epochId) view returns(uint8 status, string snapshotCid, uint256 timestamp)
func (_Contract *ContractCaller) SnapshotStatus(opts *bind.CallOpts, dataMarket common.Address, projectId string, epochId *big.Int) (struct {
	Status      uint8
	SnapshotCid string
	Timestamp   *big.Int
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "snapshotStatus", dataMarket, projectId, epochId)

	outstruct := new(struct {
		Status      uint8
		SnapshotCid string
		Timestamp   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Status = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.SnapshotCid = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Timestamp = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// SnapshotStatus is a free data retrieval call binding the contract method 0x1f588588.
//
// Solidity: function snapshotStatus(address dataMarket, string projectId, uint256 epochId) view returns(uint8 status, string snapshotCid, uint256 timestamp)
func (_Contract *ContractSession) SnapshotStatus(dataMarket common.Address, projectId string, epochId *big.Int) (struct {
	Status      uint8
	SnapshotCid string
	Timestamp   *big.Int
}, error) {
	return _Contract.Contract.SnapshotStatus(&_Contract.CallOpts, dataMarket, projectId, epochId)
}

// SnapshotStatus is a free data retrieval call binding the contract method 0x1f588588.
//
// Solidity: function snapshotStatus(address dataMarket, string projectId, uint256 epochId) view returns(uint8 status, string snapshotCid, uint256 timestamp)
func (_Contract *ContractCallerSession) SnapshotStatus(dataMarket common.Address, projectId string, epochId *big.Int) (struct {
	Status      uint8
	SnapshotCid string
	Timestamp   *big.Int
}, error) {
	return _Contract.Contract.SnapshotStatus(&_Contract.CallOpts, dataMarket, projectId, epochId)
}

// SnapshotSubmissionWindow is a free data retrieval call binding the contract method 0xf3354db0.
//
// Solidity: function snapshotSubmissionWindow(address dataMarket) view returns(uint256)
func (_Contract *ContractCaller) SnapshotSubmissionWindow(opts *bind.CallOpts, dataMarket common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "snapshotSubmissionWindow", dataMarket)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SnapshotSubmissionWindow is a free data retrieval call binding the contract method 0xf3354db0.
//
// Solidity: function snapshotSubmissionWindow(address dataMarket) view returns(uint256)
func (_Contract *ContractSession) SnapshotSubmissionWindow(dataMarket common.Address) (*big.Int, error) {
	return _Contract.Contract.SnapshotSubmissionWindow(&_Contract.CallOpts, dataMarket)
}

// SnapshotSubmissionWindow is a free data retrieval call binding the contract method 0xf3354db0.
//
// Solidity: function snapshotSubmissionWindow(address dataMarket) view returns(uint256)
func (_Contract *ContractCallerSession) SnapshotSubmissionWindow(dataMarket common.Address) (*big.Int, error) {
	return _Contract.Contract.SnapshotSubmissionWindow(&_Contract.CallOpts, dataMarket)
}

// SnapshotterState is a free data retrieval call binding the contract method 0x342050cc.
//
// Solidity: function snapshotterState() view returns(address)
func (_Contract *ContractCaller) SnapshotterState(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "snapshotterState")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SnapshotterState is a free data retrieval call binding the contract method 0x342050cc.
//
// Solidity: function snapshotterState() view returns(address)
func (_Contract *ContractSession) SnapshotterState() (common.Address, error) {
	return _Contract.Contract.SnapshotterState(&_Contract.CallOpts)
}

// SnapshotterState is a free data retrieval call binding the contract method 0x342050cc.
//
// Solidity: function snapshotterState() view returns(address)
func (_Contract *ContractCallerSession) SnapshotterState() (common.Address, error) {
	return _Contract.Contract.SnapshotterState(&_Contract.CallOpts)
}

// UserInfo is a free data retrieval call binding the contract method 0x1959a002.
//
// Solidity: function userInfo(address ) view returns(uint256 totalRewards, uint256 totalClaimed, uint256 lastClaimed, uint256 lastUpdated)
func (_Contract *ContractCaller) UserInfo(opts *bind.CallOpts, arg0 common.Address) (struct {
	TotalRewards *big.Int
	TotalClaimed *big.Int
	LastClaimed  *big.Int
	LastUpdated  *big.Int
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "userInfo", arg0)

	outstruct := new(struct {
		TotalRewards *big.Int
		TotalClaimed *big.Int
		LastClaimed  *big.Int
		LastUpdated  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TotalRewards = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.TotalClaimed = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.LastClaimed = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.LastUpdated = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// UserInfo is a free data retrieval call binding the contract method 0x1959a002.
//
// Solidity: function userInfo(address ) view returns(uint256 totalRewards, uint256 totalClaimed, uint256 lastClaimed, uint256 lastUpdated)
func (_Contract *ContractSession) UserInfo(arg0 common.Address) (struct {
	TotalRewards *big.Int
	TotalClaimed *big.Int
	LastClaimed  *big.Int
	LastUpdated  *big.Int
}, error) {
	return _Contract.Contract.UserInfo(&_Contract.CallOpts, arg0)
}

// UserInfo is a free data retrieval call binding the contract method 0x1959a002.
//
// Solidity: function userInfo(address ) view returns(uint256 totalRewards, uint256 totalClaimed, uint256 lastClaimed, uint256 lastUpdated)
func (_Contract *ContractCallerSession) UserInfo(arg0 common.Address) (struct {
	TotalRewards *big.Int
	TotalClaimed *big.Int
	LastClaimed  *big.Int
	LastUpdated  *big.Int
}, error) {
	return _Contract.Contract.UserInfo(&_Contract.CallOpts, arg0)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0xef5cfb8c.
//
// Solidity: function claimRewards(address _user) returns()
func (_Contract *ContractTransactor) ClaimRewards(opts *bind.TransactOpts, _user common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "claimRewards", _user)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0xef5cfb8c.
//
// Solidity: function claimRewards(address _user) returns()
func (_Contract *ContractSession) ClaimRewards(_user common.Address) (*types.Transaction, error) {
	return _Contract.Contract.ClaimRewards(&_Contract.TransactOpts, _user)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0xef5cfb8c.
//
// Solidity: function claimRewards(address _user) returns()
func (_Contract *ContractTransactorSession) ClaimRewards(_user common.Address) (*types.Transaction, error) {
	return _Contract.Contract.ClaimRewards(&_Contract.TransactOpts, _user)
}

// CreateDataMarket is a paid mutator transaction binding the contract method 0x1dbc586b.
//
// Solidity: function createDataMarket(address ownerAddress, uint8 epochSize, uint256 sourceChainId, uint256 sourceChainBlockTime, bool useBlockNumberAsEpochId) returns(address)
func (_Contract *ContractTransactor) CreateDataMarket(opts *bind.TransactOpts, ownerAddress common.Address, epochSize uint8, sourceChainId *big.Int, sourceChainBlockTime *big.Int, useBlockNumberAsEpochId bool) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "createDataMarket", ownerAddress, epochSize, sourceChainId, sourceChainBlockTime, useBlockNumberAsEpochId)
}

// CreateDataMarket is a paid mutator transaction binding the contract method 0x1dbc586b.
//
// Solidity: function createDataMarket(address ownerAddress, uint8 epochSize, uint256 sourceChainId, uint256 sourceChainBlockTime, bool useBlockNumberAsEpochId) returns(address)
func (_Contract *ContractSession) CreateDataMarket(ownerAddress common.Address, epochSize uint8, sourceChainId *big.Int, sourceChainBlockTime *big.Int, useBlockNumberAsEpochId bool) (*types.Transaction, error) {
	return _Contract.Contract.CreateDataMarket(&_Contract.TransactOpts, ownerAddress, epochSize, sourceChainId, sourceChainBlockTime, useBlockNumberAsEpochId)
}

// CreateDataMarket is a paid mutator transaction binding the contract method 0x1dbc586b.
//
// Solidity: function createDataMarket(address ownerAddress, uint8 epochSize, uint256 sourceChainId, uint256 sourceChainBlockTime, bool useBlockNumberAsEpochId) returns(address)
func (_Contract *ContractTransactorSession) CreateDataMarket(ownerAddress common.Address, epochSize uint8, sourceChainId *big.Int, sourceChainBlockTime *big.Int, useBlockNumberAsEpochId bool) (*types.Transaction, error) {
	return _Contract.Contract.CreateDataMarket(&_Contract.TransactOpts, ownerAddress, epochSize, sourceChainId, sourceChainBlockTime, useBlockNumberAsEpochId)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0xdb2e21bc.
//
// Solidity: function emergencyWithdraw() returns()
func (_Contract *ContractTransactor) EmergencyWithdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "emergencyWithdraw")
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0xdb2e21bc.
//
// Solidity: function emergencyWithdraw() returns()
func (_Contract *ContractSession) EmergencyWithdraw() (*types.Transaction, error) {
	return _Contract.Contract.EmergencyWithdraw(&_Contract.TransactOpts)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0xdb2e21bc.
//
// Solidity: function emergencyWithdraw() returns()
func (_Contract *ContractTransactorSession) EmergencyWithdraw() (*types.Transaction, error) {
	return _Contract.Contract.EmergencyWithdraw(&_Contract.TransactOpts)
}

// EndBatchSubmissions is a paid mutator transaction binding the contract method 0x6ee55d73.
//
// Solidity: function endBatchSubmissions(address dataMarket, uint256 epochId) returns()
func (_Contract *ContractTransactor) EndBatchSubmissions(opts *bind.TransactOpts, dataMarket common.Address, epochId *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "endBatchSubmissions", dataMarket, epochId)
}

// EndBatchSubmissions is a paid mutator transaction binding the contract method 0x6ee55d73.
//
// Solidity: function endBatchSubmissions(address dataMarket, uint256 epochId) returns()
func (_Contract *ContractSession) EndBatchSubmissions(dataMarket common.Address, epochId *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.EndBatchSubmissions(&_Contract.TransactOpts, dataMarket, epochId)
}

// EndBatchSubmissions is a paid mutator transaction binding the contract method 0x6ee55d73.
//
// Solidity: function endBatchSubmissions(address dataMarket, uint256 epochId) returns()
func (_Contract *ContractTransactorSession) EndBatchSubmissions(dataMarket common.Address, epochId *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.EndBatchSubmissions(&_Contract.TransactOpts, dataMarket, epochId)
}

// ForceCompleteConsensusAttestations is a paid mutator transaction binding the contract method 0xff614fff.
//
// Solidity: function forceCompleteConsensusAttestations(address dataMarket, string batchCid, uint256 epochId) returns()
func (_Contract *ContractTransactor) ForceCompleteConsensusAttestations(opts *bind.TransactOpts, dataMarket common.Address, batchCid string, epochId *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "forceCompleteConsensusAttestations", dataMarket, batchCid, epochId)
}

// ForceCompleteConsensusAttestations is a paid mutator transaction binding the contract method 0xff614fff.
//
// Solidity: function forceCompleteConsensusAttestations(address dataMarket, string batchCid, uint256 epochId) returns()
func (_Contract *ContractSession) ForceCompleteConsensusAttestations(dataMarket common.Address, batchCid string, epochId *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.ForceCompleteConsensusAttestations(&_Contract.TransactOpts, dataMarket, batchCid, epochId)
}

// ForceCompleteConsensusAttestations is a paid mutator transaction binding the contract method 0xff614fff.
//
// Solidity: function forceCompleteConsensusAttestations(address dataMarket, string batchCid, uint256 epochId) returns()
func (_Contract *ContractTransactorSession) ForceCompleteConsensusAttestations(dataMarket common.Address, batchCid string, epochId *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.ForceCompleteConsensusAttestations(&_Contract.TransactOpts, dataMarket, batchCid, epochId)
}

// ForceSkipEpoch is a paid mutator transaction binding the contract method 0x27856ff3.
//
// Solidity: function forceSkipEpoch(address dataMarket, uint256 begin, uint256 end) returns()
func (_Contract *ContractTransactor) ForceSkipEpoch(opts *bind.TransactOpts, dataMarket common.Address, begin *big.Int, end *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "forceSkipEpoch", dataMarket, begin, end)
}

// ForceSkipEpoch is a paid mutator transaction binding the contract method 0x27856ff3.
//
// Solidity: function forceSkipEpoch(address dataMarket, uint256 begin, uint256 end) returns()
func (_Contract *ContractSession) ForceSkipEpoch(dataMarket common.Address, begin *big.Int, end *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.ForceSkipEpoch(&_Contract.TransactOpts, dataMarket, begin, end)
}

// ForceSkipEpoch is a paid mutator transaction binding the contract method 0x27856ff3.
//
// Solidity: function forceSkipEpoch(address dataMarket, uint256 begin, uint256 end) returns()
func (_Contract *ContractTransactorSession) ForceSkipEpoch(dataMarket common.Address, begin *big.Int, end *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.ForceSkipEpoch(&_Contract.TransactOpts, dataMarket, begin, end)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address initialOwner) returns()
func (_Contract *ContractTransactor) Initialize(opts *bind.TransactOpts, initialOwner common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "initialize", initialOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address initialOwner) returns()
func (_Contract *ContractSession) Initialize(initialOwner common.Address) (*types.Transaction, error) {
	return _Contract.Contract.Initialize(&_Contract.TransactOpts, initialOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address initialOwner) returns()
func (_Contract *ContractTransactorSession) Initialize(initialOwner common.Address) (*types.Transaction, error) {
	return _Contract.Contract.Initialize(&_Contract.TransactOpts, initialOwner)
}

// LoadCurrentDay is a paid mutator transaction binding the contract method 0x59c60b30.
//
// Solidity: function loadCurrentDay(address dataMarket, uint256 _dayCounter) returns()
func (_Contract *ContractTransactor) LoadCurrentDay(opts *bind.TransactOpts, dataMarket common.Address, _dayCounter *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "loadCurrentDay", dataMarket, _dayCounter)
}

// LoadCurrentDay is a paid mutator transaction binding the contract method 0x59c60b30.
//
// Solidity: function loadCurrentDay(address dataMarket, uint256 _dayCounter) returns()
func (_Contract *ContractSession) LoadCurrentDay(dataMarket common.Address, _dayCounter *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.LoadCurrentDay(&_Contract.TransactOpts, dataMarket, _dayCounter)
}

// LoadCurrentDay is a paid mutator transaction binding the contract method 0x59c60b30.
//
// Solidity: function loadCurrentDay(address dataMarket, uint256 _dayCounter) returns()
func (_Contract *ContractTransactorSession) LoadCurrentDay(dataMarket common.Address, _dayCounter *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.LoadCurrentDay(&_Contract.TransactOpts, dataMarket, _dayCounter)
}

// LoadSlotSubmissions is a paid mutator transaction binding the contract method 0xa1ac8af0.
//
// Solidity: function loadSlotSubmissions(address dataMarket, uint256 slotId, uint256 dayId, uint256 snapshotCount) returns()
func (_Contract *ContractTransactor) LoadSlotSubmissions(opts *bind.TransactOpts, dataMarket common.Address, slotId *big.Int, dayId *big.Int, snapshotCount *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "loadSlotSubmissions", dataMarket, slotId, dayId, snapshotCount)
}

// LoadSlotSubmissions is a paid mutator transaction binding the contract method 0xa1ac8af0.
//
// Solidity: function loadSlotSubmissions(address dataMarket, uint256 slotId, uint256 dayId, uint256 snapshotCount) returns()
func (_Contract *ContractSession) LoadSlotSubmissions(dataMarket common.Address, slotId *big.Int, dayId *big.Int, snapshotCount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.LoadSlotSubmissions(&_Contract.TransactOpts, dataMarket, slotId, dayId, snapshotCount)
}

// LoadSlotSubmissions is a paid mutator transaction binding the contract method 0xa1ac8af0.
//
// Solidity: function loadSlotSubmissions(address dataMarket, uint256 slotId, uint256 dayId, uint256 snapshotCount) returns()
func (_Contract *ContractTransactorSession) LoadSlotSubmissions(dataMarket common.Address, slotId *big.Int, dayId *big.Int, snapshotCount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.LoadSlotSubmissions(&_Contract.TransactOpts, dataMarket, slotId, dayId, snapshotCount)
}

// ReleaseEpoch is a paid mutator transaction binding the contract method 0x09517cb6.
//
// Solidity: function releaseEpoch(address dataMarket, uint256 begin, uint256 end) returns()
func (_Contract *ContractTransactor) ReleaseEpoch(opts *bind.TransactOpts, dataMarket common.Address, begin *big.Int, end *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "releaseEpoch", dataMarket, begin, end)
}

// ReleaseEpoch is a paid mutator transaction binding the contract method 0x09517cb6.
//
// Solidity: function releaseEpoch(address dataMarket, uint256 begin, uint256 end) returns()
func (_Contract *ContractSession) ReleaseEpoch(dataMarket common.Address, begin *big.Int, end *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.ReleaseEpoch(&_Contract.TransactOpts, dataMarket, begin, end)
}

// ReleaseEpoch is a paid mutator transaction binding the contract method 0x09517cb6.
//
// Solidity: function releaseEpoch(address dataMarket, uint256 begin, uint256 end) returns()
func (_Contract *ContractTransactorSession) ReleaseEpoch(dataMarket common.Address, begin *big.Int, end *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.ReleaseEpoch(&_Contract.TransactOpts, dataMarket, begin, end)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contract *ContractTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contract *ContractSession) RenounceOwnership() (*types.Transaction, error) {
	return _Contract.Contract.RenounceOwnership(&_Contract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Contract *ContractTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Contract.Contract.RenounceOwnership(&_Contract.TransactOpts)
}

// SetSequencerId is a paid mutator transaction binding the contract method 0x79d01175.
//
// Solidity: function setSequencerId(address dataMarket, string _sequencerId) returns()
func (_Contract *ContractTransactor) SetSequencerId(opts *bind.TransactOpts, dataMarket common.Address, _sequencerId string) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setSequencerId", dataMarket, _sequencerId)
}

// SetSequencerId is a paid mutator transaction binding the contract method 0x79d01175.
//
// Solidity: function setSequencerId(address dataMarket, string _sequencerId) returns()
func (_Contract *ContractSession) SetSequencerId(dataMarket common.Address, _sequencerId string) (*types.Transaction, error) {
	return _Contract.Contract.SetSequencerId(&_Contract.TransactOpts, dataMarket, _sequencerId)
}

// SetSequencerId is a paid mutator transaction binding the contract method 0x79d01175.
//
// Solidity: function setSequencerId(address dataMarket, string _sequencerId) returns()
func (_Contract *ContractTransactorSession) SetSequencerId(dataMarket common.Address, _sequencerId string) (*types.Transaction, error) {
	return _Contract.Contract.SetSequencerId(&_Contract.TransactOpts, dataMarket, _sequencerId)
}

// SubmitBatchAttestation is a paid mutator transaction binding the contract method 0xcd4c1a34.
//
// Solidity: function submitBatchAttestation(address dataMarket, string batchCid, uint256 epochId, bytes32 finalizedCidsRootHash) returns()
func (_Contract *ContractTransactor) SubmitBatchAttestation(opts *bind.TransactOpts, dataMarket common.Address, batchCid string, epochId *big.Int, finalizedCidsRootHash [32]byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "submitBatchAttestation", dataMarket, batchCid, epochId, finalizedCidsRootHash)
}

// SubmitBatchAttestation is a paid mutator transaction binding the contract method 0xcd4c1a34.
//
// Solidity: function submitBatchAttestation(address dataMarket, string batchCid, uint256 epochId, bytes32 finalizedCidsRootHash) returns()
func (_Contract *ContractSession) SubmitBatchAttestation(dataMarket common.Address, batchCid string, epochId *big.Int, finalizedCidsRootHash [32]byte) (*types.Transaction, error) {
	return _Contract.Contract.SubmitBatchAttestation(&_Contract.TransactOpts, dataMarket, batchCid, epochId, finalizedCidsRootHash)
}

// SubmitBatchAttestation is a paid mutator transaction binding the contract method 0xcd4c1a34.
//
// Solidity: function submitBatchAttestation(address dataMarket, string batchCid, uint256 epochId, bytes32 finalizedCidsRootHash) returns()
func (_Contract *ContractTransactorSession) SubmitBatchAttestation(dataMarket common.Address, batchCid string, epochId *big.Int, finalizedCidsRootHash [32]byte) (*types.Transaction, error) {
	return _Contract.Contract.SubmitBatchAttestation(&_Contract.TransactOpts, dataMarket, batchCid, epochId, finalizedCidsRootHash)
}

// SubmitSubmissionBatch is a paid mutator transaction binding the contract method 0xcf396cbf.
//
// Solidity: function submitSubmissionBatch(address dataMarket, string batchCid, uint256 epochId, string[] projectIds, string[] snapshotCids, bytes32 finalizedCidsRootHash) returns()
func (_Contract *ContractTransactor) SubmitSubmissionBatch(opts *bind.TransactOpts, dataMarket common.Address, batchCid string, epochId *big.Int, projectIds []string, snapshotCids []string, finalizedCidsRootHash [32]byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "submitSubmissionBatch", dataMarket, batchCid, epochId, projectIds, snapshotCids, finalizedCidsRootHash)
}

// SubmitSubmissionBatch is a paid mutator transaction binding the contract method 0xcf396cbf.
//
// Solidity: function submitSubmissionBatch(address dataMarket, string batchCid, uint256 epochId, string[] projectIds, string[] snapshotCids, bytes32 finalizedCidsRootHash) returns()
func (_Contract *ContractSession) SubmitSubmissionBatch(dataMarket common.Address, batchCid string, epochId *big.Int, projectIds []string, snapshotCids []string, finalizedCidsRootHash [32]byte) (*types.Transaction, error) {
	return _Contract.Contract.SubmitSubmissionBatch(&_Contract.TransactOpts, dataMarket, batchCid, epochId, projectIds, snapshotCids, finalizedCidsRootHash)
}

// SubmitSubmissionBatch is a paid mutator transaction binding the contract method 0xcf396cbf.
//
// Solidity: function submitSubmissionBatch(address dataMarket, string batchCid, uint256 epochId, string[] projectIds, string[] snapshotCids, bytes32 finalizedCidsRootHash) returns()
func (_Contract *ContractTransactorSession) SubmitSubmissionBatch(dataMarket common.Address, batchCid string, epochId *big.Int, projectIds []string, snapshotCids []string, finalizedCidsRootHash [32]byte) (*types.Transaction, error) {
	return _Contract.Contract.SubmitSubmissionBatch(&_Contract.TransactOpts, dataMarket, batchCid, epochId, projectIds, snapshotCids, finalizedCidsRootHash)
}

// ToggleDataMarket is a paid mutator transaction binding the contract method 0xb34aebca.
//
// Solidity: function toggleDataMarket(address dataMarketAddress, bool enabled) returns()
func (_Contract *ContractTransactor) ToggleDataMarket(opts *bind.TransactOpts, dataMarketAddress common.Address, enabled bool) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "toggleDataMarket", dataMarketAddress, enabled)
}

// ToggleDataMarket is a paid mutator transaction binding the contract method 0xb34aebca.
//
// Solidity: function toggleDataMarket(address dataMarketAddress, bool enabled) returns()
func (_Contract *ContractSession) ToggleDataMarket(dataMarketAddress common.Address, enabled bool) (*types.Transaction, error) {
	return _Contract.Contract.ToggleDataMarket(&_Contract.TransactOpts, dataMarketAddress, enabled)
}

// ToggleDataMarket is a paid mutator transaction binding the contract method 0xb34aebca.
//
// Solidity: function toggleDataMarket(address dataMarketAddress, bool enabled) returns()
func (_Contract *ContractTransactorSession) ToggleDataMarket(dataMarketAddress common.Address, enabled bool) (*types.Transaction, error) {
	return _Contract.Contract.ToggleDataMarket(&_Contract.TransactOpts, dataMarketAddress, enabled)
}

// ToggleRewards is a paid mutator transaction binding the contract method 0x71746644.
//
// Solidity: function toggleRewards(address dataMarket) returns()
func (_Contract *ContractTransactor) ToggleRewards(opts *bind.TransactOpts, dataMarket common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "toggleRewards", dataMarket)
}

// ToggleRewards is a paid mutator transaction binding the contract method 0x71746644.
//
// Solidity: function toggleRewards(address dataMarket) returns()
func (_Contract *ContractSession) ToggleRewards(dataMarket common.Address) (*types.Transaction, error) {
	return _Contract.Contract.ToggleRewards(&_Contract.TransactOpts, dataMarket)
}

// ToggleRewards is a paid mutator transaction binding the contract method 0x71746644.
//
// Solidity: function toggleRewards(address dataMarket) returns()
func (_Contract *ContractTransactorSession) ToggleRewards(dataMarket common.Address) (*types.Transaction, error) {
	return _Contract.Contract.ToggleRewards(&_Contract.TransactOpts, dataMarket)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contract *ContractTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contract *ContractSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Contract.Contract.TransferOwnership(&_Contract.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contract *ContractTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Contract.Contract.TransferOwnership(&_Contract.TransactOpts, newOwner)
}

// UpdateAddresses is a paid mutator transaction binding the contract method 0x8d3bcb5a.
//
// Solidity: function updateAddresses(address dataMarket, uint8 role, address[] _addresses, bool[] _status) returns()
func (_Contract *ContractTransactor) UpdateAddresses(opts *bind.TransactOpts, dataMarket common.Address, role uint8, _addresses []common.Address, _status []bool) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "updateAddresses", dataMarket, role, _addresses, _status)
}

// UpdateAddresses is a paid mutator transaction binding the contract method 0x8d3bcb5a.
//
// Solidity: function updateAddresses(address dataMarket, uint8 role, address[] _addresses, bool[] _status) returns()
func (_Contract *ContractSession) UpdateAddresses(dataMarket common.Address, role uint8, _addresses []common.Address, _status []bool) (*types.Transaction, error) {
	return _Contract.Contract.UpdateAddresses(&_Contract.TransactOpts, dataMarket, role, _addresses, _status)
}

// UpdateAddresses is a paid mutator transaction binding the contract method 0x8d3bcb5a.
//
// Solidity: function updateAddresses(address dataMarket, uint8 role, address[] _addresses, bool[] _status) returns()
func (_Contract *ContractTransactorSession) UpdateAddresses(dataMarket common.Address, role uint8, _addresses []common.Address, _status []bool) (*types.Transaction, error) {
	return _Contract.Contract.UpdateAddresses(&_Contract.TransactOpts, dataMarket, role, _addresses, _status)
}

// UpdateAttestationSubmissionWindow is a paid mutator transaction binding the contract method 0x89afe86a.
//
// Solidity: function updateAttestationSubmissionWindow(address dataMarket, uint256 newattestationSubmissionWindow) returns()
func (_Contract *ContractTransactor) UpdateAttestationSubmissionWindow(opts *bind.TransactOpts, dataMarket common.Address, newattestationSubmissionWindow *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "updateAttestationSubmissionWindow", dataMarket, newattestationSubmissionWindow)
}

// UpdateAttestationSubmissionWindow is a paid mutator transaction binding the contract method 0x89afe86a.
//
// Solidity: function updateAttestationSubmissionWindow(address dataMarket, uint256 newattestationSubmissionWindow) returns()
func (_Contract *ContractSession) UpdateAttestationSubmissionWindow(dataMarket common.Address, newattestationSubmissionWindow *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.UpdateAttestationSubmissionWindow(&_Contract.TransactOpts, dataMarket, newattestationSubmissionWindow)
}

// UpdateAttestationSubmissionWindow is a paid mutator transaction binding the contract method 0x89afe86a.
//
// Solidity: function updateAttestationSubmissionWindow(address dataMarket, uint256 newattestationSubmissionWindow) returns()
func (_Contract *ContractTransactorSession) UpdateAttestationSubmissionWindow(dataMarket common.Address, newattestationSubmissionWindow *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.UpdateAttestationSubmissionWindow(&_Contract.TransactOpts, dataMarket, newattestationSubmissionWindow)
}

// UpdateBatchSubmissionWindow is a paid mutator transaction binding the contract method 0xd72e07fd.
//
// Solidity: function updateBatchSubmissionWindow(address dataMarket, uint256 newbatchSubmissionWindow) returns()
func (_Contract *ContractTransactor) UpdateBatchSubmissionWindow(opts *bind.TransactOpts, dataMarket common.Address, newbatchSubmissionWindow *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "updateBatchSubmissionWindow", dataMarket, newbatchSubmissionWindow)
}

// UpdateBatchSubmissionWindow is a paid mutator transaction binding the contract method 0xd72e07fd.
//
// Solidity: function updateBatchSubmissionWindow(address dataMarket, uint256 newbatchSubmissionWindow) returns()
func (_Contract *ContractSession) UpdateBatchSubmissionWindow(dataMarket common.Address, newbatchSubmissionWindow *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.UpdateBatchSubmissionWindow(&_Contract.TransactOpts, dataMarket, newbatchSubmissionWindow)
}

// UpdateBatchSubmissionWindow is a paid mutator transaction binding the contract method 0xd72e07fd.
//
// Solidity: function updateBatchSubmissionWindow(address dataMarket, uint256 newbatchSubmissionWindow) returns()
func (_Contract *ContractTransactorSession) UpdateBatchSubmissionWindow(dataMarket common.Address, newbatchSubmissionWindow *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.UpdateBatchSubmissionWindow(&_Contract.TransactOpts, dataMarket, newbatchSubmissionWindow)
}

// UpdateDailySnapshotQuota is a paid mutator transaction binding the contract method 0xeb4c8b89.
//
// Solidity: function updateDailySnapshotQuota(address dataMarket, uint256 _dailySnapshotQuota) returns()
func (_Contract *ContractTransactor) UpdateDailySnapshotQuota(opts *bind.TransactOpts, dataMarket common.Address, _dailySnapshotQuota *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "updateDailySnapshotQuota", dataMarket, _dailySnapshotQuota)
}

// UpdateDailySnapshotQuota is a paid mutator transaction binding the contract method 0xeb4c8b89.
//
// Solidity: function updateDailySnapshotQuota(address dataMarket, uint256 _dailySnapshotQuota) returns()
func (_Contract *ContractSession) UpdateDailySnapshotQuota(dataMarket common.Address, _dailySnapshotQuota *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.UpdateDailySnapshotQuota(&_Contract.TransactOpts, dataMarket, _dailySnapshotQuota)
}

// UpdateDailySnapshotQuota is a paid mutator transaction binding the contract method 0xeb4c8b89.
//
// Solidity: function updateDailySnapshotQuota(address dataMarket, uint256 _dailySnapshotQuota) returns()
func (_Contract *ContractTransactorSession) UpdateDailySnapshotQuota(dataMarket common.Address, _dailySnapshotQuota *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.UpdateDailySnapshotQuota(&_Contract.TransactOpts, dataMarket, _dailySnapshotQuota)
}

// UpdateDataMarketFactory is a paid mutator transaction binding the contract method 0xb48753eb.
//
// Solidity: function updateDataMarketFactory(address _address) returns()
func (_Contract *ContractTransactor) UpdateDataMarketFactory(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "updateDataMarketFactory", _address)
}

// UpdateDataMarketFactory is a paid mutator transaction binding the contract method 0xb48753eb.
//
// Solidity: function updateDataMarketFactory(address _address) returns()
func (_Contract *ContractSession) UpdateDataMarketFactory(_address common.Address) (*types.Transaction, error) {
	return _Contract.Contract.UpdateDataMarketFactory(&_Contract.TransactOpts, _address)
}

// UpdateDataMarketFactory is a paid mutator transaction binding the contract method 0xb48753eb.
//
// Solidity: function updateDataMarketFactory(address _address) returns()
func (_Contract *ContractTransactorSession) UpdateDataMarketFactory(_address common.Address) (*types.Transaction, error) {
	return _Contract.Contract.UpdateDataMarketFactory(&_Contract.TransactOpts, _address)
}

// UpdateDaySize is a paid mutator transaction binding the contract method 0x79145b89.
//
// Solidity: function updateDaySize(address dataMarket, uint256 newDaySize) returns()
func (_Contract *ContractTransactor) UpdateDaySize(opts *bind.TransactOpts, dataMarket common.Address, newDaySize *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "updateDaySize", dataMarket, newDaySize)
}

// UpdateDaySize is a paid mutator transaction binding the contract method 0x79145b89.
//
// Solidity: function updateDaySize(address dataMarket, uint256 newDaySize) returns()
func (_Contract *ContractSession) UpdateDaySize(dataMarket common.Address, newDaySize *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.UpdateDaySize(&_Contract.TransactOpts, dataMarket, newDaySize)
}

// UpdateDaySize is a paid mutator transaction binding the contract method 0x79145b89.
//
// Solidity: function updateDaySize(address dataMarket, uint256 newDaySize) returns()
func (_Contract *ContractTransactorSession) UpdateDaySize(dataMarket common.Address, newDaySize *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.UpdateDaySize(&_Contract.TransactOpts, dataMarket, newDaySize)
}

// UpdateEpochManager is a paid mutator transaction binding the contract method 0x6e81f234.
//
// Solidity: function updateEpochManager(address dataMarket, address _address) returns()
func (_Contract *ContractTransactor) UpdateEpochManager(opts *bind.TransactOpts, dataMarket common.Address, _address common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "updateEpochManager", dataMarket, _address)
}

// UpdateEpochManager is a paid mutator transaction binding the contract method 0x6e81f234.
//
// Solidity: function updateEpochManager(address dataMarket, address _address) returns()
func (_Contract *ContractSession) UpdateEpochManager(dataMarket common.Address, _address common.Address) (*types.Transaction, error) {
	return _Contract.Contract.UpdateEpochManager(&_Contract.TransactOpts, dataMarket, _address)
}

// UpdateEpochManager is a paid mutator transaction binding the contract method 0x6e81f234.
//
// Solidity: function updateEpochManager(address dataMarket, address _address) returns()
func (_Contract *ContractTransactorSession) UpdateEpochManager(dataMarket common.Address, _address common.Address) (*types.Transaction, error) {
	return _Contract.Contract.UpdateEpochManager(&_Contract.TransactOpts, dataMarket, _address)
}

// UpdateMinAttestationsForConsensus is a paid mutator transaction binding the contract method 0xe4578d51.
//
// Solidity: function updateMinAttestationsForConsensus(address dataMarket, uint256 _minAttestationsForConsensus) returns()
func (_Contract *ContractTransactor) UpdateMinAttestationsForConsensus(opts *bind.TransactOpts, dataMarket common.Address, _minAttestationsForConsensus *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "updateMinAttestationsForConsensus", dataMarket, _minAttestationsForConsensus)
}

// UpdateMinAttestationsForConsensus is a paid mutator transaction binding the contract method 0xe4578d51.
//
// Solidity: function updateMinAttestationsForConsensus(address dataMarket, uint256 _minAttestationsForConsensus) returns()
func (_Contract *ContractSession) UpdateMinAttestationsForConsensus(dataMarket common.Address, _minAttestationsForConsensus *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.UpdateMinAttestationsForConsensus(&_Contract.TransactOpts, dataMarket, _minAttestationsForConsensus)
}

// UpdateMinAttestationsForConsensus is a paid mutator transaction binding the contract method 0xe4578d51.
//
// Solidity: function updateMinAttestationsForConsensus(address dataMarket, uint256 _minAttestationsForConsensus) returns()
func (_Contract *ContractTransactorSession) UpdateMinAttestationsForConsensus(dataMarket common.Address, _minAttestationsForConsensus *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.UpdateMinAttestationsForConsensus(&_Contract.TransactOpts, dataMarket, _minAttestationsForConsensus)
}

// UpdateRewardPoolSize is a paid mutator transaction binding the contract method 0xd9ad5d2e.
//
// Solidity: function updateRewardPoolSize(address dataMarket, uint256 newRewardPoolSize) returns()
func (_Contract *ContractTransactor) UpdateRewardPoolSize(opts *bind.TransactOpts, dataMarket common.Address, newRewardPoolSize *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "updateRewardPoolSize", dataMarket, newRewardPoolSize)
}

// UpdateRewardPoolSize is a paid mutator transaction binding the contract method 0xd9ad5d2e.
//
// Solidity: function updateRewardPoolSize(address dataMarket, uint256 newRewardPoolSize) returns()
func (_Contract *ContractSession) UpdateRewardPoolSize(dataMarket common.Address, newRewardPoolSize *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.UpdateRewardPoolSize(&_Contract.TransactOpts, dataMarket, newRewardPoolSize)
}

// UpdateRewardPoolSize is a paid mutator transaction binding the contract method 0xd9ad5d2e.
//
// Solidity: function updateRewardPoolSize(address dataMarket, uint256 newRewardPoolSize) returns()
func (_Contract *ContractTransactorSession) UpdateRewardPoolSize(dataMarket common.Address, newRewardPoolSize *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.UpdateRewardPoolSize(&_Contract.TransactOpts, dataMarket, newRewardPoolSize)
}

// UpdateRewards is a paid mutator transaction binding the contract method 0x68af906d.
//
// Solidity: function updateRewards(address dataMarket, uint256[] slotIds, uint256[] submissionsList, uint256 day, uint256 eligibleNodes) returns()
func (_Contract *ContractTransactor) UpdateRewards(opts *bind.TransactOpts, dataMarket common.Address, slotIds []*big.Int, submissionsList []*big.Int, day *big.Int, eligibleNodes *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "updateRewards", dataMarket, slotIds, submissionsList, day, eligibleNodes)
}

// UpdateRewards is a paid mutator transaction binding the contract method 0x68af906d.
//
// Solidity: function updateRewards(address dataMarket, uint256[] slotIds, uint256[] submissionsList, uint256 day, uint256 eligibleNodes) returns()
func (_Contract *ContractSession) UpdateRewards(dataMarket common.Address, slotIds []*big.Int, submissionsList []*big.Int, day *big.Int, eligibleNodes *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.UpdateRewards(&_Contract.TransactOpts, dataMarket, slotIds, submissionsList, day, eligibleNodes)
}

// UpdateRewards is a paid mutator transaction binding the contract method 0x68af906d.
//
// Solidity: function updateRewards(address dataMarket, uint256[] slotIds, uint256[] submissionsList, uint256 day, uint256 eligibleNodes) returns()
func (_Contract *ContractTransactorSession) UpdateRewards(dataMarket common.Address, slotIds []*big.Int, submissionsList []*big.Int, day *big.Int, eligibleNodes *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.UpdateRewards(&_Contract.TransactOpts, dataMarket, slotIds, submissionsList, day, eligibleNodes)
}

// UpdateSnapshotSubmissionWindow is a paid mutator transaction binding the contract method 0xa02c3e9b.
//
// Solidity: function updateSnapshotSubmissionWindow(address dataMarket, uint256 newsnapshotSubmissionWindow) returns()
func (_Contract *ContractTransactor) UpdateSnapshotSubmissionWindow(opts *bind.TransactOpts, dataMarket common.Address, newsnapshotSubmissionWindow *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "updateSnapshotSubmissionWindow", dataMarket, newsnapshotSubmissionWindow)
}

// UpdateSnapshotSubmissionWindow is a paid mutator transaction binding the contract method 0xa02c3e9b.
//
// Solidity: function updateSnapshotSubmissionWindow(address dataMarket, uint256 newsnapshotSubmissionWindow) returns()
func (_Contract *ContractSession) UpdateSnapshotSubmissionWindow(dataMarket common.Address, newsnapshotSubmissionWindow *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.UpdateSnapshotSubmissionWindow(&_Contract.TransactOpts, dataMarket, newsnapshotSubmissionWindow)
}

// UpdateSnapshotSubmissionWindow is a paid mutator transaction binding the contract method 0xa02c3e9b.
//
// Solidity: function updateSnapshotSubmissionWindow(address dataMarket, uint256 newsnapshotSubmissionWindow) returns()
func (_Contract *ContractTransactorSession) UpdateSnapshotSubmissionWindow(dataMarket common.Address, newsnapshotSubmissionWindow *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.UpdateSnapshotSubmissionWindow(&_Contract.TransactOpts, dataMarket, newsnapshotSubmissionWindow)
}

// UpdateSnapshotterState is a paid mutator transaction binding the contract method 0xc9742dc1.
//
// Solidity: function updateSnapshotterState(address _address) returns()
func (_Contract *ContractTransactor) UpdateSnapshotterState(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "updateSnapshotterState", _address)
}

// UpdateSnapshotterState is a paid mutator transaction binding the contract method 0xc9742dc1.
//
// Solidity: function updateSnapshotterState(address _address) returns()
func (_Contract *ContractSession) UpdateSnapshotterState(_address common.Address) (*types.Transaction, error) {
	return _Contract.Contract.UpdateSnapshotterState(&_Contract.TransactOpts, _address)
}

// UpdateSnapshotterState is a paid mutator transaction binding the contract method 0xc9742dc1.
//
// Solidity: function updateSnapshotterState(address _address) returns()
func (_Contract *ContractTransactorSession) UpdateSnapshotterState(_address common.Address) (*types.Transaction, error) {
	return _Contract.Contract.UpdateSnapshotterState(&_Contract.TransactOpts, _address)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Contract *ContractTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Contract *ContractSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Contract.Contract.UpgradeToAndCall(&_Contract.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Contract *ContractTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Contract.Contract.UpgradeToAndCall(&_Contract.TransactOpts, newImplementation, data)
}

// ContractAdminsUpdatedIterator is returned from FilterAdminsUpdated and is used to iterate over the raw logs and unpacked data for AdminsUpdated events raised by the Contract contract.
type ContractAdminsUpdatedIterator struct {
	Event *ContractAdminsUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractAdminsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractAdminsUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractAdminsUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractAdminsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractAdminsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractAdminsUpdated represents a AdminsUpdated event raised by the Contract contract.
type ContractAdminsUpdated struct {
	DataMarketAddress common.Address
	AdminAddress      common.Address
	Allowed           bool
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterAdminsUpdated is a free log retrieval operation binding the contract event 0xcde1efc8de533d8d3476a1e5f7423ea506a579776fc27e7d5f963f6c6018ca39.
//
// Solidity: event AdminsUpdated(address indexed dataMarketAddress, address adminAddress, bool allowed)
func (_Contract *ContractFilterer) FilterAdminsUpdated(opts *bind.FilterOpts, dataMarketAddress []common.Address) (*ContractAdminsUpdatedIterator, error) {

	var dataMarketAddressRule []interface{}
	for _, dataMarketAddressItem := range dataMarketAddress {
		dataMarketAddressRule = append(dataMarketAddressRule, dataMarketAddressItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "AdminsUpdated", dataMarketAddressRule)
	if err != nil {
		return nil, err
	}
	return &ContractAdminsUpdatedIterator{contract: _Contract.contract, event: "AdminsUpdated", logs: logs, sub: sub}, nil
}

// WatchAdminsUpdated is a free log subscription operation binding the contract event 0xcde1efc8de533d8d3476a1e5f7423ea506a579776fc27e7d5f963f6c6018ca39.
//
// Solidity: event AdminsUpdated(address indexed dataMarketAddress, address adminAddress, bool allowed)
func (_Contract *ContractFilterer) WatchAdminsUpdated(opts *bind.WatchOpts, sink chan<- *ContractAdminsUpdated, dataMarketAddress []common.Address) (event.Subscription, error) {

	var dataMarketAddressRule []interface{}
	for _, dataMarketAddressItem := range dataMarketAddress {
		dataMarketAddressRule = append(dataMarketAddressRule, dataMarketAddressItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "AdminsUpdated", dataMarketAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractAdminsUpdated)
				if err := _Contract.contract.UnpackLog(event, "AdminsUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAdminsUpdated is a log parse operation binding the contract event 0xcde1efc8de533d8d3476a1e5f7423ea506a579776fc27e7d5f963f6c6018ca39.
//
// Solidity: event AdminsUpdated(address indexed dataMarketAddress, address adminAddress, bool allowed)
func (_Contract *ContractFilterer) ParseAdminsUpdated(log types.Log) (*ContractAdminsUpdated, error) {
	event := new(ContractAdminsUpdated)
	if err := _Contract.contract.UnpackLog(event, "AdminsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractBatchSubmissionsCompletedIterator is returned from FilterBatchSubmissionsCompleted and is used to iterate over the raw logs and unpacked data for BatchSubmissionsCompleted events raised by the Contract contract.
type ContractBatchSubmissionsCompletedIterator struct {
	Event *ContractBatchSubmissionsCompleted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractBatchSubmissionsCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractBatchSubmissionsCompleted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractBatchSubmissionsCompleted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractBatchSubmissionsCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractBatchSubmissionsCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractBatchSubmissionsCompleted represents a BatchSubmissionsCompleted event raised by the Contract contract.
type ContractBatchSubmissionsCompleted struct {
	DataMarketAddress common.Address
	EpochId           *big.Int
	Timestamp         *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterBatchSubmissionsCompleted is a free log retrieval operation binding the contract event 0x5ed90e6a0e12831302fecf13a2cc6f7e9439fed821340d0bddc4ae305dbf9c30.
//
// Solidity: event BatchSubmissionsCompleted(address indexed dataMarketAddress, uint256 indexed epochId, uint256 timestamp)
func (_Contract *ContractFilterer) FilterBatchSubmissionsCompleted(opts *bind.FilterOpts, dataMarketAddress []common.Address, epochId []*big.Int) (*ContractBatchSubmissionsCompletedIterator, error) {

	var dataMarketAddressRule []interface{}
	for _, dataMarketAddressItem := range dataMarketAddress {
		dataMarketAddressRule = append(dataMarketAddressRule, dataMarketAddressItem)
	}
	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "BatchSubmissionsCompleted", dataMarketAddressRule, epochIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractBatchSubmissionsCompletedIterator{contract: _Contract.contract, event: "BatchSubmissionsCompleted", logs: logs, sub: sub}, nil
}

// WatchBatchSubmissionsCompleted is a free log subscription operation binding the contract event 0x5ed90e6a0e12831302fecf13a2cc6f7e9439fed821340d0bddc4ae305dbf9c30.
//
// Solidity: event BatchSubmissionsCompleted(address indexed dataMarketAddress, uint256 indexed epochId, uint256 timestamp)
func (_Contract *ContractFilterer) WatchBatchSubmissionsCompleted(opts *bind.WatchOpts, sink chan<- *ContractBatchSubmissionsCompleted, dataMarketAddress []common.Address, epochId []*big.Int) (event.Subscription, error) {

	var dataMarketAddressRule []interface{}
	for _, dataMarketAddressItem := range dataMarketAddress {
		dataMarketAddressRule = append(dataMarketAddressRule, dataMarketAddressItem)
	}
	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "BatchSubmissionsCompleted", dataMarketAddressRule, epochIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractBatchSubmissionsCompleted)
				if err := _Contract.contract.UnpackLog(event, "BatchSubmissionsCompleted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBatchSubmissionsCompleted is a log parse operation binding the contract event 0x5ed90e6a0e12831302fecf13a2cc6f7e9439fed821340d0bddc4ae305dbf9c30.
//
// Solidity: event BatchSubmissionsCompleted(address indexed dataMarketAddress, uint256 indexed epochId, uint256 timestamp)
func (_Contract *ContractFilterer) ParseBatchSubmissionsCompleted(log types.Log) (*ContractBatchSubmissionsCompleted, error) {
	event := new(ContractBatchSubmissionsCompleted)
	if err := _Contract.contract.UnpackLog(event, "BatchSubmissionsCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractDailyTaskCompletedEventIterator is returned from FilterDailyTaskCompletedEvent and is used to iterate over the raw logs and unpacked data for DailyTaskCompletedEvent events raised by the Contract contract.
type ContractDailyTaskCompletedEventIterator struct {
	Event *ContractDailyTaskCompletedEvent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractDailyTaskCompletedEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractDailyTaskCompletedEvent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractDailyTaskCompletedEvent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractDailyTaskCompletedEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractDailyTaskCompletedEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractDailyTaskCompletedEvent represents a DailyTaskCompletedEvent event raised by the Contract contract.
type ContractDailyTaskCompletedEvent struct {
	DataMarketAddress  common.Address
	SnapshotterAddress common.Address
	SlotId             *big.Int
	DayId              *big.Int
	RewardPoints       *big.Int
	Timestamp          *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterDailyTaskCompletedEvent is a free log retrieval operation binding the contract event 0x3f80954353f9060121cb024ef2580d122fc250f7faad898ef5167aeef38f2b12.
//
// Solidity: event DailyTaskCompletedEvent(address indexed dataMarketAddress, address snapshotterAddress, uint256 slotId, uint256 dayId, uint256 rewardPoints, uint256 timestamp)
func (_Contract *ContractFilterer) FilterDailyTaskCompletedEvent(opts *bind.FilterOpts, dataMarketAddress []common.Address) (*ContractDailyTaskCompletedEventIterator, error) {

	var dataMarketAddressRule []interface{}
	for _, dataMarketAddressItem := range dataMarketAddress {
		dataMarketAddressRule = append(dataMarketAddressRule, dataMarketAddressItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "DailyTaskCompletedEvent", dataMarketAddressRule)
	if err != nil {
		return nil, err
	}
	return &ContractDailyTaskCompletedEventIterator{contract: _Contract.contract, event: "DailyTaskCompletedEvent", logs: logs, sub: sub}, nil
}

// WatchDailyTaskCompletedEvent is a free log subscription operation binding the contract event 0x3f80954353f9060121cb024ef2580d122fc250f7faad898ef5167aeef38f2b12.
//
// Solidity: event DailyTaskCompletedEvent(address indexed dataMarketAddress, address snapshotterAddress, uint256 slotId, uint256 dayId, uint256 rewardPoints, uint256 timestamp)
func (_Contract *ContractFilterer) WatchDailyTaskCompletedEvent(opts *bind.WatchOpts, sink chan<- *ContractDailyTaskCompletedEvent, dataMarketAddress []common.Address) (event.Subscription, error) {

	var dataMarketAddressRule []interface{}
	for _, dataMarketAddressItem := range dataMarketAddress {
		dataMarketAddressRule = append(dataMarketAddressRule, dataMarketAddressItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "DailyTaskCompletedEvent", dataMarketAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractDailyTaskCompletedEvent)
				if err := _Contract.contract.UnpackLog(event, "DailyTaskCompletedEvent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDailyTaskCompletedEvent is a log parse operation binding the contract event 0x3f80954353f9060121cb024ef2580d122fc250f7faad898ef5167aeef38f2b12.
//
// Solidity: event DailyTaskCompletedEvent(address indexed dataMarketAddress, address snapshotterAddress, uint256 slotId, uint256 dayId, uint256 rewardPoints, uint256 timestamp)
func (_Contract *ContractFilterer) ParseDailyTaskCompletedEvent(log types.Log) (*ContractDailyTaskCompletedEvent, error) {
	event := new(ContractDailyTaskCompletedEvent)
	if err := _Contract.contract.UnpackLog(event, "DailyTaskCompletedEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractDataMarketCreatedIterator is returned from FilterDataMarketCreated and is used to iterate over the raw logs and unpacked data for DataMarketCreated events raised by the Contract contract.
type ContractDataMarketCreatedIterator struct {
	Event *ContractDataMarketCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractDataMarketCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractDataMarketCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractDataMarketCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractDataMarketCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractDataMarketCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractDataMarketCreated represents a DataMarketCreated event raised by the Contract contract.
type ContractDataMarketCreated struct {
	OwnerAddress            common.Address
	EpochSize               uint8
	SourceChainId           *big.Int
	SourceChainBlockTime    *big.Int
	UseBlockNumberAsEpochId bool
	ProtocolState           common.Address
	DataMarketAddress       common.Address
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterDataMarketCreated is a free log retrieval operation binding the contract event 0x4ac24bfdd2a01328b05db8f3587611f1d37fd01718e1edaa23a75ec91d71f517.
//
// Solidity: event DataMarketCreated(address indexed ownerAddress, uint8 epochSize, uint256 sourceChainId, uint256 sourceChainBlockTime, bool useBlockNumberAsEpochId, address protocolState, address dataMarketAddress)
func (_Contract *ContractFilterer) FilterDataMarketCreated(opts *bind.FilterOpts, ownerAddress []common.Address) (*ContractDataMarketCreatedIterator, error) {

	var ownerAddressRule []interface{}
	for _, ownerAddressItem := range ownerAddress {
		ownerAddressRule = append(ownerAddressRule, ownerAddressItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "DataMarketCreated", ownerAddressRule)
	if err != nil {
		return nil, err
	}
	return &ContractDataMarketCreatedIterator{contract: _Contract.contract, event: "DataMarketCreated", logs: logs, sub: sub}, nil
}

// WatchDataMarketCreated is a free log subscription operation binding the contract event 0x4ac24bfdd2a01328b05db8f3587611f1d37fd01718e1edaa23a75ec91d71f517.
//
// Solidity: event DataMarketCreated(address indexed ownerAddress, uint8 epochSize, uint256 sourceChainId, uint256 sourceChainBlockTime, bool useBlockNumberAsEpochId, address protocolState, address dataMarketAddress)
func (_Contract *ContractFilterer) WatchDataMarketCreated(opts *bind.WatchOpts, sink chan<- *ContractDataMarketCreated, ownerAddress []common.Address) (event.Subscription, error) {

	var ownerAddressRule []interface{}
	for _, ownerAddressItem := range ownerAddress {
		ownerAddressRule = append(ownerAddressRule, ownerAddressItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "DataMarketCreated", ownerAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractDataMarketCreated)
				if err := _Contract.contract.UnpackLog(event, "DataMarketCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDataMarketCreated is a log parse operation binding the contract event 0x4ac24bfdd2a01328b05db8f3587611f1d37fd01718e1edaa23a75ec91d71f517.
//
// Solidity: event DataMarketCreated(address indexed ownerAddress, uint8 epochSize, uint256 sourceChainId, uint256 sourceChainBlockTime, bool useBlockNumberAsEpochId, address protocolState, address dataMarketAddress)
func (_Contract *ContractFilterer) ParseDataMarketCreated(log types.Log) (*ContractDataMarketCreated, error) {
	event := new(ContractDataMarketCreated)
	if err := _Contract.contract.UnpackLog(event, "DataMarketCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractDayStartedEventIterator is returned from FilterDayStartedEvent and is used to iterate over the raw logs and unpacked data for DayStartedEvent events raised by the Contract contract.
type ContractDayStartedEventIterator struct {
	Event *ContractDayStartedEvent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractDayStartedEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractDayStartedEvent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractDayStartedEvent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractDayStartedEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractDayStartedEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractDayStartedEvent represents a DayStartedEvent event raised by the Contract contract.
type ContractDayStartedEvent struct {
	DataMarketAddress common.Address
	DayId             *big.Int
	Timestamp         *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterDayStartedEvent is a free log retrieval operation binding the contract event 0x9bf380fe36617cd5d995968abb4ae6d3657a763c126535d590b4503ff2542426.
//
// Solidity: event DayStartedEvent(address indexed dataMarketAddress, uint256 dayId, uint256 timestamp)
func (_Contract *ContractFilterer) FilterDayStartedEvent(opts *bind.FilterOpts, dataMarketAddress []common.Address) (*ContractDayStartedEventIterator, error) {

	var dataMarketAddressRule []interface{}
	for _, dataMarketAddressItem := range dataMarketAddress {
		dataMarketAddressRule = append(dataMarketAddressRule, dataMarketAddressItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "DayStartedEvent", dataMarketAddressRule)
	if err != nil {
		return nil, err
	}
	return &ContractDayStartedEventIterator{contract: _Contract.contract, event: "DayStartedEvent", logs: logs, sub: sub}, nil
}

// WatchDayStartedEvent is a free log subscription operation binding the contract event 0x9bf380fe36617cd5d995968abb4ae6d3657a763c126535d590b4503ff2542426.
//
// Solidity: event DayStartedEvent(address indexed dataMarketAddress, uint256 dayId, uint256 timestamp)
func (_Contract *ContractFilterer) WatchDayStartedEvent(opts *bind.WatchOpts, sink chan<- *ContractDayStartedEvent, dataMarketAddress []common.Address) (event.Subscription, error) {

	var dataMarketAddressRule []interface{}
	for _, dataMarketAddressItem := range dataMarketAddress {
		dataMarketAddressRule = append(dataMarketAddressRule, dataMarketAddressItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "DayStartedEvent", dataMarketAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractDayStartedEvent)
				if err := _Contract.contract.UnpackLog(event, "DayStartedEvent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDayStartedEvent is a log parse operation binding the contract event 0x9bf380fe36617cd5d995968abb4ae6d3657a763c126535d590b4503ff2542426.
//
// Solidity: event DayStartedEvent(address indexed dataMarketAddress, uint256 dayId, uint256 timestamp)
func (_Contract *ContractFilterer) ParseDayStartedEvent(log types.Log) (*ContractDayStartedEvent, error) {
	event := new(ContractDayStartedEvent)
	if err := _Contract.contract.UnpackLog(event, "DayStartedEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractDelayedAttestationSubmittedIterator is returned from FilterDelayedAttestationSubmitted and is used to iterate over the raw logs and unpacked data for DelayedAttestationSubmitted events raised by the Contract contract.
type ContractDelayedAttestationSubmittedIterator struct {
	Event *ContractDelayedAttestationSubmitted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractDelayedAttestationSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractDelayedAttestationSubmitted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractDelayedAttestationSubmitted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractDelayedAttestationSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractDelayedAttestationSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractDelayedAttestationSubmitted represents a DelayedAttestationSubmitted event raised by the Contract contract.
type ContractDelayedAttestationSubmitted struct {
	DataMarketAddress common.Address
	BatchCid          string
	EpochId           *big.Int
	Timestamp         *big.Int
	ValidatorAddr     common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterDelayedAttestationSubmitted is a free log retrieval operation binding the contract event 0x23db68f6127736e9b15d04ab69c48dc844e007832c59edaf8f60cae2dd638808.
//
// Solidity: event DelayedAttestationSubmitted(address indexed dataMarketAddress, string batchCid, uint256 indexed epochId, uint256 timestamp, address indexed validatorAddr)
func (_Contract *ContractFilterer) FilterDelayedAttestationSubmitted(opts *bind.FilterOpts, dataMarketAddress []common.Address, epochId []*big.Int, validatorAddr []common.Address) (*ContractDelayedAttestationSubmittedIterator, error) {

	var dataMarketAddressRule []interface{}
	for _, dataMarketAddressItem := range dataMarketAddress {
		dataMarketAddressRule = append(dataMarketAddressRule, dataMarketAddressItem)
	}

	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}

	var validatorAddrRule []interface{}
	for _, validatorAddrItem := range validatorAddr {
		validatorAddrRule = append(validatorAddrRule, validatorAddrItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "DelayedAttestationSubmitted", dataMarketAddressRule, epochIdRule, validatorAddrRule)
	if err != nil {
		return nil, err
	}
	return &ContractDelayedAttestationSubmittedIterator{contract: _Contract.contract, event: "DelayedAttestationSubmitted", logs: logs, sub: sub}, nil
}

// WatchDelayedAttestationSubmitted is a free log subscription operation binding the contract event 0x23db68f6127736e9b15d04ab69c48dc844e007832c59edaf8f60cae2dd638808.
//
// Solidity: event DelayedAttestationSubmitted(address indexed dataMarketAddress, string batchCid, uint256 indexed epochId, uint256 timestamp, address indexed validatorAddr)
func (_Contract *ContractFilterer) WatchDelayedAttestationSubmitted(opts *bind.WatchOpts, sink chan<- *ContractDelayedAttestationSubmitted, dataMarketAddress []common.Address, epochId []*big.Int, validatorAddr []common.Address) (event.Subscription, error) {

	var dataMarketAddressRule []interface{}
	for _, dataMarketAddressItem := range dataMarketAddress {
		dataMarketAddressRule = append(dataMarketAddressRule, dataMarketAddressItem)
	}

	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}

	var validatorAddrRule []interface{}
	for _, validatorAddrItem := range validatorAddr {
		validatorAddrRule = append(validatorAddrRule, validatorAddrItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "DelayedAttestationSubmitted", dataMarketAddressRule, epochIdRule, validatorAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractDelayedAttestationSubmitted)
				if err := _Contract.contract.UnpackLog(event, "DelayedAttestationSubmitted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDelayedAttestationSubmitted is a log parse operation binding the contract event 0x23db68f6127736e9b15d04ab69c48dc844e007832c59edaf8f60cae2dd638808.
//
// Solidity: event DelayedAttestationSubmitted(address indexed dataMarketAddress, string batchCid, uint256 indexed epochId, uint256 timestamp, address indexed validatorAddr)
func (_Contract *ContractFilterer) ParseDelayedAttestationSubmitted(log types.Log) (*ContractDelayedAttestationSubmitted, error) {
	event := new(ContractDelayedAttestationSubmitted)
	if err := _Contract.contract.UnpackLog(event, "DelayedAttestationSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractDelayedBatchSubmittedIterator is returned from FilterDelayedBatchSubmitted and is used to iterate over the raw logs and unpacked data for DelayedBatchSubmitted events raised by the Contract contract.
type ContractDelayedBatchSubmittedIterator struct {
	Event *ContractDelayedBatchSubmitted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractDelayedBatchSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractDelayedBatchSubmitted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractDelayedBatchSubmitted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractDelayedBatchSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractDelayedBatchSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractDelayedBatchSubmitted represents a DelayedBatchSubmitted event raised by the Contract contract.
type ContractDelayedBatchSubmitted struct {
	DataMarketAddress common.Address
	BatchCid          string
	EpochId           *big.Int
	Timestamp         *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterDelayedBatchSubmitted is a free log retrieval operation binding the contract event 0xa48b3f37894ae648db7ab45db773764cc131b15cb19978e17da7bcd79eaf64be.
//
// Solidity: event DelayedBatchSubmitted(address indexed dataMarketAddress, string batchCid, uint256 indexed epochId, uint256 timestamp)
func (_Contract *ContractFilterer) FilterDelayedBatchSubmitted(opts *bind.FilterOpts, dataMarketAddress []common.Address, epochId []*big.Int) (*ContractDelayedBatchSubmittedIterator, error) {

	var dataMarketAddressRule []interface{}
	for _, dataMarketAddressItem := range dataMarketAddress {
		dataMarketAddressRule = append(dataMarketAddressRule, dataMarketAddressItem)
	}

	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "DelayedBatchSubmitted", dataMarketAddressRule, epochIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractDelayedBatchSubmittedIterator{contract: _Contract.contract, event: "DelayedBatchSubmitted", logs: logs, sub: sub}, nil
}

// WatchDelayedBatchSubmitted is a free log subscription operation binding the contract event 0xa48b3f37894ae648db7ab45db773764cc131b15cb19978e17da7bcd79eaf64be.
//
// Solidity: event DelayedBatchSubmitted(address indexed dataMarketAddress, string batchCid, uint256 indexed epochId, uint256 timestamp)
func (_Contract *ContractFilterer) WatchDelayedBatchSubmitted(opts *bind.WatchOpts, sink chan<- *ContractDelayedBatchSubmitted, dataMarketAddress []common.Address, epochId []*big.Int) (event.Subscription, error) {

	var dataMarketAddressRule []interface{}
	for _, dataMarketAddressItem := range dataMarketAddress {
		dataMarketAddressRule = append(dataMarketAddressRule, dataMarketAddressItem)
	}

	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "DelayedBatchSubmitted", dataMarketAddressRule, epochIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractDelayedBatchSubmitted)
				if err := _Contract.contract.UnpackLog(event, "DelayedBatchSubmitted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDelayedBatchSubmitted is a log parse operation binding the contract event 0xa48b3f37894ae648db7ab45db773764cc131b15cb19978e17da7bcd79eaf64be.
//
// Solidity: event DelayedBatchSubmitted(address indexed dataMarketAddress, string batchCid, uint256 indexed epochId, uint256 timestamp)
func (_Contract *ContractFilterer) ParseDelayedBatchSubmitted(log types.Log) (*ContractDelayedBatchSubmitted, error) {
	event := new(ContractDelayedBatchSubmitted)
	if err := _Contract.contract.UnpackLog(event, "DelayedBatchSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractDelayedSnapshotSubmittedIterator is returned from FilterDelayedSnapshotSubmitted and is used to iterate over the raw logs and unpacked data for DelayedSnapshotSubmitted events raised by the Contract contract.
type ContractDelayedSnapshotSubmittedIterator struct {
	Event *ContractDelayedSnapshotSubmitted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractDelayedSnapshotSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractDelayedSnapshotSubmitted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractDelayedSnapshotSubmitted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractDelayedSnapshotSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractDelayedSnapshotSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractDelayedSnapshotSubmitted represents a DelayedSnapshotSubmitted event raised by the Contract contract.
type ContractDelayedSnapshotSubmitted struct {
	DataMarketAddress common.Address
	SnapshotterAddr   common.Address
	SlotId            *big.Int
	SnapshotCid       string
	EpochId           *big.Int
	ProjectId         string
	Timestamp         *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterDelayedSnapshotSubmitted is a free log retrieval operation binding the contract event 0x3857c02d90218ce4e6decef6b24babba8e0d8331c159392de06c1ce2c7a2d3be.
//
// Solidity: event DelayedSnapshotSubmitted(address indexed dataMarketAddress, address indexed snapshotterAddr, uint256 slotId, string snapshotCid, uint256 indexed epochId, string projectId, uint256 timestamp)
func (_Contract *ContractFilterer) FilterDelayedSnapshotSubmitted(opts *bind.FilterOpts, dataMarketAddress []common.Address, snapshotterAddr []common.Address, epochId []*big.Int) (*ContractDelayedSnapshotSubmittedIterator, error) {

	var dataMarketAddressRule []interface{}
	for _, dataMarketAddressItem := range dataMarketAddress {
		dataMarketAddressRule = append(dataMarketAddressRule, dataMarketAddressItem)
	}
	var snapshotterAddrRule []interface{}
	for _, snapshotterAddrItem := range snapshotterAddr {
		snapshotterAddrRule = append(snapshotterAddrRule, snapshotterAddrItem)
	}

	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "DelayedSnapshotSubmitted", dataMarketAddressRule, snapshotterAddrRule, epochIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractDelayedSnapshotSubmittedIterator{contract: _Contract.contract, event: "DelayedSnapshotSubmitted", logs: logs, sub: sub}, nil
}

// WatchDelayedSnapshotSubmitted is a free log subscription operation binding the contract event 0x3857c02d90218ce4e6decef6b24babba8e0d8331c159392de06c1ce2c7a2d3be.
//
// Solidity: event DelayedSnapshotSubmitted(address indexed dataMarketAddress, address indexed snapshotterAddr, uint256 slotId, string snapshotCid, uint256 indexed epochId, string projectId, uint256 timestamp)
func (_Contract *ContractFilterer) WatchDelayedSnapshotSubmitted(opts *bind.WatchOpts, sink chan<- *ContractDelayedSnapshotSubmitted, dataMarketAddress []common.Address, snapshotterAddr []common.Address, epochId []*big.Int) (event.Subscription, error) {

	var dataMarketAddressRule []interface{}
	for _, dataMarketAddressItem := range dataMarketAddress {
		dataMarketAddressRule = append(dataMarketAddressRule, dataMarketAddressItem)
	}
	var snapshotterAddrRule []interface{}
	for _, snapshotterAddrItem := range snapshotterAddr {
		snapshotterAddrRule = append(snapshotterAddrRule, snapshotterAddrItem)
	}

	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "DelayedSnapshotSubmitted", dataMarketAddressRule, snapshotterAddrRule, epochIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractDelayedSnapshotSubmitted)
				if err := _Contract.contract.UnpackLog(event, "DelayedSnapshotSubmitted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDelayedSnapshotSubmitted is a log parse operation binding the contract event 0x3857c02d90218ce4e6decef6b24babba8e0d8331c159392de06c1ce2c7a2d3be.
//
// Solidity: event DelayedSnapshotSubmitted(address indexed dataMarketAddress, address indexed snapshotterAddr, uint256 slotId, string snapshotCid, uint256 indexed epochId, string projectId, uint256 timestamp)
func (_Contract *ContractFilterer) ParseDelayedSnapshotSubmitted(log types.Log) (*ContractDelayedSnapshotSubmitted, error) {
	event := new(ContractDelayedSnapshotSubmitted)
	if err := _Contract.contract.UnpackLog(event, "DelayedSnapshotSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractEmergencyWithdrawIterator is returned from FilterEmergencyWithdraw and is used to iterate over the raw logs and unpacked data for EmergencyWithdraw events raised by the Contract contract.
type ContractEmergencyWithdrawIterator struct {
	Event *ContractEmergencyWithdraw // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractEmergencyWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractEmergencyWithdraw)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractEmergencyWithdraw)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractEmergencyWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractEmergencyWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractEmergencyWithdraw represents a EmergencyWithdraw event raised by the Contract contract.
type ContractEmergencyWithdraw struct {
	Owner  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEmergencyWithdraw is a free log retrieval operation binding the contract event 0x5fafa99d0643513820be26656b45130b01e1c03062e1266bf36f88cbd3bd9695.
//
// Solidity: event EmergencyWithdraw(address indexed owner, uint256 amount)
func (_Contract *ContractFilterer) FilterEmergencyWithdraw(opts *bind.FilterOpts, owner []common.Address) (*ContractEmergencyWithdrawIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "EmergencyWithdraw", ownerRule)
	if err != nil {
		return nil, err
	}
	return &ContractEmergencyWithdrawIterator{contract: _Contract.contract, event: "EmergencyWithdraw", logs: logs, sub: sub}, nil
}

// WatchEmergencyWithdraw is a free log subscription operation binding the contract event 0x5fafa99d0643513820be26656b45130b01e1c03062e1266bf36f88cbd3bd9695.
//
// Solidity: event EmergencyWithdraw(address indexed owner, uint256 amount)
func (_Contract *ContractFilterer) WatchEmergencyWithdraw(opts *bind.WatchOpts, sink chan<- *ContractEmergencyWithdraw, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "EmergencyWithdraw", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractEmergencyWithdraw)
				if err := _Contract.contract.UnpackLog(event, "EmergencyWithdraw", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseEmergencyWithdraw is a log parse operation binding the contract event 0x5fafa99d0643513820be26656b45130b01e1c03062e1266bf36f88cbd3bd9695.
//
// Solidity: event EmergencyWithdraw(address indexed owner, uint256 amount)
func (_Contract *ContractFilterer) ParseEmergencyWithdraw(log types.Log) (*ContractEmergencyWithdraw, error) {
	event := new(ContractEmergencyWithdraw)
	if err := _Contract.contract.UnpackLog(event, "EmergencyWithdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractEpochReleasedIterator is returned from FilterEpochReleased and is used to iterate over the raw logs and unpacked data for EpochReleased events raised by the Contract contract.
type ContractEpochReleasedIterator struct {
	Event *ContractEpochReleased // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractEpochReleasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractEpochReleased)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractEpochReleased)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractEpochReleasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractEpochReleasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractEpochReleased represents a EpochReleased event raised by the Contract contract.
type ContractEpochReleased struct {
	DataMarketAddress common.Address
	EpochId           *big.Int
	Begin             *big.Int
	End               *big.Int
	Timestamp         *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterEpochReleased is a free log retrieval operation binding the contract event 0xf7d2257d4a1c445138ab52bd3c22425bfed29da81d0173961c697dc14fcba60c.
//
// Solidity: event EpochReleased(address indexed dataMarketAddress, uint256 indexed epochId, uint256 begin, uint256 end, uint256 timestamp)
func (_Contract *ContractFilterer) FilterEpochReleased(opts *bind.FilterOpts, dataMarketAddress []common.Address, epochId []*big.Int) (*ContractEpochReleasedIterator, error) {

	var dataMarketAddressRule []interface{}
	for _, dataMarketAddressItem := range dataMarketAddress {
		dataMarketAddressRule = append(dataMarketAddressRule, dataMarketAddressItem)
	}
	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "EpochReleased", dataMarketAddressRule, epochIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractEpochReleasedIterator{contract: _Contract.contract, event: "EpochReleased", logs: logs, sub: sub}, nil
}

// WatchEpochReleased is a free log subscription operation binding the contract event 0xf7d2257d4a1c445138ab52bd3c22425bfed29da81d0173961c697dc14fcba60c.
//
// Solidity: event EpochReleased(address indexed dataMarketAddress, uint256 indexed epochId, uint256 begin, uint256 end, uint256 timestamp)
func (_Contract *ContractFilterer) WatchEpochReleased(opts *bind.WatchOpts, sink chan<- *ContractEpochReleased, dataMarketAddress []common.Address, epochId []*big.Int) (event.Subscription, error) {

	var dataMarketAddressRule []interface{}
	for _, dataMarketAddressItem := range dataMarketAddress {
		dataMarketAddressRule = append(dataMarketAddressRule, dataMarketAddressItem)
	}
	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "EpochReleased", dataMarketAddressRule, epochIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractEpochReleased)
				if err := _Contract.contract.UnpackLog(event, "EpochReleased", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseEpochReleased is a log parse operation binding the contract event 0xf7d2257d4a1c445138ab52bd3c22425bfed29da81d0173961c697dc14fcba60c.
//
// Solidity: event EpochReleased(address indexed dataMarketAddress, uint256 indexed epochId, uint256 begin, uint256 end, uint256 timestamp)
func (_Contract *ContractFilterer) ParseEpochReleased(log types.Log) (*ContractEpochReleased, error) {
	event := new(ContractEpochReleased)
	if err := _Contract.contract.UnpackLog(event, "EpochReleased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Contract contract.
type ContractInitializedIterator struct {
	Event *ContractInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractInitialized represents a Initialized event raised by the Contract contract.
type ContractInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Contract *ContractFilterer) FilterInitialized(opts *bind.FilterOpts) (*ContractInitializedIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ContractInitializedIterator{contract: _Contract.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Contract *ContractFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractInitialized) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractInitialized)
				if err := _Contract.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Contract *ContractFilterer) ParseInitialized(log types.Log) (*ContractInitialized, error) {
	event := new(ContractInitialized)
	if err := _Contract.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Contract contract.
type ContractOwnershipTransferredIterator struct {
	Event *ContractOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOwnershipTransferred represents a OwnershipTransferred event raised by the Contract contract.
type ContractOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contract *ContractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractOwnershipTransferredIterator{contract: _Contract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contract *ContractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOwnershipTransferred)
				if err := _Contract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contract *ContractFilterer) ParseOwnershipTransferred(log types.Log) (*ContractOwnershipTransferred, error) {
	event := new(ContractOwnershipTransferred)
	if err := _Contract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractRewardsClaimedIterator is returned from FilterRewardsClaimed and is used to iterate over the raw logs and unpacked data for RewardsClaimed events raised by the Contract contract.
type ContractRewardsClaimedIterator struct {
	Event *ContractRewardsClaimed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractRewardsClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractRewardsClaimed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractRewardsClaimed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractRewardsClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractRewardsClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractRewardsClaimed represents a RewardsClaimed event raised by the Contract contract.
type ContractRewardsClaimed struct {
	User      common.Address
	Amount    *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRewardsClaimed is a free log retrieval operation binding the contract event 0xdacbdde355ba930696a362ea6738feb9f8bd52dfb3d81947558fd3217e23e325.
//
// Solidity: event RewardsClaimed(address indexed user, uint256 amount, uint256 timestamp)
func (_Contract *ContractFilterer) FilterRewardsClaimed(opts *bind.FilterOpts, user []common.Address) (*ContractRewardsClaimedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "RewardsClaimed", userRule)
	if err != nil {
		return nil, err
	}
	return &ContractRewardsClaimedIterator{contract: _Contract.contract, event: "RewardsClaimed", logs: logs, sub: sub}, nil
}

// WatchRewardsClaimed is a free log subscription operation binding the contract event 0xdacbdde355ba930696a362ea6738feb9f8bd52dfb3d81947558fd3217e23e325.
//
// Solidity: event RewardsClaimed(address indexed user, uint256 amount, uint256 timestamp)
func (_Contract *ContractFilterer) WatchRewardsClaimed(opts *bind.WatchOpts, sink chan<- *ContractRewardsClaimed, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "RewardsClaimed", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractRewardsClaimed)
				if err := _Contract.contract.UnpackLog(event, "RewardsClaimed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRewardsClaimed is a log parse operation binding the contract event 0xdacbdde355ba930696a362ea6738feb9f8bd52dfb3d81947558fd3217e23e325.
//
// Solidity: event RewardsClaimed(address indexed user, uint256 amount, uint256 timestamp)
func (_Contract *ContractFilterer) ParseRewardsClaimed(log types.Log) (*ContractRewardsClaimed, error) {
	event := new(ContractRewardsClaimed)
	if err := _Contract.contract.UnpackLog(event, "RewardsClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSequencersUpdatedIterator is returned from FilterSequencersUpdated and is used to iterate over the raw logs and unpacked data for SequencersUpdated events raised by the Contract contract.
type ContractSequencersUpdatedIterator struct {
	Event *ContractSequencersUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractSequencersUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSequencersUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractSequencersUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractSequencersUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSequencersUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSequencersUpdated represents a SequencersUpdated event raised by the Contract contract.
type ContractSequencersUpdated struct {
	DataMarketAddress common.Address
	SequencerAddress  common.Address
	Allowed           bool
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterSequencersUpdated is a free log retrieval operation binding the contract event 0xad12010237fe83915c67abe51836e693a16f8a2592b9a4e959374ab33ae7a697.
//
// Solidity: event SequencersUpdated(address indexed dataMarketAddress, address sequencerAddress, bool allowed)
func (_Contract *ContractFilterer) FilterSequencersUpdated(opts *bind.FilterOpts, dataMarketAddress []common.Address) (*ContractSequencersUpdatedIterator, error) {

	var dataMarketAddressRule []interface{}
	for _, dataMarketAddressItem := range dataMarketAddress {
		dataMarketAddressRule = append(dataMarketAddressRule, dataMarketAddressItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "SequencersUpdated", dataMarketAddressRule)
	if err != nil {
		return nil, err
	}
	return &ContractSequencersUpdatedIterator{contract: _Contract.contract, event: "SequencersUpdated", logs: logs, sub: sub}, nil
}

// WatchSequencersUpdated is a free log subscription operation binding the contract event 0xad12010237fe83915c67abe51836e693a16f8a2592b9a4e959374ab33ae7a697.
//
// Solidity: event SequencersUpdated(address indexed dataMarketAddress, address sequencerAddress, bool allowed)
func (_Contract *ContractFilterer) WatchSequencersUpdated(opts *bind.WatchOpts, sink chan<- *ContractSequencersUpdated, dataMarketAddress []common.Address) (event.Subscription, error) {

	var dataMarketAddressRule []interface{}
	for _, dataMarketAddressItem := range dataMarketAddress {
		dataMarketAddressRule = append(dataMarketAddressRule, dataMarketAddressItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "SequencersUpdated", dataMarketAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSequencersUpdated)
				if err := _Contract.contract.UnpackLog(event, "SequencersUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSequencersUpdated is a log parse operation binding the contract event 0xad12010237fe83915c67abe51836e693a16f8a2592b9a4e959374ab33ae7a697.
//
// Solidity: event SequencersUpdated(address indexed dataMarketAddress, address sequencerAddress, bool allowed)
func (_Contract *ContractFilterer) ParseSequencersUpdated(log types.Log) (*ContractSequencersUpdated, error) {
	event := new(ContractSequencersUpdated)
	if err := _Contract.contract.UnpackLog(event, "SequencersUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSnapshotBatchAttestationSubmittedIterator is returned from FilterSnapshotBatchAttestationSubmitted and is used to iterate over the raw logs and unpacked data for SnapshotBatchAttestationSubmitted events raised by the Contract contract.
type ContractSnapshotBatchAttestationSubmittedIterator struct {
	Event *ContractSnapshotBatchAttestationSubmitted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractSnapshotBatchAttestationSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSnapshotBatchAttestationSubmitted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractSnapshotBatchAttestationSubmitted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractSnapshotBatchAttestationSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSnapshotBatchAttestationSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSnapshotBatchAttestationSubmitted represents a SnapshotBatchAttestationSubmitted event raised by the Contract contract.
type ContractSnapshotBatchAttestationSubmitted struct {
	DataMarketAddress common.Address
	BatchCid          string
	EpochId           *big.Int
	Timestamp         *big.Int
	ValidatorAddr     common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterSnapshotBatchAttestationSubmitted is a free log retrieval operation binding the contract event 0x0b0cf2ab04f090685b5b2ac9b8beb1cc821e28298914a0ae86d397c79fe63c01.
//
// Solidity: event SnapshotBatchAttestationSubmitted(address indexed dataMarketAddress, string batchCid, uint256 indexed epochId, uint256 timestamp, address indexed validatorAddr)
func (_Contract *ContractFilterer) FilterSnapshotBatchAttestationSubmitted(opts *bind.FilterOpts, dataMarketAddress []common.Address, epochId []*big.Int, validatorAddr []common.Address) (*ContractSnapshotBatchAttestationSubmittedIterator, error) {

	var dataMarketAddressRule []interface{}
	for _, dataMarketAddressItem := range dataMarketAddress {
		dataMarketAddressRule = append(dataMarketAddressRule, dataMarketAddressItem)
	}

	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}

	var validatorAddrRule []interface{}
	for _, validatorAddrItem := range validatorAddr {
		validatorAddrRule = append(validatorAddrRule, validatorAddrItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "SnapshotBatchAttestationSubmitted", dataMarketAddressRule, epochIdRule, validatorAddrRule)
	if err != nil {
		return nil, err
	}
	return &ContractSnapshotBatchAttestationSubmittedIterator{contract: _Contract.contract, event: "SnapshotBatchAttestationSubmitted", logs: logs, sub: sub}, nil
}

// WatchSnapshotBatchAttestationSubmitted is a free log subscription operation binding the contract event 0x0b0cf2ab04f090685b5b2ac9b8beb1cc821e28298914a0ae86d397c79fe63c01.
//
// Solidity: event SnapshotBatchAttestationSubmitted(address indexed dataMarketAddress, string batchCid, uint256 indexed epochId, uint256 timestamp, address indexed validatorAddr)
func (_Contract *ContractFilterer) WatchSnapshotBatchAttestationSubmitted(opts *bind.WatchOpts, sink chan<- *ContractSnapshotBatchAttestationSubmitted, dataMarketAddress []common.Address, epochId []*big.Int, validatorAddr []common.Address) (event.Subscription, error) {

	var dataMarketAddressRule []interface{}
	for _, dataMarketAddressItem := range dataMarketAddress {
		dataMarketAddressRule = append(dataMarketAddressRule, dataMarketAddressItem)
	}

	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}

	var validatorAddrRule []interface{}
	for _, validatorAddrItem := range validatorAddr {
		validatorAddrRule = append(validatorAddrRule, validatorAddrItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "SnapshotBatchAttestationSubmitted", dataMarketAddressRule, epochIdRule, validatorAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSnapshotBatchAttestationSubmitted)
				if err := _Contract.contract.UnpackLog(event, "SnapshotBatchAttestationSubmitted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSnapshotBatchAttestationSubmitted is a log parse operation binding the contract event 0x0b0cf2ab04f090685b5b2ac9b8beb1cc821e28298914a0ae86d397c79fe63c01.
//
// Solidity: event SnapshotBatchAttestationSubmitted(address indexed dataMarketAddress, string batchCid, uint256 indexed epochId, uint256 timestamp, address indexed validatorAddr)
func (_Contract *ContractFilterer) ParseSnapshotBatchAttestationSubmitted(log types.Log) (*ContractSnapshotBatchAttestationSubmitted, error) {
	event := new(ContractSnapshotBatchAttestationSubmitted)
	if err := _Contract.contract.UnpackLog(event, "SnapshotBatchAttestationSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSnapshotBatchFinalizedIterator is returned from FilterSnapshotBatchFinalized and is used to iterate over the raw logs and unpacked data for SnapshotBatchFinalized events raised by the Contract contract.
type ContractSnapshotBatchFinalizedIterator struct {
	Event *ContractSnapshotBatchFinalized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractSnapshotBatchFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSnapshotBatchFinalized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractSnapshotBatchFinalized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractSnapshotBatchFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSnapshotBatchFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSnapshotBatchFinalized represents a SnapshotBatchFinalized event raised by the Contract contract.
type ContractSnapshotBatchFinalized struct {
	DataMarketAddress common.Address
	EpochId           *big.Int
	BatchCid          common.Hash
	Timestamp         *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterSnapshotBatchFinalized is a free log retrieval operation binding the contract event 0x5ce21fc72041bd91ef828e11b07f6d1107a20642ea918f190e1e2029d85fedfe.
//
// Solidity: event SnapshotBatchFinalized(address indexed dataMarketAddress, uint256 indexed epochId, string indexed batchCid, uint256 timestamp)
func (_Contract *ContractFilterer) FilterSnapshotBatchFinalized(opts *bind.FilterOpts, dataMarketAddress []common.Address, epochId []*big.Int, batchCid []string) (*ContractSnapshotBatchFinalizedIterator, error) {

	var dataMarketAddressRule []interface{}
	for _, dataMarketAddressItem := range dataMarketAddress {
		dataMarketAddressRule = append(dataMarketAddressRule, dataMarketAddressItem)
	}
	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}
	var batchCidRule []interface{}
	for _, batchCidItem := range batchCid {
		batchCidRule = append(batchCidRule, batchCidItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "SnapshotBatchFinalized", dataMarketAddressRule, epochIdRule, batchCidRule)
	if err != nil {
		return nil, err
	}
	return &ContractSnapshotBatchFinalizedIterator{contract: _Contract.contract, event: "SnapshotBatchFinalized", logs: logs, sub: sub}, nil
}

// WatchSnapshotBatchFinalized is a free log subscription operation binding the contract event 0x5ce21fc72041bd91ef828e11b07f6d1107a20642ea918f190e1e2029d85fedfe.
//
// Solidity: event SnapshotBatchFinalized(address indexed dataMarketAddress, uint256 indexed epochId, string indexed batchCid, uint256 timestamp)
func (_Contract *ContractFilterer) WatchSnapshotBatchFinalized(opts *bind.WatchOpts, sink chan<- *ContractSnapshotBatchFinalized, dataMarketAddress []common.Address, epochId []*big.Int, batchCid []string) (event.Subscription, error) {

	var dataMarketAddressRule []interface{}
	for _, dataMarketAddressItem := range dataMarketAddress {
		dataMarketAddressRule = append(dataMarketAddressRule, dataMarketAddressItem)
	}
	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}
	var batchCidRule []interface{}
	for _, batchCidItem := range batchCid {
		batchCidRule = append(batchCidRule, batchCidItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "SnapshotBatchFinalized", dataMarketAddressRule, epochIdRule, batchCidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSnapshotBatchFinalized)
				if err := _Contract.contract.UnpackLog(event, "SnapshotBatchFinalized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSnapshotBatchFinalized is a log parse operation binding the contract event 0x5ce21fc72041bd91ef828e11b07f6d1107a20642ea918f190e1e2029d85fedfe.
//
// Solidity: event SnapshotBatchFinalized(address indexed dataMarketAddress, uint256 indexed epochId, string indexed batchCid, uint256 timestamp)
func (_Contract *ContractFilterer) ParseSnapshotBatchFinalized(log types.Log) (*ContractSnapshotBatchFinalized, error) {
	event := new(ContractSnapshotBatchFinalized)
	if err := _Contract.contract.UnpackLog(event, "SnapshotBatchFinalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSnapshotBatchSubmittedIterator is returned from FilterSnapshotBatchSubmitted and is used to iterate over the raw logs and unpacked data for SnapshotBatchSubmitted events raised by the Contract contract.
type ContractSnapshotBatchSubmittedIterator struct {
	Event *ContractSnapshotBatchSubmitted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractSnapshotBatchSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSnapshotBatchSubmitted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractSnapshotBatchSubmitted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractSnapshotBatchSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSnapshotBatchSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSnapshotBatchSubmitted represents a SnapshotBatchSubmitted event raised by the Contract contract.
type ContractSnapshotBatchSubmitted struct {
	DataMarketAddress common.Address
	BatchCid          string
	EpochId           *big.Int
	Timestamp         *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterSnapshotBatchSubmitted is a free log retrieval operation binding the contract event 0xcbd08c0021940d15cc732c670afbb27d5170795ee3a4ba0e47368452fae58c5d.
//
// Solidity: event SnapshotBatchSubmitted(address indexed dataMarketAddress, string batchCid, uint256 indexed epochId, uint256 timestamp)
func (_Contract *ContractFilterer) FilterSnapshotBatchSubmitted(opts *bind.FilterOpts, dataMarketAddress []common.Address, epochId []*big.Int) (*ContractSnapshotBatchSubmittedIterator, error) {

	var dataMarketAddressRule []interface{}
	for _, dataMarketAddressItem := range dataMarketAddress {
		dataMarketAddressRule = append(dataMarketAddressRule, dataMarketAddressItem)
	}

	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "SnapshotBatchSubmitted", dataMarketAddressRule, epochIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractSnapshotBatchSubmittedIterator{contract: _Contract.contract, event: "SnapshotBatchSubmitted", logs: logs, sub: sub}, nil
}

// WatchSnapshotBatchSubmitted is a free log subscription operation binding the contract event 0xcbd08c0021940d15cc732c670afbb27d5170795ee3a4ba0e47368452fae58c5d.
//
// Solidity: event SnapshotBatchSubmitted(address indexed dataMarketAddress, string batchCid, uint256 indexed epochId, uint256 timestamp)
func (_Contract *ContractFilterer) WatchSnapshotBatchSubmitted(opts *bind.WatchOpts, sink chan<- *ContractSnapshotBatchSubmitted, dataMarketAddress []common.Address, epochId []*big.Int) (event.Subscription, error) {

	var dataMarketAddressRule []interface{}
	for _, dataMarketAddressItem := range dataMarketAddress {
		dataMarketAddressRule = append(dataMarketAddressRule, dataMarketAddressItem)
	}

	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "SnapshotBatchSubmitted", dataMarketAddressRule, epochIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSnapshotBatchSubmitted)
				if err := _Contract.contract.UnpackLog(event, "SnapshotBatchSubmitted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSnapshotBatchSubmitted is a log parse operation binding the contract event 0xcbd08c0021940d15cc732c670afbb27d5170795ee3a4ba0e47368452fae58c5d.
//
// Solidity: event SnapshotBatchSubmitted(address indexed dataMarketAddress, string batchCid, uint256 indexed epochId, uint256 timestamp)
func (_Contract *ContractFilterer) ParseSnapshotBatchSubmitted(log types.Log) (*ContractSnapshotBatchSubmitted, error) {
	event := new(ContractSnapshotBatchSubmitted)
	if err := _Contract.contract.UnpackLog(event, "SnapshotBatchSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSnapshotFinalizedIterator is returned from FilterSnapshotFinalized and is used to iterate over the raw logs and unpacked data for SnapshotFinalized events raised by the Contract contract.
type ContractSnapshotFinalizedIterator struct {
	Event *ContractSnapshotFinalized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractSnapshotFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSnapshotFinalized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractSnapshotFinalized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractSnapshotFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSnapshotFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSnapshotFinalized represents a SnapshotFinalized event raised by the Contract contract.
type ContractSnapshotFinalized struct {
	DataMarketAddress common.Address
	EpochId           *big.Int
	EpochEnd          *big.Int
	ProjectId         string
	SnapshotCid       string
	Timestamp         *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterSnapshotFinalized is a free log retrieval operation binding the contract event 0x1746fd35c2b9c905f8e7ef34232f796acb536f2fd87f01d201f57fb338eab9a6.
//
// Solidity: event SnapshotFinalized(address indexed dataMarketAddress, uint256 indexed epochId, uint256 epochEnd, string projectId, string snapshotCid, uint256 timestamp)
func (_Contract *ContractFilterer) FilterSnapshotFinalized(opts *bind.FilterOpts, dataMarketAddress []common.Address, epochId []*big.Int) (*ContractSnapshotFinalizedIterator, error) {

	var dataMarketAddressRule []interface{}
	for _, dataMarketAddressItem := range dataMarketAddress {
		dataMarketAddressRule = append(dataMarketAddressRule, dataMarketAddressItem)
	}
	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "SnapshotFinalized", dataMarketAddressRule, epochIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractSnapshotFinalizedIterator{contract: _Contract.contract, event: "SnapshotFinalized", logs: logs, sub: sub}, nil
}

// WatchSnapshotFinalized is a free log subscription operation binding the contract event 0x1746fd35c2b9c905f8e7ef34232f796acb536f2fd87f01d201f57fb338eab9a6.
//
// Solidity: event SnapshotFinalized(address indexed dataMarketAddress, uint256 indexed epochId, uint256 epochEnd, string projectId, string snapshotCid, uint256 timestamp)
func (_Contract *ContractFilterer) WatchSnapshotFinalized(opts *bind.WatchOpts, sink chan<- *ContractSnapshotFinalized, dataMarketAddress []common.Address, epochId []*big.Int) (event.Subscription, error) {

	var dataMarketAddressRule []interface{}
	for _, dataMarketAddressItem := range dataMarketAddress {
		dataMarketAddressRule = append(dataMarketAddressRule, dataMarketAddressItem)
	}
	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "SnapshotFinalized", dataMarketAddressRule, epochIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSnapshotFinalized)
				if err := _Contract.contract.UnpackLog(event, "SnapshotFinalized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSnapshotFinalized is a log parse operation binding the contract event 0x1746fd35c2b9c905f8e7ef34232f796acb536f2fd87f01d201f57fb338eab9a6.
//
// Solidity: event SnapshotFinalized(address indexed dataMarketAddress, uint256 indexed epochId, uint256 epochEnd, string projectId, string snapshotCid, uint256 timestamp)
func (_Contract *ContractFilterer) ParseSnapshotFinalized(log types.Log) (*ContractSnapshotFinalized, error) {
	event := new(ContractSnapshotFinalized)
	if err := _Contract.contract.UnpackLog(event, "SnapshotFinalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractTriggerBatchResubmissionIterator is returned from FilterTriggerBatchResubmission and is used to iterate over the raw logs and unpacked data for TriggerBatchResubmission events raised by the Contract contract.
type ContractTriggerBatchResubmissionIterator struct {
	Event *ContractTriggerBatchResubmission // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractTriggerBatchResubmissionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractTriggerBatchResubmission)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractTriggerBatchResubmission)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractTriggerBatchResubmissionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractTriggerBatchResubmissionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractTriggerBatchResubmission represents a TriggerBatchResubmission event raised by the Contract contract.
type ContractTriggerBatchResubmission struct {
	DataMarketAddress common.Address
	EpochId           *big.Int
	BatchCid          common.Hash
	Timestamp         *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterTriggerBatchResubmission is a free log retrieval operation binding the contract event 0x826e6849ff24825cbaeb8adb637217c5a8ef9fa9d8cd09ae58c5223254c25408.
//
// Solidity: event TriggerBatchResubmission(address indexed dataMarketAddress, uint256 indexed epochId, string indexed batchCid, uint256 timestamp)
func (_Contract *ContractFilterer) FilterTriggerBatchResubmission(opts *bind.FilterOpts, dataMarketAddress []common.Address, epochId []*big.Int, batchCid []string) (*ContractTriggerBatchResubmissionIterator, error) {

	var dataMarketAddressRule []interface{}
	for _, dataMarketAddressItem := range dataMarketAddress {
		dataMarketAddressRule = append(dataMarketAddressRule, dataMarketAddressItem)
	}
	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}
	var batchCidRule []interface{}
	for _, batchCidItem := range batchCid {
		batchCidRule = append(batchCidRule, batchCidItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "TriggerBatchResubmission", dataMarketAddressRule, epochIdRule, batchCidRule)
	if err != nil {
		return nil, err
	}
	return &ContractTriggerBatchResubmissionIterator{contract: _Contract.contract, event: "TriggerBatchResubmission", logs: logs, sub: sub}, nil
}

// WatchTriggerBatchResubmission is a free log subscription operation binding the contract event 0x826e6849ff24825cbaeb8adb637217c5a8ef9fa9d8cd09ae58c5223254c25408.
//
// Solidity: event TriggerBatchResubmission(address indexed dataMarketAddress, uint256 indexed epochId, string indexed batchCid, uint256 timestamp)
func (_Contract *ContractFilterer) WatchTriggerBatchResubmission(opts *bind.WatchOpts, sink chan<- *ContractTriggerBatchResubmission, dataMarketAddress []common.Address, epochId []*big.Int, batchCid []string) (event.Subscription, error) {

	var dataMarketAddressRule []interface{}
	for _, dataMarketAddressItem := range dataMarketAddress {
		dataMarketAddressRule = append(dataMarketAddressRule, dataMarketAddressItem)
	}
	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}
	var batchCidRule []interface{}
	for _, batchCidItem := range batchCid {
		batchCidRule = append(batchCidRule, batchCidItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "TriggerBatchResubmission", dataMarketAddressRule, epochIdRule, batchCidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractTriggerBatchResubmission)
				if err := _Contract.contract.UnpackLog(event, "TriggerBatchResubmission", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTriggerBatchResubmission is a log parse operation binding the contract event 0x826e6849ff24825cbaeb8adb637217c5a8ef9fa9d8cd09ae58c5223254c25408.
//
// Solidity: event TriggerBatchResubmission(address indexed dataMarketAddress, uint256 indexed epochId, string indexed batchCid, uint256 timestamp)
func (_Contract *ContractFilterer) ParseTriggerBatchResubmission(log types.Log) (*ContractTriggerBatchResubmission, error) {
	event := new(ContractTriggerBatchResubmission)
	if err := _Contract.contract.UnpackLog(event, "TriggerBatchResubmission", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Contract contract.
type ContractUpgradedIterator struct {
	Event *ContractUpgraded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractUpgraded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractUpgraded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractUpgraded represents a Upgraded event raised by the Contract contract.
type ContractUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Contract *ContractFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*ContractUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &ContractUpgradedIterator{contract: _Contract.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Contract *ContractFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *ContractUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractUpgraded)
				if err := _Contract.contract.UnpackLog(event, "Upgraded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Contract *ContractFilterer) ParseUpgraded(log types.Log) (*ContractUpgraded, error) {
	event := new(ContractUpgraded)
	if err := _Contract.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractValidatorAttestationsInvalidatedIterator is returned from FilterValidatorAttestationsInvalidated and is used to iterate over the raw logs and unpacked data for ValidatorAttestationsInvalidated events raised by the Contract contract.
type ContractValidatorAttestationsInvalidatedIterator struct {
	Event *ContractValidatorAttestationsInvalidated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractValidatorAttestationsInvalidatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractValidatorAttestationsInvalidated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractValidatorAttestationsInvalidated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractValidatorAttestationsInvalidatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractValidatorAttestationsInvalidatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractValidatorAttestationsInvalidated represents a ValidatorAttestationsInvalidated event raised by the Contract contract.
type ContractValidatorAttestationsInvalidated struct {
	DataMarketAddress common.Address
	EpochId           *big.Int
	BatchCid          common.Hash
	Validator         common.Address
	Timestamp         *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterValidatorAttestationsInvalidated is a free log retrieval operation binding the contract event 0xace836e24d26220daa574b979a6c9f9614da0cf8dd180e3fbfe974cf21c01c2a.
//
// Solidity: event ValidatorAttestationsInvalidated(address indexed dataMarketAddress, uint256 indexed epochId, string indexed batchCid, address validator, uint256 timestamp)
func (_Contract *ContractFilterer) FilterValidatorAttestationsInvalidated(opts *bind.FilterOpts, dataMarketAddress []common.Address, epochId []*big.Int, batchCid []string) (*ContractValidatorAttestationsInvalidatedIterator, error) {

	var dataMarketAddressRule []interface{}
	for _, dataMarketAddressItem := range dataMarketAddress {
		dataMarketAddressRule = append(dataMarketAddressRule, dataMarketAddressItem)
	}
	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}
	var batchCidRule []interface{}
	for _, batchCidItem := range batchCid {
		batchCidRule = append(batchCidRule, batchCidItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "ValidatorAttestationsInvalidated", dataMarketAddressRule, epochIdRule, batchCidRule)
	if err != nil {
		return nil, err
	}
	return &ContractValidatorAttestationsInvalidatedIterator{contract: _Contract.contract, event: "ValidatorAttestationsInvalidated", logs: logs, sub: sub}, nil
}

// WatchValidatorAttestationsInvalidated is a free log subscription operation binding the contract event 0xace836e24d26220daa574b979a6c9f9614da0cf8dd180e3fbfe974cf21c01c2a.
//
// Solidity: event ValidatorAttestationsInvalidated(address indexed dataMarketAddress, uint256 indexed epochId, string indexed batchCid, address validator, uint256 timestamp)
func (_Contract *ContractFilterer) WatchValidatorAttestationsInvalidated(opts *bind.WatchOpts, sink chan<- *ContractValidatorAttestationsInvalidated, dataMarketAddress []common.Address, epochId []*big.Int, batchCid []string) (event.Subscription, error) {

	var dataMarketAddressRule []interface{}
	for _, dataMarketAddressItem := range dataMarketAddress {
		dataMarketAddressRule = append(dataMarketAddressRule, dataMarketAddressItem)
	}
	var epochIdRule []interface{}
	for _, epochIdItem := range epochId {
		epochIdRule = append(epochIdRule, epochIdItem)
	}
	var batchCidRule []interface{}
	for _, batchCidItem := range batchCid {
		batchCidRule = append(batchCidRule, batchCidItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "ValidatorAttestationsInvalidated", dataMarketAddressRule, epochIdRule, batchCidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractValidatorAttestationsInvalidated)
				if err := _Contract.contract.UnpackLog(event, "ValidatorAttestationsInvalidated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseValidatorAttestationsInvalidated is a log parse operation binding the contract event 0xace836e24d26220daa574b979a6c9f9614da0cf8dd180e3fbfe974cf21c01c2a.
//
// Solidity: event ValidatorAttestationsInvalidated(address indexed dataMarketAddress, uint256 indexed epochId, string indexed batchCid, address validator, uint256 timestamp)
func (_Contract *ContractFilterer) ParseValidatorAttestationsInvalidated(log types.Log) (*ContractValidatorAttestationsInvalidated, error) {
	event := new(ContractValidatorAttestationsInvalidated)
	if err := _Contract.contract.UnpackLog(event, "ValidatorAttestationsInvalidated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractValidatorsUpdatedIterator is returned from FilterValidatorsUpdated and is used to iterate over the raw logs and unpacked data for ValidatorsUpdated events raised by the Contract contract.
type ContractValidatorsUpdatedIterator struct {
	Event *ContractValidatorsUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractValidatorsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractValidatorsUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractValidatorsUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractValidatorsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractValidatorsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractValidatorsUpdated represents a ValidatorsUpdated event raised by the Contract contract.
type ContractValidatorsUpdated struct {
	DataMarketAddress common.Address
	ValidatorAddress  common.Address
	Allowed           bool
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterValidatorsUpdated is a free log retrieval operation binding the contract event 0x2a93c48a2a98c035cd37b0e8a3e25c15ce5dd0caa1fb0553603c249c37db24b8.
//
// Solidity: event ValidatorsUpdated(address indexed dataMarketAddress, address validatorAddress, bool allowed)
func (_Contract *ContractFilterer) FilterValidatorsUpdated(opts *bind.FilterOpts, dataMarketAddress []common.Address) (*ContractValidatorsUpdatedIterator, error) {

	var dataMarketAddressRule []interface{}
	for _, dataMarketAddressItem := range dataMarketAddress {
		dataMarketAddressRule = append(dataMarketAddressRule, dataMarketAddressItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "ValidatorsUpdated", dataMarketAddressRule)
	if err != nil {
		return nil, err
	}
	return &ContractValidatorsUpdatedIterator{contract: _Contract.contract, event: "ValidatorsUpdated", logs: logs, sub: sub}, nil
}

// WatchValidatorsUpdated is a free log subscription operation binding the contract event 0x2a93c48a2a98c035cd37b0e8a3e25c15ce5dd0caa1fb0553603c249c37db24b8.
//
// Solidity: event ValidatorsUpdated(address indexed dataMarketAddress, address validatorAddress, bool allowed)
func (_Contract *ContractFilterer) WatchValidatorsUpdated(opts *bind.WatchOpts, sink chan<- *ContractValidatorsUpdated, dataMarketAddress []common.Address) (event.Subscription, error) {

	var dataMarketAddressRule []interface{}
	for _, dataMarketAddressItem := range dataMarketAddress {
		dataMarketAddressRule = append(dataMarketAddressRule, dataMarketAddressItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "ValidatorsUpdated", dataMarketAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractValidatorsUpdated)
				if err := _Contract.contract.UnpackLog(event, "ValidatorsUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseValidatorsUpdated is a log parse operation binding the contract event 0x2a93c48a2a98c035cd37b0e8a3e25c15ce5dd0caa1fb0553603c249c37db24b8.
//
// Solidity: event ValidatorsUpdated(address indexed dataMarketAddress, address validatorAddress, bool allowed)
func (_Contract *ContractFilterer) ParseValidatorsUpdated(log types.Log) (*ContractValidatorsUpdated, error) {
	event := new(ContractValidatorsUpdated)
	if err := _Contract.contract.UnpackLog(event, "ValidatorsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
