package cache

import (
	"demo/config"
	"fmt"
	"github.com/go-redis/redis"
)

var RedisClient *redis.Client

func RedisInit() {
	rConfig := config.Config.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", rConfig.RedisHost, rConfig.RedisPort),
		Password: rConfig.RedisPassword,
		DB:       rConfig.RedisDbName,
	})
	_, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}
	RedisClient = client
}
