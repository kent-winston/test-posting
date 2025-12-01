package controller

import (
	"myapp/model"
	"myapp/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AuthRefreshToken godoc
// @Summary Get new access token
// @Description Get new access token using refresh token
// @Tags Token
// @Accept json
// @Produce json
// @Param body body model.RefreshTokenInput true "Refresh token data"
// @Success 200 {object} model.TokenResponse
// @Failure 400 {object} model.TokenResponse
// @Failure 500 {object} model.TokenResponse
// @Router /auth/refresh [post]
func AuthRefreshToken(c *gin.Context) {
	var (
		input model.RefreshTokenInput
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

	data, _ := s.RefreshTokensGenerateAccessToken(c.Request.Context(), input.RefreshToken)
	s.Commit()

	c.JSON(http.StatusOK, &model.TokenResponse{
		Success: true,
		Message: "Success",
		Data:    data,
	})
}
