package shop

import "github.com/gin-gonic/gin"

func Routers(e *gin.Engine){
	r:=e.Group("shop")
	r.POST("",CreateShop)
}