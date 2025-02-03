package redis

import (
	"context"
	"errors"
	"fmt"
	"protocol-state-cacher/config"
	"protocol-state-cacher/pkgs/reporting"
	"time"

	"github.com/go-redis/redis/v8"
)

var RedisClient *redis.Client

// TODO: Pool size failures to be checked
func NewRedisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%s", config.SettingsObj.RedisHost, config.SettingsObj.RedisPort), // Redis server address
		Password:     "",                                                                               // no password set
		DB:           config.SettingsObj.RedisDB,
		PoolSize:     1000,
		ReadTimeout:  200 * time.Millisecond,
		WriteTimeout: 200 * time.Millisecond,
		DialTimeout:  5 * time.Second,
		IdleTimeout:  5 * time.Minute,
	})
}

func AddToSet(ctx context.Context, set string, keys ...string) error {
	if err := RedisClient.SAdd(ctx, set, keys).Err(); err != nil {
		reporting.SendFailureNotification("Redis error", err.Error(), time.Now().String(), "High")
		return err
	}

	return nil
}

func GetSetKeys(ctx context.Context, set string) []string {
	return RedisClient.SMembers(ctx, set).Val()
}

func RemoveFromSet(ctx context.Context, set, key string) error {
	return RedisClient.SRem(context.Background(), set, key).Err()
}

func Delete(ctx context.Context, set string) error {
	return RedisClient.Del(ctx, set).Err()
}

func PersistKey(ctx context.Context, key string) error {
	return RedisClient.Persist(ctx, key).Err()
}

func Get(ctx context.Context, key string) (string, error) {
	val, err := RedisClient.Get(ctx, key).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return "", nil
		} else {
			reporting.SendFailureNotification("Redis error", err.Error(), time.Now().String(), "High")
			return "", err
		}
	}
	return val, nil
}

func Set(ctx context.Context, key, value string, expiration time.Duration) error {
	return RedisClient.Set(ctx, key, value, expiration).Err()
}
