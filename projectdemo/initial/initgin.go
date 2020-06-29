package initial

import (
	"projectdemo/internal/app/middleware"
	"projectdemo/internal/app/shop"
	"projectdemo/internal/app/user"
	"projectdemo/routers"

	"log"

	"github.com/spf13/viper"
)

//InitEngine gin配置初始化
func InitEngine() {

	// 加载多个APP的路由配置
	routers.Include(middleware.Routers, user.Routers, shop.Routers)
	// 初始化路由
	r := routers.Init()
	port := viper.GetString("server.port")
	if port != "" {
		panic(r.Run(":" + port))
	}

	if err := r.Run(); err != nil {
		log.Fatal("[log.DEBUG]startup service failed, err:" + err.Error())
	}
}
