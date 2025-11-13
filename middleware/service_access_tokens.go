package middleware

import (
	"context"
	"myapp/config"
	"myapp/model"
	"myapp/tools"
	"time"
)

func AccessTokenCheckExistByRawToken(ctx context.Context, token string) (bool, error) {
	var (
		db          = config.GetDB()
		accessToken model.AccessTokens
		count       int64
	)

	hashedToken := tools.HashSHA256(token)

	if err := db.Model(&accessToken).Where("token_hash = ? AND expired_at > ? AND revoked_at IS NULL", hashedToken, time.Now().UTC()).Count(&count).Error; err != nil {
		return false, err
	}

	if int(count) > 0 {
		return true, nil
	}

	return false, nil
}
