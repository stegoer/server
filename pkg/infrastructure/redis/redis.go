package redis

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"

	"github.com/stegoer/server/pkg/infrastructure/env"
	"github.com/stegoer/server/pkg/infrastructure/log"
)

// Client wraps redis.Client.
type Client = redis.Client

// MustNew ensures that a new Client is created and panics if not.
func MustNew(config *env.Config, logger *log.Logger) *Client {
	client, err := New(config)
	if err != nil {
		logger.Panic(err)
	}

	return client
}

// New returns a new instance of redis.Client.
func New(config *env.Config) (*Client, error) {
	redisOptions, err := redis.ParseURL(config.RedisURL)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to parse %s as a redis url: %w",
			config.RedisURL,
			err,
		)
	}

	redisClient := redis.NewClient(redisOptions)

	_, err = redisClient.Ping(context.Background()).Result()
	if err != nil {
		return nil, fmt.Errorf("redis client is not responding: %w", err)
	}

	return redisClient, nil
}
