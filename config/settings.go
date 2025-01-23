package config

import (
	"encoding/json"
	"os"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	log "github.com/sirupsen/logrus"
)

var SettingsObj *Settings

type Settings struct {
	ClientUrl                   string
	ContractAddress             string
	RedisHost                   string
	RedisPort                   string
	SlackReportingUrl           string
	DataMarketAddresses         []string
	DataMarketContractAddresses []common.Address
	RedisDB                     int
	BlockInterval               int
	BlockOffset                 int
	StatePollingInterval        int
	SlotSyncInterval            int
	PollingStaticStateVariables bool
}

func LoadConfig() {
	dataMarketAddresses := getEnv("DATA_MARKET_ADDRESSES", "[]")
	dataMarketAddressesList := []string{}
	dataMarketContractAddresses := []common.Address{}

	err := json.Unmarshal([]byte(dataMarketAddresses), &dataMarketAddressesList)
	if err != nil {
		log.Fatalf("Failed to parse DATA_MARKET_ADDRESSES environment variable: %v", err)
	}
	if len(dataMarketAddressesList) == 0 {
		log.Fatalf("DATA_MARKET_ADDRESSES environment variable has an empty array")
	}

	for _, dataMarketAddress := range dataMarketAddressesList {
		dataMarketContractAddresses = append(dataMarketContractAddresses, common.HexToAddress(dataMarketAddress))
	}

	pollingStaticStateVariables, pollingStaticStateVariablesErr := strconv.ParseBool(getEnv("POLLING_STATIC_STATE_VARIABLES", "true"))
	if pollingStaticStateVariablesErr != nil {
		log.Fatalf("Failed to parse POLLING_STATIC_STATE_VARIABLES environment variable: %v", pollingStaticStateVariablesErr)
	}

	statePollingInterval, err := strconv.Atoi(getEnv("STATE_POLLING_INTERVAL", "60"))
	if err != nil {
		log.Fatalf("Invalid STATE_POLLING_INTERVAL value: %v", err)
	}

	slotSyncInterval, err := strconv.Atoi(getEnv("SLOT_SYNC_INTERVAL", "60"))
	if err != nil {
		log.Fatalf("Invalid SLOT_SYNC_INTERVAL value: %v", err)
	}

	config := Settings{
		ClientUrl:                   getEnv("PROST_RPC_URL", ""),
		ContractAddress:             getEnv("PROTOCOL_STATE_CONTRACT", ""),
		RedisHost:                   getEnv("REDIS_HOST", ""),
		RedisPort:                   getEnv("REDIS_PORT", ""),
		SlackReportingUrl:           getEnv("SLACK_REPORTING_URL", ""),
		DataMarketAddresses:         dataMarketAddressesList,
		DataMarketContractAddresses: dataMarketContractAddresses,
		StatePollingInterval:        statePollingInterval,
		PollingStaticStateVariables: pollingStaticStateVariables,
		SlotSyncInterval:            slotSyncInterval,
	}

	redisDB, redisDBParseErr := strconv.Atoi(getEnv("REDIS_DB", ""))
	if redisDBParseErr != nil || redisDB < 0 {
		log.Fatalf("Failed to parse REDIS_DB environment variable: %v", redisDBParseErr)
	}
	config.RedisDB = redisDB

	blockInterval, blockIntervalParseErr := strconv.Atoi(getEnv("BLOCK_INTERVAL", "5"))
	if blockIntervalParseErr != nil {
		log.Fatalf("Failed to parse BLOCK_INTERVAL environment variable: %v", blockIntervalParseErr)
	}
	config.BlockInterval = blockInterval

	blockOffset, blockOffsetParseErr := strconv.Atoi(getEnv("BLOCK_OFFSET", "2"))
	if blockOffsetParseErr != nil {
		log.Fatalf("Failed to parse BLOCK_OFFSET environment variable: %v", blockOffsetParseErr)
	}
	config.BlockOffset = blockOffset

	SettingsObj = &config
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
