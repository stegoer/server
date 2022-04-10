// inspired by https://gqlgen.com/reference/apq/#usage

// Package apq implement the graphql.Cache interface which is used
// for automatic persisted queries.
package apq

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

// Cache is a shared store for APQ and query AST caching.
type Cache struct {
	client redis.Client
	ttl    time.Duration
}

const (
	keyPrefix = "apq:"
	ttl       = time.Hour * 24
)

// NewCache returns a new Cache instance for APQ.
func NewCache(client redis.Client) *Cache {
	return &Cache{client: client, ttl: ttl}
}

// Add adds a value to the cache.
func (c *Cache) Add(ctx context.Context, key string, value interface{}) {
	c.client.Set(ctx, getKey(key), value, c.ttl)
}

// Get looks up a key's value from the cache.
func (c *Cache) Get(ctx context.Context, key string) (interface{}, bool) {
	res, err := c.client.Get(ctx, getKey(key)).Result()
	if err != nil {
		return nil, false
	}

	return res, true
}

func getKey(key string) string {
	return keyPrefix + key
}
