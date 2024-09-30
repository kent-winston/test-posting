package model

func (t *Post) TableName() string {
	return "post"
}

func (t *User) TableName() string {
	return "user"
}
