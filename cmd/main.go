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

	// Fetch all slots once at startup
	prost.FetchAllSlots()

	var wg sync.WaitGroup

	wg.Add(2)
	go prost.MonitorEvents()          // Start monitoring events for updates
	go prost.StartPeriodicStateSync() // Start periodic state sync
	wg.Wait()
}
