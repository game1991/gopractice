package main

import (
	"log"

	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
)

//Account 账户信息结构体
type Account struct {
	Id      int
	Name    string `xorm:"unique"`
	Balance float64
	//Version int `xorm:"version"` //乐观锁
}

var x *xorm.Engine

func init() {
	var err error
	x, err = xorm.NewEngine("sqlite3", "./bank.db")
	if err != nil {
		log.Fatal("Fail to create engine,err:", err)
		return
	}
	if err = x.Sync2(new(Account)); err != nil {
		log.Fatal("Fail to sync database:", err)
		return
	}
}

func newAccount(name string, balance float64) error {
	_, err := x.Insert(&Account{Name: name, Balance: balance})
	return err
}

//查询计数
func getAccount() (int64, error) {
	return x.Count(new(Account))
}

// //迭代查询
//如果想要迭代查询某个表中符合条件的所有记录，则可以使用 Iterate 方法
// func ()  {

// }
