package cs_redis

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"strings"

	"github.com/go-redis/redis/v9"
	"github.com/joho/godotenv"
)

var RdbClient *redis.Client

func Init() {

	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	redisAddr, redisAddrExists := os.LookupEnv("REDIS_NODE")
	redisPort, redisPortExists := os.LookupEnv("REDIS_PORT")
	if !redisAddrExists || !redisPortExists || len(redisAddr) == 0 || len(redisPort) == 0 {
		log.Fatal("FATAL ERROR: ENV not properly configured, check .env file or REDIS_NODE and REDIS_PORT")
	}

	RdbClient = redis.NewClient(&redis.Options{
		Addr:     strings.Join([]string{redisAddr, ":", redisPort}, ""),
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

func PutMessageQueue(streamName, key string, value interface{}) {
	ctx := context.Background()
	vbyte, err := json.Marshal(value)
	if err != nil {
		log.Println("PutMessageQueue Marshal err:", value, err.Error())
	}
	result, err := RdbClient.XAdd(ctx, &redis.XAddArgs{Stream: streamName, Values: []string{key, string(vbyte)}}).Result()
	if err != nil {
		log.Println("PutMessageQueue xadd redis result, err:", result, err.Error())
	}
}
