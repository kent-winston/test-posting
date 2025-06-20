package controller

import (
	"myapp/model"
	"myapp/service"
	"myapp/tools"
	"net/http"

	"github.com/gin-gonic/gin"
)

//	UserRegister godoc
//
// @Summarry 		User Register
// @Description 	Register a use with name, email, and password
// @Tags 			users
// @Accept 			json
// @Produce 		json
// @Param 			input body model.NewUser true "New User Payload"
// @Success 		200 {object} model.User
// @Failure 		500 {object} model.GlobalResponse
// @Router 			/user/register [post]
func UserRegister(c *gin.Context) {
	var (
		input model.NewUser
	)

	err := c.ShouldBind(&input)
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

	s.UserOnCreate(c.Request.Context(), input)
	s.Commit()

	c.JSON(http.StatusOK, &model.GlobalResponse{
		Success: true,
		Message: "Success",
	})
}

// UserLogin godoc
// @Summary      User Login
// @Description  Authenticate a user and return a token
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        input body model.UserLogin true "User Login Payload"
// @Success      200 {object} model.TokenResponse
// @Failure      500 {object} model.TokenResponse
// @Router       /user/login [post]
func UserLogin(c *gin.Context) {
	var (
		input model.UserLogin
	)

	err := c.ShouldBind(&input)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, &model.TokenResponse{
			Success: false,
			Message: err.Error(),
			Token:   nil,
		})
		return
	}

	s := service.GetService()
	defer func() {
		if r := recover(); r != nil {
			err := s.ErrorCheck(r)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, &model.TokenResponse{
					Success: false,
					Message: err.Error(),
					Token:   nil,
				})
				return
			}
		}
	}()

	token, _ := s.UserLogin(c.Request.Context(), input)

	c.JSON(http.StatusOK, &model.TokenResponse{
		Success: true,
		Message: "Success",
		Token:   &token,
	})
}

// UserGetMe godoc
// @Summary      Get current user profile
// @Description  Returns the authenticated user's information
// @Tags         users
// @Security 	 BearerAuth
// @Produce      json
// @Success      200  {object}  model.UserResponse
// @Failure      500  {object}  model.UserResponse
// @Router       /user/me [get]
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
