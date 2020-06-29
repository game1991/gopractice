package middleware

import (
	"net/http"

	"projectdemo/internal/pkg/middleware"

	"github.com/gin-gonic/gin"
)

//generateToken 生产token
func generateToken(c *gin.Context) {
	token, err := middleware.GenerateCsrfToken(c.ClientIP())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status_code": http.StatusBadRequest,
			"error":       err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status_code": http.StatusOK,
		"message":     token,
	})
}
