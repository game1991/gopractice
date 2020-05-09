package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name      string `gorm:"type:varchar(50);not null;unique;index"`
	Telephone string `gorm:"varchar(11);not null;unique;index"`
	Email     string `gorm:"varchar(255);unique;index"`
	Password  string `gorm:"size:255;not null"`
}
