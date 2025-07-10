package db

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
)

var RDB *redis.Client
var Ctx = context.Background()

func ConnectRedis() {
	RDB = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", 
		DB:       0,
	})

	_, err := RDB.Ping(Ctx).Result()
	if err != nil {
		log.Fatalf(" Redis connection failed: %v", err)
	}

	log.Println("Redis connected successfully")
}
