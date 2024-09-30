package model

type GlobalResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type TokenResponse struct {
	Success bool    `json:"success"`
	Message string  `json:"message"`
	Token   *string `json:"token"`
}
