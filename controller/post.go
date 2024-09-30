package controller

import (
	"myapp/model"
	"myapp/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func PostCreate(c *gin.Context) {
	var (
		input model.NewPost
	)

	err := c.ShouldBind(&input)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, &model.PostResponse{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	s := service.GetTransaction()
	defer func() {
		if r := recover(); r != nil {
			err := s.Rollback(r)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, &model.PostResponse{
					Success: false,
					Message: err.Error(),
					Data:    nil,
				})
				return
			}
		}
	}()

	post, _ := s.PostCreate(c.Request.Context(), input)
	s.Commit()

	c.JSON(http.StatusOK, &model.PostResponse{
		Success: true,
		Message: "Success",
		Data:    post,
	})
}

func PostUpdate(c *gin.Context) {
	var (
		input model.UpdatePost
	)

	err := c.ShouldBind(&input)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, &model.PostResponse{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	s := service.GetTransaction()
	defer func() {
		if r := recover(); r != nil {
			err := s.Rollback(r)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, &model.PostResponse{
					Success: false,
					Message: err.Error(),
					Data:    nil,
				})
				return
			}
		}
	}()

	post, _ := s.PostUpdate(c.Request.Context(), input)
	s.Commit()

	c.JSON(http.StatusOK, &model.PostResponse{
		Success: true,
		Message: "Success",
		Data:    post,
	})
}

func PostDelete(c *gin.Context) {
	postIDStr := c.Query("id")

	postID, err := strconv.Atoi(postIDStr)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, &model.GlobalResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	s := service.GetTransaction()
	defer func() {
		if r := recover(); r != nil {
			err := s.Rollback(r)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, &model.GlobalResponse{
					Success: false,
					Message: err.Error(),
				})
				return
			}
		}
	}()

	resp, _ := s.PostDeleteByID(c.Request.Context(), postID)
	s.Commit()

	c.JSON(http.StatusInternalServerError, &model.GlobalResponse{
		Success: true,
		Message: resp,
	})
}

func PostGetAll(c *gin.Context) {
	s := service.GetService()
	defer func() {
		if r := recover(); r != nil {
			err := s.ErrorCheck(r)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, &model.PostMultipleResponse{
					Success: false,
					Message: err.Error(),
					Data:    nil,
				})
				return
			}
		}
	}()

	posts, _ := s.PostGetAll(c.Request.Context())

	c.JSON(http.StatusOK, &model.PostMultipleResponse{
		Success: true,
		Message: "Success",
		Data:    posts,
	})
}

func PostGetByID(c *gin.Context) {
	postIDStr := c.Query("id")

	postID, err := strconv.Atoi(postIDStr)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, &model.PostResponse{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	s := service.GetService()
	defer func() {
		if r := recover(); r != nil {
			err := s.ErrorCheck(r)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, &model.PostResponse{
					Success: false,
					Message: err.Error(),
					Data:    nil,
				})
				return
			}
		}
	}()

	post, _ := s.PostGetByID(c.Request.Context(), postID)

	c.JSON(http.StatusOK, &model.PostResponse{
		Success: true,
		Message: "Success",
		Data:    post,
	})
}
