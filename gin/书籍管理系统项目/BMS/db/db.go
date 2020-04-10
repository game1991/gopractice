package db

import (
	"BMS/v1/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

var db *sqlx.DB

func InitDB() (err error) {
	dsn := "root:root@tcp(localhost:3306)/go_test?charset=utf8&parseTime=True&loc=Local"
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		log.Fatalf("mysql connect failed,err:%v\n", err)
		return err
	}
	//设置最大连接数以及最大空闲连接
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(16)
	return
}

func QueryAllBook() (bookList []*models.Book, err error) {
	sqlStr := "select id,title,price from book"
	err = db.Select(&bookList, sqlStr)
	if err != nil {
		log.Fatalf("查询所有书籍信息错误,err:%v\n", err)
		return
	}
	return
}

func InsertBook(title string, price float64) (err error) {
	sqlStr := "insert into book(title,price) values(?,?)"
	_, err = db.Exec(sqlStr, title, price)
	if err != nil {
		log.Fatalf("插入书籍信息错误,err:%v\n", err)
		return
	}
	return
}

//删除书籍
func DeleteBook(id int) (err error) {
	sqlStr := "delete  from book where id =?"
	_, err = db.Exec(sqlStr, id)
	if err != nil {
		log.Fatalf("删除书籍信息失败,err:%v\n", err)
		return
	}
	return
}
