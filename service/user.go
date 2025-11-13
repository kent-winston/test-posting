package service

import (
	"context"
	"fmt"
	"myapp/middleware"
	"myapp/model"
	"myapp/tools"
	"strings"
	"time"

	"gorm.io/gorm"
)

func (s *Service) UserOnCreate(ctx context.Context, input model.NewUser) (*model.User, error) {
	if input.Email == "" || input.Password == "" {
		panic(fmt.Errorf("invalid email/password"))
	}

	input.Email = strings.ToLower(input.Email)
	input.Email = strings.TrimSpace(input.Email)

	valid := tools.CheckEmailValidity(input.Email)
	if !valid {
		panic(fmt.Errorf("invalid email/password"))
	}

	exist, _ := s.UserCheckExistByEmail(ctx, input.Email)
	if exist {
		panic(fmt.Errorf("email already used"))
	}

	password, err := tools.HashAndSalt(input.Password)
	if err != nil {
		panic(err)
	}

	user, _ := s.UserCreate(ctx, input, password)

	return user, nil
}

func (s *Service) UserCreate(ctx context.Context, input model.NewUser, password string) (*model.User, error) {
	user := model.User{
		Name:      input.Name,
		Email:     input.Email,
		Password:  password,
		CreatedAt: time.Now().UTC(),
	}

	if err := s.DB.Model(&user).Omit("updated_at").Create(&user).Error; err != nil {
		panic(err)
	}

	return &user, nil
}

func (s *Service) UserLogin(ctx context.Context, input model.UserLogin) (*model.TokenDataResponse, error) {
	if input.Email == "" || input.Password == "" {
		panic(fmt.Errorf("invalid email/password"))
	}

	input.Email = strings.ToLower(input.Email)
	input.Email = strings.TrimSpace(input.Email)

	valid := tools.CheckEmailValidity(input.Email)
	if !valid {
		panic(fmt.Errorf("invalid email/password"))
	}

	user, err := s.UserGetByEmailWithErrNotFound(ctx, input.Email)
	if err != nil {
		panic(fmt.Errorf("invalid email/password"))
	}

	valid, err = tools.CompareHash(user.Password, input.Password)
	if err != nil {
		panic(err)
	}

	if !valid {
		panic(fmt.Errorf("invalid email/password"))
	}

	return s.UserCreateAccessAndRefreshToken(ctx, user.ID)
}

func (s *Service) UserCreateAccessAndRefreshToken(ctx context.Context, id int) (*model.TokenDataResponse, error) {
	accessToken, accessExpiredAt := tools.TokenCreate(id)

	refreshToken, err := tools.GenerateSecureTokenHex(32)
	if err != nil {
		panic(err)
	}

	inputRefreshTokenData := model.RefreshTokens{
		UserID:    id,
		TokenHash: tools.HashSHA256(refreshToken),
		CreatedAt: time.Now().UTC(),
		ExpiredAt: time.Now().UTC().AddDate(0, 0, 7),
	}

	refreshTokenData, _ := s.RefreshTokensCreate(ctx, inputRefreshTokenData)

	inputAccessTokenData := model.AccessTokens{
		UserID:         id,
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

func (s *Service) UserGetMe(ctx context.Context) (*model.User, error) {
	var (
		getUser = middleware.AuthContext(ctx)
	)

	return s.UserGetByID(ctx, getUser.ID)
}

func (s *Service) UserGetByID(ctx context.Context, id int) (*model.User, error) {
	var (
		user model.User
	)

	if err := s.DB.Model(&user).Scopes(tools.IsDeletedAtNull).Where("id = ?", id).First(&user).Error; err == gorm.ErrRecordNotFound {
		panic(fmt.Errorf("user not found"))
	} else if err != nil {
		panic(err)
	}

	return &user, nil
}

func (s *Service) UserGetByEmailWithErrNotFound(ctx context.Context, email string) (*model.User, error) {
	var (
		user model.User
	)

	if err := s.DB.Model(&user).Scopes(tools.IsDeletedAtNull).Where("email = ?", email).First(&user).Error; err == gorm.ErrRecordNotFound {
		return nil, err
	} else if err != nil {
		panic(err)
	}

	return &user, nil
}

func (s *Service) UserCheckExistByEmail(ctx context.Context, email string) (bool, error) {
	var (
		user  model.User
		count int64
		exist bool = false
	)

	if err := s.DB.Model(&user).Scopes(tools.IsDeletedAtNull).Where("email = ?", email).Count(&count).Error; err != nil {
		panic(err)
	}

	if int(count) > 0 {
		exist = true
	}

	return exist, nil
}
