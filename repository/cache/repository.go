package cache

import (
	"context"
	"crypto/tls"
	"net/url"
	"strings"

	"github.com/go-redis/redis/extra/redisotel"
	redis "github.com/go-redis/redis/v9"
)

// ---- redis
const defaultNS = "redis"

// Cache redis cache object
type Cache struct {
	client        *redis.Client
	ns            string
	clusterClient *redis.ClusterClient
}

func newRedisClient(uri *url.URL) (*redis.Client, error) {
	p, _ := url.User.Password()
	opt := &redis.Options{
		Addr:     url.Host,
		Password: p,
		DB:       0, //use default DB
	}
	if ts := url.Query().Get("tls"); ts != "" {
		opt.TLSConfig = &tls.Config{
			ServerName: ts,
		}
	}
	rClient := redis.NewClient(opt)
	rClient.AddHook(redisotel.TracingHook)
	ns := strings.TrimPrefix(url.Path, "/")
	if ns == "" {
		ns = defaultNS
	}

	cache := &Cache{
		client: rClient,
		ns:     strings.TrimPrefix(url.Path, "/"),
	}
	_, err := cache.client.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}
	return cache, nil
}
