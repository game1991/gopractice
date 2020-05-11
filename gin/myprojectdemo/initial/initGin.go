package initial

import (
	"mydemo/internal/app/user"
	"mydemo/routers"

	"github.com/spf13/viper"
	"pkg.deepin.com/service/lib/log"
)

func InitEngine() {

	// 加载多个APP的路由配置
	routers.Include(user.Routers)
	// 初始化路由
	r := routers.Init()
	port := viper.GetString("server.port")
	if port != "" {
		panic(r.Run(":" + port))
	}

	if err := r.Run(); err != nil {
		log.Debug("[log.DEBUG]startup service failed, err:" + err.Error())
	}
}
