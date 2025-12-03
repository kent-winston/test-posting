package controller

import (
	"myapp/model"
	"myapp/service"
	"myapp/tools"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// PostCreate godoc
// @Summary Create post
// @Description Create post
// @Tags Post
// @Accept json
// @Produce json
// @Param body body model.NewPost true "Post data"
// @Param Authorization header string true "Bearer JWT token"
// @Success 200 {object} model.PostResponse
// @Failure 400 {object} model.PostResponse
// @Failure 500 {object} model.PostResponse
// @Router /post [post]
func PostCreate(c *gin.Context) {
	var (
		input model.NewPost
	)

	err := c.ShouldBind(&input)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, &model.PostResponse{
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
				code, message := tools.APIErrorResponse(err)
				c.AbortWithStatusJSON(code, &model.PostResponse{
					Success: false,
					Message: message,
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

// PostUpdate godoc
// @Summary Update post
// @Description Update post
// @Tags Post
// @Accept json
// @Produce json
// @Param body body model.UpdatePost true "Post data"
// @Param Authorization header string true "Bearer JWT token"
// @Success 200 {object} model.PostResponse
// @Failure 400 {object} model.PostResponse
// @Failure 500 {object} model.PostResponse
// @Router /post [put]
func PostUpdate(c *gin.Context) {
	var (
		input model.UpdatePost
	)

	err := c.ShouldBind(&input)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, &model.PostResponse{
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

// PostDelete godoc
// @Summary Delete post
// @Description Delete post
// @Tags Post
// @Accept json
// @Produce json
// @Param id query int true "Post id"
// @Param Authorization header string true "Bearer JWT token"
// @Success 200 {object} model.GlobalResponse
// @Failure 400 {object} model.GlobalResponse
// @Failure 500 {object} model.GlobalResponse
// @Router /post [delete]
func PostDelete(c *gin.Context) {
	postIDStr := c.Query("id")

	postID, err := strconv.Atoi(postIDStr)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, &model.GlobalResponse{
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

// PostGetAll godoc
// @Summary Get all posts
// @Description Get all posts from all user
// @Tags Post
// @Accept json
// @Produce json
// @Success 200 {object} model.PostMultipleResponse
// @Failure 500 {object} model.PostMultipleResponse
// @Router /posts [get]
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

// PostGetByID godoc
// @Summary Get post by id
// @Description Get post by id
// @Tags Post
// @Accept json
// @Produce json
// @Param id query int true "Post id"
// @Success 200 {object} model.PostResponse
// @Success 400 {object} model.PostResponse
// @Failure 500 {object} model.PostResponse
// @Router /post [get]
func PostGetByID(c *gin.Context) {
	postIDStr := c.Query("id")

	postID, err := strconv.Atoi(postIDStr)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, &model.PostResponse{
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
