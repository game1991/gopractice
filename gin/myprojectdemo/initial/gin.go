package initial

import (
	"mydemo/api"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func InitEngine() {
	engine := gin.Default()
	API := engine.Group("api")
	{
		api.RegisteredRoute(API)
	}
	port := viper.GetString("server.port")
	if port != "" {
		panic(engine.Run(":" + port))
	}
	panic(engine.Run())
}
