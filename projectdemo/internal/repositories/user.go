package repositories

import (
	"projectdemo/internal/pkg/models"
	"projectdemo/internal/pkg/myutil"

	"github.com/goutilpkg/phpbbencrypt"
	"github.com/jinzhu/gorm"
)

//Curd 实现增删改查
type Curd interface {
	UserList(*models.User) ([]*models.User, error)
	QueryUserByName(*models.User) (*models.User, error)
	CreateUser(*models.User) error
	DeleteUser(*models.User) error
	UpdateUser(*models.User, *models.User) (int64, error)
}

type user struct {
	orm *gorm.DB
}

//创建用户
func (u *user) CreateUser(user *models.User) error {
	if user.Name == "" {
		user.Name = myutil.RandomString(8)
	}
	user.Password = phpbbencrypt.Encrypt(user.Password)
	result := u.orm.Create(user)
	return result.Error
}

//删除用户
func (u *user) DeleteUser(user *models.User) error {
	reult := u.orm.Delete(user)
	return reult.Error
}

//更新用户
func (u *user) UpdateUser(user *models.User, update *models.User) (int64, error) {
	reult := u.orm.Model(user).Where(user).Update(update)
	return reult.RowsAffected, reult.Error
}

//查询用户
func (u *user) QueryUserByName(user *models.User) (*models.User, error) {
	return user, u.orm.Where(&models.User{Name: user.Name}).First(&user).Error
}

//用户列表
func (u *user) UserList(user *models.User) ([]*models.User, error) {
	result := make([]*models.User, 0)
	return result, u.orm.Find(&result).Error
}

//InitStore 实例化user存储
func InitStore(db *gorm.DB) (Curd, error) {
	return &user{
		orm: db,
	}, nil
}
