package initial

import (
	"ganlei.github.com/gopractice/gin/myprojectdemo/internal/app/shop"
	"ganlei.github.com/gopractice/gin/myprojectdemo/internal/app/user"
	"ganlei.github.com/gopractice/gin/myprojectdemo/routers"
	"log"

	"github.com/spf13/viper"
)

//gin配置初始化
func InitEngine() {

	// 加载多个APP的路由配置
	routers.Include(user.Routers, shop.Routers)
	// 初始化路由
	r := routers.Init()
	port := viper.GetString("server.port")
	if port != "" {
		panic(r.Run(":" + port))
	}

	if err := r.Run(); err != nil {
		log.Fatal("[log.Fatal]startup service failed, err:" + err.Error())
	}
}
