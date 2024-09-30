package model

import "time"

type Post struct {
	ID        int        `json:"id"`
	Title     string     `json:"title"`
	Content   string     `json:"content"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type NewPost struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type UpdatePost struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type PostResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    *Post  `json:"data"`
}

type PostMultipleResponse struct {
	Success bool    `json:"success"`
	Message string  `json:"message"`
	Data    []*Post `json:"data"`
}
