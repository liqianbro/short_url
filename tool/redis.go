package tool

import (
	"log"

	"github.com/go-redis/redis"
)

var (
	RedisClient   *redis.Client
	redisHost     = "127.0.0.1"
	redisPort     = "6379"
	redisPassword = "123456"
)

func RedisInit() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     redisHost + ":" + redisPort,
		Password: redisPassword,
		DB:       1,
	})
	_, err := RedisClient.Ping().Result()
	if err != nil {
		panic("redis ping error")
	}
	log.Println("-----redis init success------")
}
