package routers

import (
	"mydemo/internal/manage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CollectRouter(group *gin.RouterGroup) {
	r := group.Group("user")
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "user",
		})
	})

	r.GET("/curd",manage.QueryUser)
	r.POST("/curd",manage.CreateUser)
	r.DELETE("/curd",manage.DeleteUser)
	r.PUT("/curd",manage.UpdateUser)
}
