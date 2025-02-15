package redis

import (
	"fmt"
	"protocol-state-cacher/pkgs"
	"strings"
)

func AllSlotInfo() string {
	return pkgs.AllSlotsInfoKey
}

func TotalNodesCountKey() string {
	return pkgs.TotalNodesCount
}

func ContractStateVariable(varName string) string {
	return fmt.Sprintf("ProtocolState.%s", varName)
}

func ContractStateVariableWithDataMarket(dataMarketAddress string, varName string) string {
	return fmt.Sprintf("ProtocolState.%s.%s", dataMarketAddress, varName)
}

func SlotInfo(slotId string) string {
	return fmt.Sprintf("%s.%s", ContractStateVariable("SlotInfo"), slotId)
}

func CurrentEpochID(dataMarketAddress string) string {
	return fmt.Sprintf("%s.%s", strings.ToLower(dataMarketAddress), pkgs.CurrentEpochID)
}

func DataMarketCurrentDay(dataMarketAddress string) string {
	return fmt.Sprintf("%s.%s", pkgs.CurrentDayKey, strings.ToLower(dataMarketAddress))
}

func GetDailySnapshotQuotaTableKey() string {
	return pkgs.DailySnapshotQuotaTableKey
}