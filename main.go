package main

import (
	"fmt"
	core "portfolio-backend/init"

	"github.com/redis/go-redis/v9"
)

func main() {
	var redisClient *redis.Client
	redisClient = core.SetupRedis()
	fmt.Println("Hello world", redisClient)
}