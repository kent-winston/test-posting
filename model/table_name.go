package model

func (t *AccessTokens) TableName() string {
	return "access_tokens"
}

func (t *Post) TableName() string {
	return "post"
}

func (t *User) TableName() string {
	return "user"
}

func (t *RefreshTokens) TableName() string {
	return "refresh_tokens"
}
