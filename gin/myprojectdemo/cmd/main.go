package main

import (
	"mydemo/initial"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	initial.InitConfig()

	initial.InitDB()

	initial.InitEngine()

}
