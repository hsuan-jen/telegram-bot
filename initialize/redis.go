package initialize

import (
	"log"

	"github.com/go-redis/redis"
)

func NewRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1",
		Password: "", // no password set
		DB: 0, // use default DB
	})
	_, err := client.Ping().Result()
	if err != nil {
		log.Fatal("redis connect ping failed, err:", err.Error())
		return nil
	}
	return client
}
