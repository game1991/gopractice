package models

import "github.com/jinzhu/gorm"

//User 用户信息对象
type User struct {
	gorm.Model
	IP        string  `gorm:"column:ip;type:varchar(255)"`
	Name      string  `gorm:"type:varchar(50);not null;unique;index"`
	Telephone string  `gorm:"varchar(11);not null;unique;index"`
	Email     string  `gorm:"varchar(255);unique;index"`
	Password  string  `gorm:"size:255;not null"`
	Shops     []*Shop `gorm:"many2many:user_shops"`
}
