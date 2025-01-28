package main

import (
	"protocol-state-cacher/config"
	"protocol-state-cacher/pkgs/prost"
	"protocol-state-cacher/pkgs/redis"
	"protocol-state-cacher/pkgs/utils"
	"sync"
)

func main() {
	// Initiate logger
	utils.InitLogger()

	// Load the config object
	config.LoadConfig()

	// Initialize the RPC client, contract, and ABI instance
	prost.Initialize()

	// Setup redis
	redis.RedisClient = redis.NewRedisClient()

	// Set static state variables once
	prost.StaticStateVariables()

	var wg sync.WaitGroup

	wg.Add(4)

	// Start static state sync
	if config.SettingsObj.PollingStaticStateVariables {
		go prost.StaticStateSync()
	}

	// Start dynamic state sync
	go prost.DynamicStateSync()

	// Start monitoring events for updates
	go prost.MonitorEvents()

	// Start syncing all slots
	go prost.SyncAllSlots()

	wg.Wait()
}
