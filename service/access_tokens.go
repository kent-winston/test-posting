package service

import (
	"context"
	"myapp/model"
	"myapp/tools"
	"time"
)

func (s *Service) AccessTokensCreate(ctx context.Context, accessToken model.AccessTokens) (*model.AccessTokens, error) {
	if err := s.DB.Model(&accessToken).Create(&accessToken).Error; err != nil {
		panic(err)
	}

	return &accessToken, nil
}

func (s *Service) AccessTokensCheckExistByRawToken(ctx context.Context, token string) (bool, error) {
	var (
		accessToken model.AccessTokens
		count       int64
	)

	hashedToken := tools.HashSHA256(token)

	if err := s.DB.Model(&accessToken).Where("token_hash = ? AND expired_at > ? AND revoked_at IS NULL", hashedToken, time.Now().UTC()).Count(&count).Error; err != nil {
		panic(err)
	}

	if int(count) > 0 {
		return true, nil
	}

	return false, nil
}
