package main

import (
	"ganlei.github.com/gopractice/gin/myprojectdemo/initial"
	"ganlei.github.com/gopractice/gin/myprojectdemo/internal/app/shop"
	"ganlei.github.com/gopractice/gin/myprojectdemo/internal/app/user"
	repo "ganlei.github.com/gopractice/gin/myprojectdemo/internal/repositories"

	"log"

	_ "github.com/go-sql-driver/mysql"
)

//程序入口函数
func main() {
	//初始化配置
	initial.InitConfig()

	db := initial.InitDB()
	curd, err := repo.InitStore(db)
	if err != nil {
		log.Fatal("[log.Fatal]" + err.Error())
		return
	}
	shopInterface, err := repo.InitShop(db)
	if err != nil {
		log.Fatal("[log.Fatal]" + err.Error())
		return
	}
	//user相关接口初始化
	user.InitCurd(curd)
	//shop相关接口初始化
	shop.InitShop(shopInterface)
	//gin启动初始化
	initial.InitEngine()

}
