package cache

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/go-redis/redis/extra/redisotel"
	redis "github.com/go-redis/redis/v8"
)

// ---- redis
const defaultNS = "redis"
const NotFound = CacheError("[cache] not found")

// Cache redis cache object
type Cache struct {
	client        *redis.Client
	ns            string
	clusterClient *redis.ClusterClient
}
type CacheError string

type Option func(options *redis.Options)

type ClusterOption func(options *redis.ClusterOptions)

type DeleteCache struct {
	Pattern string
}
type DeleteOptions func(options *DeleteCache)

// NewRedisCache creating instance of redis cache
func NewRedisCache(ns string, option ...Option) (*Cache, error) {
	r := &redis.Options{}
	for _, o := range option {
		o(r)
	}
	rClient := redis.NewClient(r)
	rClient.AddHook(redisotel.TracingHook{})
	cache := &Cache{
		client: rClient,
		ns:     ns,
	}
	_, err := cache.client.Ping(context.Background()).Result()
	return cache, err
}

func NewRedisCluster(addresses []string, option ...ClusterOption) (*Cache, error) {
	r := &redis.ClusterOptions{}
	for _, o := range option {
		o(r)
	}
	rClient := redis.NewClusterClient(r)
	rClient.AddHook(redisotel.TracingHook{})

	cache := &Cache{
		clusterClient: rClient,
	}
	_, err := cache.clusterClient.Ping(context.Background()).Result()
	return cache, err
}

// Set set value
func (c *Cache) Set(ctx context.Context, key string, value interface{}, expiration int) error {
	switch value.(type) {
	case string, bool, int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64, []byte:
		return c.client.Set(ctx, c.ns+key, value, time.Duration(expiration)*time.Second).Err()
	default:
		b, err := json.Marshal(value)
		if err != nil {
			return err
		}
		return c.client.Set(ctx, c.ns+key, b, time.Duration(expiration)*time.Second).Err()
	}
}

// Get get value
func (c *Cache) Get(ctx context.Context, key string) ([]byte, error) {
	b, err := c.client.Get(ctx, c.ns+key).Bytes()
	if err != nil {
		if err == redis.Nil {
			return nil, errors.New(string(NotFound))
		}
		return nil, err
	}
	return b, nil
}

// GetObject get object value
func (c *Cache) GetObject(ctx context.Context, key string, doc interface{}) error {
	b, err := c.client.Get(ctx, c.ns+key).Bytes()
	if err != nil {
		if err == redis.Nil {
			return errors.New(string(NotFound))
		}
		return err
	}
	return json.Unmarshal(b, doc)
}

// GetString get string value
func (c *Cache) GetString(ctx context.Context, key string) (string, error) {
	s, err := c.client.Get(ctx, c.ns+key).Result()
	if err != nil {
		if err == redis.Nil {
			return "", errors.New(string(NotFound))
		}
		return "", err
	}
	return s, nil
}

// GetInt get int value
func (c *Cache) GetInt(ctx context.Context, key string) (int64, error) {
	i, err := c.client.Get(ctx, c.ns+key).Int64()
	if err != nil {
		if err == redis.Nil {
			return 0, errors.New(string(NotFound))
		}
		return 0, err
	}
	return i, nil
}

// GetFloat get float value
func (c *Cache) GetFloat(ctx context.Context, key string) (float64, error) {
	f, err := c.client.Get(ctx, c.ns+key).Float64()
	if err != nil {
		if err == redis.Nil {
			return 0, errors.New(string(NotFound))
		}
		return 0, err
	}
	return f, nil
}

// Exist check if key exist
func (c *Cache) Exist(ctx context.Context, key string) bool {
	return c.client.Exists(ctx, c.ns+key).Val() > 0
}

// Delete delete record
func (c *Cache) Delete(ctx context.Context, key string, opts ...DeleteOptions) error {
	deleteCache := &DeleteCache{}
	for _, opt := range opts {
		opt(deleteCache)
	}

	if deleteCache.Pattern != "" {
		return c.deletePattern(ctx, deleteCache.Pattern)
	}
	return c.client.Del(ctx, c.ns+key).Err()
}

func (c *Cache) GetKeys(ctx context.Context, pattern string) []string {
	cmd := c.client.Keys(ctx, pattern)
	keys, err := cmd.Result()
	if err != nil {
		return nil
	}
	return keys
}

// deletePattern delete record by pattern
func (c *Cache) deletePattern(ctx context.Context, pattern string) error {
	iter := c.client.Scan(ctx, 0, c.ns+pattern, 0).Iterator()
	var localKeys []string

	for iter.Next(ctx) {
		localKeys = append(localKeys, iter.Val())
	}

	if err := iter.Err(); err != nil {
		return err
	}

	if len(localKeys) > 0 {
		_, err := c.client.Pipelined(ctx, func(pipeline redis.Pipeliner) error {
			pipeline.Del(ctx, localKeys...)
			return nil
		})

		if err != nil {
			return err
		}
	}

	return nil
}

// RemainingTime get remaining time
func (c *Cache) RemainingTime(ctx context.Context, key string) int {
	return int(c.client.TTL(ctx, c.ns+key).Val().Seconds())
}

// Close close connection
func (c *Cache) Close() error {
	return c.client.Close()
}

func (c *Cache) As(i interface{}) bool {
	p, ok := i.(**redis.Client)
	if !ok {
		return false
	}
	*p = c.client
	return true
}

func (c *Cache) Flush(ctx context.Context) error {
	return c.client.FlushDBAsync(ctx).Err()
}
