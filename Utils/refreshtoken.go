package utils

import (
	"time"

	"github.com/phy749/LearnEnglish/config"
)

// Dùng Client & Context từ package config

func SaveRefreshToken(userID string, token string, expiration time.Duration) error {
	return config.Client.Set(config.Ctx, "refresh_token:"+userID, token, expiration).Err()
}

func GetRefreshToken(userID string) (string, error) {
	return config.Client.Get(config.Ctx, "refresh_token:"+userID).Result()
}

func DeleteRefreshToken(userID string) error {
	return config.Client.Del(config.Ctx, "refresh_token:"+userID).Err()
}
