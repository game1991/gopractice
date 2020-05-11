package user

import (
	"github.com/gin-gonic/gin"
)

func Routers(e *gin.Engine) {
	r := e.Group("user")
	// r.GET("/hello", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "user",
	// 	})
	// })

	r.POST("", CreateUser)

	r = r.Group("/:name")
	r.GET("", QueryUser)
	r.DELETE("", DeleteUser)
	r.PUT("", UpdateUser)
}

// http:

// head:
// 	url:
// 	method:
// 	COntext-type:
// body
// 	data:xml
// 	data:JSON
// 	data:stream
