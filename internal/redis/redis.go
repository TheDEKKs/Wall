package redis

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

var Ctx = context.Background()
var Rdb *redis.Client


//Смысла коментировать я не вижу
func InitReddis() {
	Rdb = redis.NewClient(&redis.Options{
		Addr: "redis-cache:6379",
		Password: "2242",
		DB: 0,
	})

	_, err := Rdb.Ping(Ctx).Result()


	if err != nil {
		log.Println("Error conect redis", err)
		return 
	}

}