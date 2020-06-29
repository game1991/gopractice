package initial

import (
	"log"

	"github.com/spf13/viper"
)

//InitConfig 配置文件初始化
func InitConfig() error{
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Fatal error config file: %s \n", err)
		return err
	}
    return nil
}
