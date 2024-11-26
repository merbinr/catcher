package helpers

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var Ctx = context.Background()
var Redis_client *redis.Client

func LoadRedisClient() error {
	Redis_client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	_, err := Redis_client.Ping(Ctx).Result()
	if err != nil {
		return fmt.Errorf("unable to load redis client, error: %s", err)
	}
	return nil
}
