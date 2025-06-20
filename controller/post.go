package controller

import (
	"myapp/model"
	"myapp/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// PostCreate godoc
// @Summary      Create a new post
// @Description  Creates a new post with JSON input which includes the `title` and `content` fields
// @Tags         posts
// @Accept       json
// @Produce      json
// @Param        post  body      model.NewPost       true  "Post creation payload"
// @Success      200   {object}  model.PostResponse
// @Failure      500   {object}  model.PostResponse
// @Router       /post/create [post]
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

// PostUpdate godoc
// @Summary      Update an existing post
// @Description  Updates an existing post with JSON input containing `id`, `title`, and `content` fields
// @Tags         posts
// @Accept       json
// @Produce      json
// @Param        post  body      model.UpdatePost       true  "Post update payload"
// @Success      200   {object}  model.PostResponse
// @Failure      500   {object}  model.PostResponse
// @Router       /post/update [put]
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

// PostDelete godoc
// @Summary      Delete a post by ID
// @Description  Deletes a post based on the `id` query parameter
// @Tags         posts
// @Produce      json
// @Param        id   query     int  true  "Post ID to delete"
// @Success      200   {object}  model.GlobalResponse
// @Failure      500   {object}  model.GlobalResponse
// @Router       /post/delete [delete]
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

// PostGetAll godoc
// @Summary      Show all posts
// @Description  returns a list of all posts
// @Tags         posts
// @Produce      json
// @Success      200  {object}  model.PostMultipleResponse
// @Failure      500  {object}  model.PostMultipleResponse
// @Router       /posts [get]
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
// @Summary      Show a post
// @Description  Return a single post by ID
// @Tags         posts
// @Produce      json
// @Param 		 id query int true "Post ID"
// @Success      200  {object}  model.PostResponse
// @Failure      500  {object}  model.PostResponse
// @Router       /post [get]
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
