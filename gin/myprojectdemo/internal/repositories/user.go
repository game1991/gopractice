package repositories

import (
	"mydemo/internal/pkg/models"
	"mydemo/internal/pkg/myutil"

	"github.com/goutilpkg/phpbbencrypt"
	"github.com/jinzhu/gorm"
)

//CURD 实现增删改查
type CURD interface {
	QueryUser(*models.User) ([]*models.User, error)
	CreateUser(*models.User) error
	DeleteUser(*models.User) error
	UpdateUser(*models.User, *models.User) (int64, error)
}

type user struct {
	orm *gorm.DB
}

func (u *user) CreateUser(user *models.User) error {
	if user.Name == "" {
		user.Name = myutil.RandomString(8)
	}
	user.Password = phpbbencrypt.Encrypt(user.Password)
	result := u.orm.Create(user)
	return result.Error
}

func (u *user) DeleteUser(user *models.User) error {
	reult := u.orm.Delete(user)
	return reult.Error
}

func (u *user) UpdateUser(user *models.User, update *models.User) (int64, error) {
	reult := u.orm.Model(user).Where(user).Update(update)
	return reult.RowsAffected, reult.Error
}

func (u *user) QueryUser(user *models.User) ([]*models.User, error) {
	result := make([]*models.User, 0)
	return result, u.orm.Find(&result, user).Error
}

//InitStore 实例化user存储
func InitStore(db *gorm.DB) (CURD, error) {
	return &user{
		orm: db,
	}, nil
}