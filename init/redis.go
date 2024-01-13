package init

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func SetupRedis() *redis.Client {
	ctx := context.Background()

	RedisClient = redis.NewClient(&redis.Options{Addr: "localhost:6379", Password: "", DB: 0})
	pingResult, err := RedisClient.Ping(ctx).Result()

	if err != nil {
		panic(err)
	}
	fmt.Println("PING:", pingResult)
	return RedisClient
}