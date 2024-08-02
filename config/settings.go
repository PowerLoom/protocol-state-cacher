package config

import (
	log "github.com/sirupsen/logrus"
	"os"
)

var SettingsObj *Settings

type Settings struct {
	ClientUrl         string
	ContractAddress   string
	RedisHost         string
	RedisPort         string
	SlackReportingUrl string
	DataMarketAddress string
}

func LoadConfig() {
	config := Settings{
		ClientUrl:         getEnv("PROST_RPC_URL", ""),
		ContractAddress:   getEnv("PROTOCOL_STATE_CONTRACT", ""),
		RedisHost:         getEnv("REDIS_HOST", ""),
		RedisPort:         getEnv("REDIS_PORT", ""),
		SlackReportingUrl: getEnv("SLACK_REPORTING_URL", ""),
		DataMarketAddress: getEnv("DATA_MARKET_ADDRESS", ""),
	}

	// Check for any missing required environment variables and log errors
	missingEnvVars := []string{}
	if config.ClientUrl == "" {
		missingEnvVars = append(missingEnvVars, "PROST_RPC_URL")
	}
	if config.ContractAddress == "" {
		missingEnvVars = append(missingEnvVars, "PROTOCOL_STATE_CONTRACT")
	}
	if config.DataMarketAddress == "" {
		missingEnvVars = append(missingEnvVars, "DATA_MARKET_ADDRESS")
	}

	if len(missingEnvVars) > 0 {
		log.Fatalf("Missing required environment variables: %v", missingEnvVars)
	}

	checkOptionalEnvVar(config.SlackReportingUrl, "SLACK_REPORTING_URL")
	checkOptionalEnvVar(config.RedisHost, "REDIS_HOST")
	checkOptionalEnvVar(config.RedisPort, "REDIS_PORT")

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
