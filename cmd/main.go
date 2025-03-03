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

	// Create EventProcessor instance with a reasonable worker limit
	ep := prost.NewEventProcessor(10) // Adjust worker limit based on your needs

	var wg sync.WaitGroup
	routineCount := 3 // Base count for guaranteed routines
	if config.SettingsObj.PollingStaticStateVariables {
		routineCount++
	}
	wg.Add(routineCount)

	go func() {
		ep.MonitorEvents() // Use the EventProcessor instance
		wg.Done()
	}()
	go func() {
		prost.DynamicStateSync()
		wg.Done()
	}()
	if config.SettingsObj.PollingStaticStateVariables {
		go func() {
			prost.StaticStateSync()
			wg.Done()
		}()
	}
	go func() {
		prost.SyncAllSlots()
		wg.Done()
	}()

	wg.Wait()
}
