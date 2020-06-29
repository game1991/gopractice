package user

import (
	"projectdemo/internal/pkg/middleware"

	"github.com/gin-gonic/gin"
)

//Routers 用户路由设定
func Routers(e *gin.Engine) {
	//用户相关的路由组
	r := e.Group("user")

	r.POST("", CreateUser)
	r.GET("", middleware.CsrfMiddleware(), QueryAllUser)

	r = r.Group("/:name")
	r.GET("", QueryUserByName)
	r.DELETE("", DeleteUser)
	r.PUT("", UpdateUser)
}
