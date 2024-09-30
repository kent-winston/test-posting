package service

import (
	"context"
	"fmt"
	"myapp/model"
	"myapp/tools"
	"time"

	"gorm.io/gorm"
)

func (s *Service) PostCreate(ctx context.Context, input model.NewPost) (*model.Post, error) {
	post := model.Post{
		Title:     input.Title,
		Content:   input.Content,
		CreatedAt: time.Now().UTC(),
	}

	if err := s.DB.Model(&post).Omit("updated_at").Create(&post).Error; err != nil {
		panic(err)
	}

	return &post, nil
}

func (s *Service) PostUpdate(ctx context.Context, input model.UpdatePost) (*model.Post, error) {
	var (
		post model.Post
	)

	if err := s.DB.Model(&post).Scopes(tools.IsDeletedAtNull).Where("id = ?", input.ID).Updates(map[string]interface{}{
		"title":      input.Title,
		"content":    input.Content,
		"updated_at": time.Now().UTC(),
	}).Error; err != nil {
		panic(err)
	}

	return s.PostGetByID(ctx, input.ID)
}

func (s *Service) PostDeleteByID(ctx context.Context, id int) (string, error) {
	var (
		post model.Post
	)

	if err := s.DB.Model(&post).Scopes(tools.IsDeletedAtNull).Where("id = ?", id).Omit("updated_at").Update("deleted_at", time.Now().UTC()).Error; err != nil {
		panic(err)
	}

	return "Success", nil
}

func (s *Service) PostGetAll(ctx context.Context) ([]*model.Post, error) {
	var (
		posts []*model.Post
	)

	if err := s.DB.Model(&posts).Scopes(tools.IsDeletedAtNull).Find(&posts).Error; err != nil {
		panic(err)
	}

	return posts, nil
}

func (s *Service) PostGetByID(ctx context.Context, id int) (*model.Post, error) {
	var (
		post model.Post
	)

	if err := s.DB.Model(&post).Scopes(tools.IsDeletedAtNull).Where("id = ?", id).First(&post).Error; err == gorm.ErrRecordNotFound {
		panic(fmt.Errorf("post not found"))
	} else if err != nil {
		panic(err)
	}

	return &post, nil
}
