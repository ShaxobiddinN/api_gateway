package handlers

import (
	"blogpost/api_gateway/models"
	"blogpost/api_gateway/protogen/blogpost"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Authmiddleware...
func (h handler) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		hasAccessResponse, err := h.grpcClients.Auth.HasAccess(c.Request.Context(), &blogpost.TokenRequest{
			Token: token,
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.JSONErrorResponce{
				Error: err.Error(),
			})
			c.Abort()
			return
		}

		if !hasAccessResponse.HasAccess {
			c.JSON(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}
		c.Next()
	}
}
