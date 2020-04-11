package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"pkg.deepin.io/ginessential/common"

	"pkg.deepin.io/ginessential/controller"
	"pkg.deepin.io/ginessential/model"
)

func main() {
	//数据库配置初始化
	db := common.InitDB()
	defer db.Close()

	r := gin.Default()
	r = CollectRouter(r)

	panic(engine.Run())

}
