package middleware

import (
	"context"
	"myapp/model"
	"myapp/tools"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

var CtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}

type User struct {
	ID int `json:"id"`
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authToken := c.Request.Header.Get("Authorization")
		if authToken == "" {
			c.Next()
			return
		}

		authTokens := strings.Split(authToken, " ")
		if authTokens[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, &model.GlobalResponse{
				Success: false,
				Message: "Unauthorized",
			})
			return
		}

		jwtToken, err := tools.TokenValidate(authTokens[1])
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, &model.GlobalResponse{
				Success: false,
				Message: "Invalid token",
			})
			return
		}

		claims, ok := jwtToken.Claims.(*tools.JwtClaim)
		if !ok || !jwtToken.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, &model.GlobalResponse{
				Success: false,
				Message: "Invalid token",
			})
			return
		}

		valid, err := AccessTokenCheckExistByRawToken(c.Request.Context(), authTokens[1])
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, &model.GlobalResponse{
				Success: false,
				Message: err.Error(),
			})
			return
		}

		if !valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, &model.GlobalResponse{
				Success: false,
				Message: "Invalid token",
			})
			return
		}

		ctx := context.WithValue(c.Request.Context(), CtxKey, &User{
			ID: claims.ID,
		})

		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func IsLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := AuthContext(c.Request.Context())
		if user == nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, &model.GlobalResponse{
				Success: false,
				Message: "Unauthorized",
			})
			return
		}

		c.Next()
	}
}

func AuthContext(ctx context.Context) *User {
	raw, _ := ctx.Value(CtxKey).(*User)
	return raw
}
