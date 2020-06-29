package route

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"github.com/swaggo/gin-swagger"
)
//InitRouter ...
func InitRouter(e *gin.Engine) {
	url := ginSwagger.URL("http://localhost:9090/swagger/doc.json")
	e.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

}
