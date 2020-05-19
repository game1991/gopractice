package repositories

import (
	"fmt"
	"os"
	"testing"

	"ganlei.github.com/gopractice/gin/myprojectdemo/internal/pkg/models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	store Curd
	db    *gorm.DB
)

func init() {
	var err error
	db, err = gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/testforuser?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	err = db.DB().Ping()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	db.LogMode(true)
	db.AutoMigrate(&models.User{})
	store, _ = InitStore(db)
}

func TestCreateUser(t *testing.T) {
	db.Exec(`truncate table user`)
	user1 := &models.User{
		Name:      "gamer",
		Email:     "gan@163.com",
		Telephone: "13345678901",
		Password:  "12345asdf",
	}

	err := store.CreateUser(user1)

	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}

	//重名
	user2 := &models.User{
		Name:      "gamer",
		Email:     "gan1@163.com",
		Telephone: "13345678902",
		Password:  "12345asdf",
	}
	err = store.CreateUser(user2)
	if err == nil {
		t.Fail()
	} else {
		t.Log(err.Error())
	}

	//重邮箱
	err = store.CreateUser(&models.User{
		Name:      "gamer1",
		Email:     "gan@163.com",
		Telephone: "13345678900",
		Password:  "12345asdf",
	})
	if err == nil {
		t.Fail()
	} else {
		t.Log(err.Error())
	}

	//重手机号
	err = store.CreateUser(&models.User{
		Name:      "gamer2",
		Email:     "gan1@163.com",
		Telephone: "13345678901",
		Password:  "12345asdf",
	})
	if err == nil {
		t.Fail()
	} else {
		t.Log(err.Error())
	}
}
