package main

import (
	"protocol-state-cacher/config"
	"protocol-state-cacher/pkgs/prost"
	"protocol-state-cacher/pkgs/redis"
	"protocol-state-cacher/pkgs/utils"
	"sync"
)

func main() {
	utils.InitLogger()
	config.LoadConfig()

	prost.ConfigureClient()
	prost.ConfigureContractInstance()
	redis.RedisClient = redis.NewRedisClient()

	wg := &sync.WaitGroup{}
	wg.Add(2)
	go prost.ColdSyncMappings()
	go prost.ColdSyncValues()
	wg.Wait()
}
