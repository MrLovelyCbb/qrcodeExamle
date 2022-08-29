package common

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var Rdb *redis.Client
var Ctx = context.Background()

func InitRedis() {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     "192.168.2.15:6378",
		Password: "q1w2e3r4", // no password set
		DB:       0,          // use default DB
	})

	_, err := Rdb.Ping(Ctx).Result()
	if err != nil {
		fmt.Println("failed")
	}

	fmt.Println("connected redis success")
}
