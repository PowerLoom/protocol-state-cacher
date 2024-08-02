package redis

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"protocol-state-cacher/config"
	"protocol-state-cacher/pkgs/common"
	"time"
)

var RedisClient *redis.Client

// TODO: Pool size failures to be checked
func NewRedisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%s", config.SettingsObj.RedisHost, config.SettingsObj.RedisPort), // Redis server address
		Password:     "",                                                                               // no password set
		DB:           0,
		PoolSize:     1000,
		ReadTimeout:  200 * time.Millisecond,
		WriteTimeout: 200 * time.Millisecond,
		DialTimeout:  5 * time.Second,
		IdleTimeout:  5 * time.Minute,
	})
}

func AddToSet(ctx context.Context, set string, keys ...string) error {
	if err := RedisClient.SAdd(ctx, set, keys).Err(); err != nil {
		common.SendFailureNotification("Redis error", err.Error(), time.Now().String(), "High")
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

func SetSubmission(ctx context.Context, key string, value string, set string, expiration time.Duration) error {
	if err := RedisClient.SAdd(ctx, set, key).Err(); err != nil {
		common.SendFailureNotification("Redis error", err.Error(), time.Now().String(), "High")
		return err
	}
	if err := RedisClient.Set(ctx, key, value, expiration).Err(); err != nil {
		common.SendFailureNotification("Redis error", err.Error(), time.Now().String(), "High")
		return err
	}
	return nil
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
			common.SendFailureNotification("Redis error", err.Error(), time.Now().String(), "High")
			return "", err
		}
	}
	return val, nil
}

func Set(ctx context.Context, key, value string, expiration time.Duration) error {
	return RedisClient.Set(ctx, key, value, expiration).Err()
}

// Save log to Redis
func SetProcessLog(ctx context.Context, key string, logEntry map[string]interface{}, exp time.Duration) error {
	data, err := json.Marshal(logEntry)
	if err != nil {
		common.SendFailureNotification("Redis SetProcessLog marshalling", err.Error(), time.Now().String(), "High")
		return fmt.Errorf("failed to marshal log entry: %w", err)
	}

	err = RedisClient.Set(ctx, key, data, exp).Err()
	if err != nil {
		common.SendFailureNotification("Redis error", err.Error(), time.Now().String(), "High")
		return fmt.Errorf("failed to set log entry in Redis: %w", err)
	}

	return nil
}
