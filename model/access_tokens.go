package model

import "time"

type AccessTokens struct {
	ID             int        `json:"id"`
	UserID         int        `json:"user_id"`
	RefreshTokenID int        `json:"refresh_token_id"`
	TokenHash      string     `json:"token_hash"`
	CreatedAt      time.Time  `json:"created_at"`
	ExpiredAt      time.Time  `json:"expired_at"`
	RevokedAt      *time.Time `json:"revoked_at"`
}
