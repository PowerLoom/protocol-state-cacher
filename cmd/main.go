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
	go prost.MonitorEvents()    // Start monitoring events for updates
	go prost.DynamicStateSync() // Start dynamic state sync
	if config.SettingsObj.PollingStaticStateVariables {
		go prost.StaticStateSync() // Start static state sync
	}
	go prost.SyncAllSlots() // Start syncing all slots
	wg.Wait()
}
