package initial

import (
	"fmt"
	"mydemo/internal/pkg/models"

	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

func InitDB() *gorm.DB {
	drivename := viper.GetString("datasource.drivename")
	host := viper.GetString("datasource.host")
	port := viper.GetString("datasource.port")
	database := viper.GetString("datasource.database")
	username := viper.GetString("datasource.username")
	password := viper.GetString("datasource.password")
	charset := viper.GetString("datasource.charset")
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username, password, host, port, database, charset)
	db, err := gorm.Open(drivename, dns)
	if err != nil {
		panic("failed to connect database,err:" + err.Error())
	}
	db.AutoMigrate(&models.User{})

	return db
}
