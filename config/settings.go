package config

import (
	"encoding/json"
	"os"
	"strconv"
	"time"

	rpchelper "github.com/powerloom/rpc-helper"

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
	BlockInterval               float64
	BlockOffset                 int
	StatePollingInterval        float64
	SlotSyncInterval            int
	PollingStaticStateVariables bool
	RedisFlushBatchSize         int
	RPCNodes                    []string `json:"rpc_nodes"`
	ArchiveNodes                []string `json:"archive_nodes"`
	MaxRetries                  int      `json:"max_retries"`
	RetryDelayMs                int      `json:"retry_delay_ms"`
	MaxRetryDelayS              int      `json:"max_retry_delay_s"`
	RequestTimeoutS             int      `json:"request_timeout_s"`
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

	statePollingInterval, err := strconv.ParseFloat(getEnv("STATE_POLLING_INTERVAL", "60"), 64)
	if err != nil {
		log.Fatalf("Invalid STATE_POLLING_INTERVAL value: %v", err)
	}

	slotSyncInterval, err := strconv.Atoi(getEnv("SLOT_SYNC_INTERVAL", "60"))
	if err != nil {
		log.Fatalf("Invalid SLOT_SYNC_INTERVAL value: %v", err)
	}

	redisFlushBatchSize, err := strconv.Atoi(getEnv("REDIS_FLUSH_BATCH_SIZE", "100"))
	if err != nil {
		log.Fatalf("Invalid REDIS_FLUSH_BATCH_SIZE value: %v", err)
	}

	// Parse RPC nodes from environment
	rpcNodesStr := getEnv("RPC_NODES", "[]")
	var rpcNodes []string
	if err := json.Unmarshal([]byte(rpcNodesStr), &rpcNodes); err != nil {
		log.Fatalf("Failed to parse RPC_NODES environment variable: %v", err)
	}

	// Parse archive nodes from environment
	archiveNodesStr := getEnv("ARCHIVE_NODES", "[]")
	var archiveNodes []string
	if err := json.Unmarshal([]byte(archiveNodesStr), &archiveNodes); err != nil {
		log.Fatalf("Failed to parse ARCHIVE_NODES environment variable: %v", err)
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
		RedisFlushBatchSize:         redisFlushBatchSize,
		RPCNodes:                    rpcNodes,
		ArchiveNodes:                archiveNodes,
		MaxRetries:                  getEnvInt("MAX_RETRIES", 3),
		RetryDelayMs:                getEnvInt("RETRY_DELAY_MS", 500),
		MaxRetryDelayS:              getEnvInt("MAX_RETRY_DELAY_S", 30),
		RequestTimeoutS:             getEnvInt("REQUEST_TIMEOUT_S", 30),
	}

	redisDB, redisDBParseErr := strconv.Atoi(getEnv("REDIS_DB", ""))
	if redisDBParseErr != nil || redisDB < 0 {
		log.Fatalf("Failed to parse REDIS_DB environment variable: %v", redisDBParseErr)
	}
	config.RedisDB = redisDB

	blockInterval, blockIntervalParseErr := strconv.ParseFloat(getEnv("BLOCK_INTERVAL", "5.0"), 64)
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

func getEnvInt(key string, defaultValue int) int {
	value := getEnv(key, "")
	if value == "" {
		return defaultValue
	}
	intValue, err := strconv.Atoi(value)
	if err != nil {
		log.Fatalf("Failed to parse %s environment variable: %v", key, err)
	}
	return intValue
}

func (s *Settings) ToRPCConfig() *rpchelper.RPCConfig {
	return &rpchelper.RPCConfig{
		Nodes: func() []rpchelper.NodeConfig {
			var nodes []rpchelper.NodeConfig
			for _, url := range s.RPCNodes {
				nodes = append(nodes, rpchelper.NodeConfig{URL: url})
			}
			return nodes
		}(),
		ArchiveNodes: func() []rpchelper.NodeConfig {
			var nodes []rpchelper.NodeConfig
			for _, url := range s.ArchiveNodes {
				nodes = append(nodes, rpchelper.NodeConfig{URL: url})
			}
			return nodes
		}(),
		MaxRetries:     s.MaxRetries,
		RetryDelay:     time.Duration(s.RetryDelayMs) * time.Millisecond,
		MaxRetryDelay:  time.Duration(s.MaxRetryDelayS) * time.Second,
		RequestTimeout: time.Duration(s.RequestTimeoutS) * time.Second,
	}
}
