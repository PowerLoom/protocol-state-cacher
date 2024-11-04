package redis

import (
	"fmt"
	"protocol-state-cacher/pkgs"
	"strings"
)

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
