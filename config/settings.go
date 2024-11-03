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
	RedisDb                     int
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
	for _, addr := range dataMarketAddressesList {
		dataMarketContractAddresses = append(dataMarketContractAddresses, common.HexToAddress(addr))
	}
	config := Settings{
		ClientUrl:                   getEnv("PROST_RPC_URL", ""),
		ContractAddress:             getEnv("PROTOCOL_STATE_CONTRACT", ""),
		RedisHost:                   getEnv("REDIS_HOST", ""),
		RedisPort:                   getEnv("REDIS_PORT", ""),
		SlackReportingUrl:           getEnv("SLACK_REPORTING_URL", ""),
		DataMarketAddresses:         dataMarketAddressesList,
		DataMarketContractAddresses: dataMarketContractAddresses,
	}

	// Check for any missing required environment variables and log errors
	missingEnvVars := []string{}
	if config.ClientUrl == "" {
		missingEnvVars = append(missingEnvVars, "PROST_RPC_URL")
	}
	if config.ContractAddress == "" {
		missingEnvVars = append(missingEnvVars, "PROTOCOL_STATE_CONTRACT")
	}
	if len(config.DataMarketAddresses) == 0 {
		missingEnvVars = append(missingEnvVars, "DATA_MARKET_ADDRESSES")
	}
	if getEnv("REDIS_DB", "") == "" {
		missingEnvVars = append(missingEnvVars, "REDIS_DB")
	}
	if config.RedisHost == "" {
		missingEnvVars = append(missingEnvVars, "REDIS_HOST")
	}
	if config.RedisPort == "" {
		missingEnvVars = append(missingEnvVars, "REDIS_PORT")
	}

	if len(missingEnvVars) > 0 {
		log.Fatalf("Missing required environment variables: %v", missingEnvVars)
	}

	redisDb, err := strconv.Atoi(getEnv("REDIS_DB", ""))
	if err != nil || redisDb < 0 {
		log.Fatalf("Invalid REDIS_DB value: %v", redisDb)
	}

	config.RedisDb = redisDb

	checkOptionalEnvVar(config.SlackReportingUrl, "SLACK_REPORTING_URL")

	SettingsObj = &config
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func checkOptionalEnvVar(value, key string) {
	if value == "" {
		log.Warnf("Optional environment variable %s is not set", key)
	}
}
