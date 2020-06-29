package middleware

import "github.com/gin-gonic/gin"

//Routers ..
func Routers(e *gin.Engine){
	r:=e.Group("middleware")
	r.POST("/token",generateToken)
}