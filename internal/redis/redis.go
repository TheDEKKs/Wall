package redis

import (
	"context"
	"log"

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

	_, err := rdb.Ping(ctx).Result()


	if err != nil {
		log.Println("Error conect redis", err)
		return 
	}

}