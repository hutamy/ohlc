package redis

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	Client *redis.Client
}

func NewRedisClient(ctx context.Context, host, password string, port int) (*RedisClient, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", host, port),
		Password: password,
		DB:       0,
	})

	log.Println("Redis Connection: ", client.Ping(ctx))
	if err := client.Ping(ctx).Err(); err != nil {
		log.Println("Redis Connection Error Log: ", err)
		return nil, err
	}

	return &RedisClient{
		Client: client,
	}, nil
}

func (r *RedisClient) Get(ctx context.Context, key string) (string, error) {
	return r.Client.Get(ctx, key).Result()
}

func (r *RedisClient) Close() error {
	return r.Client.Close()
}
