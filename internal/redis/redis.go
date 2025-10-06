package redis

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()
var rdb *redis.Client

func InitReddis() {
	rdb = redis.NewClient(&redis.Options{
		Addr: "redis-cache:6379",
		Password: "2242",
		DB: 0,
	})

	pong, err := rdb.Ping(ctx).Result()


	if err != nil {
		fmt.Println("Error conect redis", err)
		return 
	}

	fmt.Println("Ping ", pong)

}