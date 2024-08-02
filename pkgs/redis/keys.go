package redis

import (
	"fmt"
)

func ContractStateVariable(varName string) string {
	return fmt.Sprintf("ProtocolState.%s", varName)
}

func SlotInfo(slotId string) string {
	return fmt.Sprintf("%s.%s", ContractStateVariable("SlotInfo"), slotId)
}
