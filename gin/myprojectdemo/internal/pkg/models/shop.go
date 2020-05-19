package models

import "github.com/jinzhu/gorm"

//Shop 商品信息对象
type Shop struct {
	gorm.Model
	Name  string  `gorm:"type:varchar(50)"`
	Users []*User `gorm:"many2many:user_shops"`
}
