package main

import (
	"log"
	"projectdemo/initial"
	"projectdemo/internal/app/shop"
	"projectdemo/internal/app/user"
	repo "projectdemo/internal/repositories"

	_ "github.com/go-sql-driver/mysql"
)

//程序入口函数
func main() {
	//初始化配置
	initial.InitConfig()

	db := initial.InitDB()
	curd, err := repo.InitStore(db)
	if err != nil {
		log.Fatal("[log.Debug]" + err.Error())
		return
	}
	shopInterface, err := repo.InitShop(db)
	if err != nil {
		log.Fatal("[log.Debug]" + err.Error())
		return
	}
	//user相关接口初始化
	user.InitCurd(curd)
	//shop相关接口初始化
	shop.InitShop(shopInterface)
	//gin启动初始化
	initial.InitEngine()

}
