package main

import (
	"mydemo/initial"
	"mydemo/internal/app/user"
	repo "mydemo/internal/repositories"

	_ "github.com/go-sql-driver/mysql"
	"pkg.deepin.com/service/lib/log"
)

func main() {
	initial.InitConfig()

	db := initial.InitDB()
	curd, err := repo.InitStore(db)
	if err != nil {
		log.Debug("[log.Debug]" + err.Error())
		return
	}
	user.InitCurd(curd)

	initial.InitEngine()

}
