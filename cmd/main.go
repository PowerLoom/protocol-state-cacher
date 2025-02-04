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

	// Start dynamic state sync
	wg.Add(1)
	go prost.DynamicStateSync() // Start dynamic state sync

	// Start static state sync if enabled
	if config.SettingsObj.PollingStaticStateVariables {
		wg.Add(1)
		go prost.StaticStateSync() // Start static state sync
	}

	// Start syncing all slots
	wg.Add(1)
	go prost.SyncAllSlots() // Start syncing all slots

	wg.Wait()
}
