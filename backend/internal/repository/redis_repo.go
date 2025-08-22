package repository

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

func NewRedisClient(addr string) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr: addr,
	})

	if err := rdb.Ping(context.Background()).Err(); err != nil {
		return nil, fmt.Errorf("failed to connect to redis %w", err)
	}

	log.Println("Successfully connected to redis.")
	return rdb, nil
}
