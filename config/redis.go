package config

import (
	"context"
	"fmt"
	"os"

	"github.com/redis/go-redis/v9"
)

var Ctx = context.Background()

var RedisClient *redis.Client

func ConnectRedis() {

		client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})

	_,err := client.Ping(Ctx).Result()

	if err!=nil{
		panic(err)
	}

	fmt.Println("Redis Connected")

	RedisClient=client

}
