package service

import (
	"context"
	"fmt"
	"myapp/model"
	"myapp/tools"
	"time"

	"gorm.io/gorm"
)

func (s *Service) RefreshTokensCreate(ctx context.Context, refreshToken model.RefreshTokens) (*model.RefreshTokens, error) {
	if err := s.DB.Model(&refreshToken).Create(&refreshToken).Error; err != nil {
		panic(err)
	}

	return &refreshToken, nil
}

func (s *Service) RefreshTokensGenerateAccessToken(ctx context.Context, refreshToken string) (*model.TokenDataResponse, error) {
	refreshTokenData, _ := s.RefreshTokensGetByRawToken(ctx, refreshToken)

	accessToken, accessExpiredAt := tools.TokenCreate(refreshTokenData.UserID)

	inputAccessTokenData := model.AccessTokens{
		UserID:         refreshTokenData.UserID,
		RefreshTokenID: refreshTokenData.ID,
		TokenHash:      tools.HashSHA256(accessToken),
		CreatedAt:      time.Now().UTC(),
		ExpiredAt:      accessExpiredAt,
	}

	s.AccessTokensCreate(ctx, inputAccessTokenData)

	return &model.TokenDataResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *Service) RefreshTokensGetByRawToken(ctx context.Context, token string) (*model.RefreshTokens, error) {
	var (
		refreshToken model.RefreshTokens
	)

	hashedToken := tools.HashSHA256(token)

	if err := s.DB.Model(&refreshToken).Where("token_hash = ? AND expired_at > ? AND revoked_at IS NULL", hashedToken, time.Now().UTC()).Take(&refreshToken).Error; err == gorm.ErrRecordNotFound {
		panic(fmt.Errorf("invalid refresh token"))
	} else if err != nil {
		panic(err)
	}

	return &refreshToken, nil
}
