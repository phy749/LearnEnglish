package utils

import (
	"context"
	"github.com/redis/go-redis/v9" // hoặc github.com/go-redis/redis/v8 tuỳ phiên bản bạn dùng
	"time"
)

// Redis client & context được khởi tạo sẵn ở nơi khác (ví dụ: trong redisdb.go)
var Client *redis.Client
var Ctx = context.Background()

func SaveRefreshToken(userID string, token string, expiration time.Duration) error {
	return Client.Set(Ctx, "refresh_token:"+userID, token, expiration).Err()
}

func GetRefreshToken(userID string) (string, error) {
	return Client.Get(Ctx, "refresh_token:"+userID).Result()
}

func DeleteRefreshToken(userID string) error {
	return Client.Del(Ctx, "refresh_token:"+userID).Err()
}
