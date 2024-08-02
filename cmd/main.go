package main

import (
	"protocol-state-cacher/config"
	"protocol-state-cacher/pkgs/prost"
	"protocol-state-cacher/pkgs/redis"
	"protocol-state-cacher/pkgs/utils"
)

func main() {
	utils.InitLogger()
	config.LoadConfig()

	prost.ConfigureClient()
	prost.ConfigureContractInstance()
	redis.RedisClient = redis.NewRedisClient()

	prost.ColdSync()
}
