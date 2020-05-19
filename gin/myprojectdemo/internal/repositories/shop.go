package repositories

import (
	"ganlei.github.com/gopractice/gin/myprojectdemo/internal/pkg/models"

	"github.com/jinzhu/gorm"
)

type Shoper interface {
	CreateShop(*models.Shop) error
}

type shop struct {
	orm *gorm.DB
}

//CreateShop 存储到数据库
func (s shop) CreateShop(shop *models.Shop) error {
	return s.orm.Create(shop).Error
}

//InitShop 实例化shop对象
func InitShop(db *gorm.DB) (Shoper, error) {
	return &shop{
		orm: db,
	}, nil
}
