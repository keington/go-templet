package cache

import (
	"context"
	"errors"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/keington/go-templet/pkg/zlog"
	"github.com/redis/go-redis/v9"
)

/**
 * @author: x.gallagher.anderson@gmail.com
 * @time: 2023/11/14 0:33
 * @file: cache.go
 * @description: cache
 */

type RedisConfig struct {
	Type             string
	Address          string
	Username         string
	Password         string
	DB               int
	MasterName       string
	SentinelUsername string
	SentinelPassword string
}

var (
	r    redis.Cmdable
	once sync.Once
)

type Redis redis.Cmdable

func NewRedis(cfg *RedisConfig) (Redis, error) {
	var redisClient Redis
	switch cfg.Type {
	case "standalone":
		redisOptions := &redis.Options{
			Addr:     cfg.Address,
			Username: cfg.Username,
			Password: cfg.Password,
			DB:       cfg.DB,
		}

		once.Do(func() {
			redisClient = redis.NewClient(redisOptions)
		})

	case "cluster":
		redisOptions := &redis.ClusterOptions{
			Addrs:    strings.Split(cfg.Address, ","),
			Username: cfg.Username,
			Password: cfg.Password,
		}

		once.Do(func() {
			redisClient = redis.NewClusterClient(redisOptions)
		})

	case "sentinel":
		redisOptions := &redis.FailoverOptions{
			MasterName:       cfg.MasterName,
			SentinelAddrs:    strings.Split(cfg.Address, ","),
			Username:         cfg.Username,
			Password:         cfg.Password,
			DB:               cfg.DB,
			SentinelUsername: cfg.SentinelUsername,
			SentinelPassword: cfg.SentinelPassword,
		}

		once.Do(func() {
			redisClient = redis.NewFailoverClient(redisOptions)
		})

	default:
		zlog.Errorf("failed to init cache , cache type is illegal: %s", cfg.Type)
		os.Exit(1)
	}

	err := redisClient.Ping(context.Background()).Err()
	if err != nil {
		zlog.Errorf("failed to ping cache: %s", err)
		os.Exit(1)
	}

	r = redisClient
	return redisClient, nil
}

// Get returns the value of key.
func Get(ctx context.Context, key string) []byte {
	val, err := r.Get(ctx, key).Bytes()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil
		}
		zlog.Errorf("failed to get key: %s, err: %s", key, err)
	}
	return val
}

// Set sets the value of key.
func Set(ctx context.Context, key string, value interface{}, expires time.Duration) error {
	return r.Set(ctx, key, value, expires).Err()
}

// Del deletes the key.
func Del(ctx context.Context, key string) error {
	return r.Del(ctx, key).Err()
}

// HMGet returns the value of field in the hash stored at key.
func HMGet(ctx context.Context, key string, fields ...string) [][]byte {
	var values [][]byte
	cmd := r.HMGet(ctx, key, fields...)
	if cmd.Err() != nil {
		zlog.Errorf("failed to get key: %s, err: %s", key, cmd.Err())
		return values
	}

	for _, val := range cmd.Val() {
		if val == nil {
			continue
		}
		values = append(values, []byte(val.(string)))
	}

	return values
}

// HMSet sets the value of field in the hash stored at key.
func HMSet(ctx context.Context, key string, fields map[string]interface{}) error {
	return r.HMSet(ctx, key, fields).Err()
}

// Expire sets a timeout on key.
func Expire(ctx context.Context, key string, expiration time.Duration) error {
	return r.Expire(ctx, key, expiration).Err()
}

// MGet returns the values of all specified keys.
func MGet(ctx context.Context, keys []string) [][]byte {
	var values [][]byte
	pipe := r.Pipeline()
	for _, key := range keys {
		pipe.Get(ctx, key)
	}
	command, _ := pipe.Exec(ctx)

	for i, key := range keys {
		cmd := command[i]
		if errors.Is(cmd.Err(), redis.Nil) {
			continue
		}

		if cmd.Err() != nil {
			zlog.Errorf("failed to get key: %s, err: %s", key, cmd.Err())
			continue
		}
		val := []byte(cmd.(*redis.StringCmd).Val())
		values = append(values, val)
	}

	return values
}

// MSet sets the given keys to their respective values.
func MSet(ctx context.Context, m map[string]interface{}) error {
	pipe := r.Pipeline()
	for k, v := range m {
		pipe.Set(ctx, k, v, 0)
	}
	_, err := pipe.Exec(ctx)
	return err
}
