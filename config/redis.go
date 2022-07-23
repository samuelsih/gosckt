package config

import (
	"os"
	"time"

	"github.com/go-redis/redis/v8"
	_ "github.com/joho/godotenv/autoload"
)

func RedisSetup() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS"),
		Password: "",
		DB: 0,
		DialTimeout: 3 * time.Second,
		ReadTimeout: 3 * time.Second,
	})

	if rdb == nil {
		panic("cant create redis client")
	}

	return rdb
}