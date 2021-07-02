// +build bd_all bd_redis redis

package blackdatura

import (
	"context"
	"net/url"

	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

// Redis create redis sink instance
func Redis(addr, pwd, key string, db int) RedisSink {
	return RedisClient(key, redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pwd,
		DB:       db,
	}))
}

func (r RedisSink) check() RedisSink {
	if r.Client != nil && r.Client.Ping(context.Background()).Err() != nil {
		return r
	}
	panic("redis client error")
}

// RedisClient sink instance
func RedisClient(key string, client *redis.Client) RedisSink {
	return RedisSink{
		Client: client,
		Key:    key,
	}.check()
}

type RedisSink struct {
	Client *redis.Client
	Key    string
}

func (r RedisSink) Sink(*url.URL) (zap.Sink, error) { return r, nil }

// Close implement zap.Sink func Close
func (r RedisSink) Close() error { return r.Client.Close() }

// Write implement zap.Sink func Write
func (r RedisSink) Write(b []byte) (n int, err error) {
	return len(b), r.Client.Publish(context.Background(), r.Key, string(b)).Err()
}

// Sync implement zap.Sink func Sync
func (r RedisSink) Sync() error { return nil }

func (r RedisSink) String() string { return "redis" }
