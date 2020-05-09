package api

import (
	"mydemo/routers"

	"github.com/gin-gonic/gin"
)

//RegisteredRoute Registered api to route
func RegisteredRoute(route *gin.RouterGroup) {
	v1Group := route.Group("v1")
	routers.CollectRouter(v1Group)
}
