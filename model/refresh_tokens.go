package model

import "time"

type RefreshTokens struct {
	ID        int        `json:"id"`
	UserID    int        `json:"user_id"`
	TokenHash string     `json:"token_hash"`
	CreatedAt time.Time  `json:"created_at"`
	ExpiredAt time.Time  `json:"expired_at"`
	RevokedAt *time.Time `json:"revoked_at"`
}

type RefreshTokenInput struct {
	RefreshToken string `json:"refresh_token"`
}
