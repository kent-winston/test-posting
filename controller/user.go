package controller

import (
	"myapp/model"
	"myapp/service"
	"myapp/tools"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserRegister godoc
// @Summary Register user account
// @Description Register account for user
// @Tags User
// @Accept json
// @Produce json
// @Param body body model.NewUser true "User registration data"
// @Success 200 {object} model.GlobalResponse
// @Failure 400 {object} model.GlobalResponse
// @Failure 500 {object} model.GlobalResponse
// @Router /user/register [post]
func UserRegister(c *gin.Context) {
	var (
		input model.NewUser
	)

	err := c.ShouldBind(&input)
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

	s.UserOnCreate(c.Request.Context(), input)
	s.Commit()

	c.JSON(http.StatusOK, &model.GlobalResponse{
		Success: true,
		Message: "Success",
	})
}

// UserLogin godoc
// @Summary Login user account
// @Description Login to user account
// @Tags User
// @Accept json
// @Produce json
// @Param body body model.UserLogin true "User login data"
// @Success 200 {object} model.TokenResponse
// @Failure 400 {object} model.TokenResponse
// @Failure 500 {object} model.TokenResponse
// @Router /user/login [post]
func UserLogin(c *gin.Context) {
	var (
		input model.UserLogin
	)

	err := c.ShouldBind(&input)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, &model.TokenResponse{
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
				c.AbortWithStatusJSON(http.StatusInternalServerError, &model.TokenResponse{
					Success: false,
					Message: err.Error(),
					Data:    nil,
				})
				return
			}
		}
	}()

	data, _ := s.UserLogin(c.Request.Context(), input)
	s.Commit()

	c.JSON(http.StatusOK, &model.TokenResponse{
		Success: true,
		Message: "Success",
		Data:    data,
	})
}

// UserGetMe godoc
// @Summary Login user account
// @Description Login to user account
// @Tags User
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer JWT token"
// @Success 200 {object} model.UserResponse
// @Failure 500 {object} model.UserResponse
// @Router /user/me [get]
func UserGetMe(c *gin.Context) {
	s := service.GetService()
	defer func() {
		if r := recover(); r != nil {
			err := s.ErrorCheck(r)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, &model.UserResponse{
					Success: false,
					Message: err.Error(),
					Data:    nil,
				})
				return
			}
		}
	}()

	user, _ := s.UserGetMe(c.Request.Context())

	c.JSON(http.StatusOK, &model.UserResponse{
		Success: true,
		Message: "Success",
		Data:    tools.UserToUserData(*user),
	})
}
