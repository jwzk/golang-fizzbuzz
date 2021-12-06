package store

import (
	"context"
	"fmt"
	"os"

	"github.com/go-redis/redis/v8"
)

type Redis struct{
	Client redis.Client	
}

func New() *Redis {
	client := createRedisClient()
	return &Redis{Client: *client}
}

func (store Redis) GetDataFromKey(ctx context.Context, key string) string {
	data, err := store.Client.Get(ctx, key).Result()
	if err != nil {
		return ""
	}

	return data
}

func (store Redis) GetKeys(ctx context.Context, key string) ([]string, error) {
	keys, err := store.Client.Keys(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	return keys, nil
}

func (store Redis) Increment(ctx context.Context, key string) error {
	err := store.Client.Incr(ctx, key).Err()
	if err != nil {
		return err
	}

	return nil
}

func createRedisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")),
		Password: "",
		DB:       0,
	})
}
