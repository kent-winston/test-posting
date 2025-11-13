package controller

import (
	"myapp/model"
	"myapp/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthRefreshToken(c *gin.Context) {
	var (
		input model.RefreshTokenInput
	)

	err := c.ShouldBind(&input)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, &model.TokenResponse{
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
