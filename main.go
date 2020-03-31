package main

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	//_ "pkg.deepin.com/golang/libs/gorm"
)

type UserIndex struct {
	SQLEmail       sql.NullString `json:"-" gorm:"column:email;type:varchar(191);unique;"`
	SQLPhoneNumber sql.NullString `json:"-" gorm:"column:phone_number;unique;type:varchar(32)"`
}

type UniqueID struct {
	ID int

	UserIndex

	Username    string `json:"username" gorm:"type:varchar(191);not null;unique;index"`
	Email       string `form:"email" json:"email" gorm:"-"`
	PhoneNumber string `gorm:"-" form:"phone_number" json:"phone_number" `
	Region      string `json:"region" gorm:"type:varchar(128);not null;default:''"`

	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-"`
}

func main() {
	db, err := gorm.Open("mysql", "uosid:uosid@tcp(10.0.10.18:3306)/uosid?charset=utf8&parseTime=true")
	//db, err := gorm.Open("mysql", "power:Zp900704@tcp(rm-8vbu2us3722ae2sm21o.mysql.zhangbei.rds.aliyuncs.com:3306)/uosid?charset=utf8&parseTime=true")
	//db, err := gorm.Open("mysql", "root:Zp900704@tcp(192.168.31.200:3306)/uosid?charset=utf8&parseTime=true")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//db.SetLogger(gorm.Logger{revel.})
	db.LogMode(true)
	db.SingularTable(true)

	for i := 0; i < 1; i++ {
		us := make([]*UniqueID, 0, 1)
		for j := 0; j < 10; j++ {
			us = append(us, &UniqueID{
				Username:    "zhao" + strconv.Itoa(i*100+j),
				Email:       "zhao" + strconv.Itoa(i*100+j) + "@haiouyueyue.top",
				PhoneNumber: strconv.FormatInt(10000000000+int64(i*100+j), 10),
			})
		}
		//db.Create(us)
		fmt.Println(CreateOrderStrands(us, db))
	}
	//time.Sleep(10 * time.Second)
}

func CreateOrderStrands(orderStrands []*UniqueID, db *gorm.DB) (int64, error) {
	sqlStr := "INSERT INTO `unique_id`(`email`,`phone_number`,`username`,`region`,`created_at`,`updated_at`,`deleted_at`) VALUES "
	vals := []interface{}{}

	for _, row := range orderStrands {
		sqlStr += "(?,?,?,?,?,?,?),"
		vals = append(vals, row.Email, row.PhoneNumber, row.Username, row.Region, row.CreatedAt, row.UpdatedAt, row.DeletedAt)
	}

	sqlStr = strings.TrimSuffix(sqlStr, ",")
	stmt, err := db.DB().Prepare(sqlStr)
	fmt.Println(sqlStr)
	if err != nil {
		return 0, err
	}

	defer stmt.Close()
	fmt.Println(vals)

	res, err := stmt.Exec(vals...)
	if err != nil {
		return 0, err
	}

	affect, err := res.RowsAffected()
	if err != nil || affect == 0 {
		return 0, err
	}

	return affect, nil
}
