package model

type GlobalResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type TokenResponse struct {
	Success bool               `json:"success"`
	Message string             `json:"message"`
	Data    *TokenDataResponse `json:"data"`
}

type TokenDataResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
