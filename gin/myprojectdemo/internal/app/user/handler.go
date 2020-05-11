package user

import (
	"strings"

	"mydemo/internal/pkg/models"
	repo "mydemo/internal/repositories"
	"mydemo/response"

	"github.com/gin-gonic/gin"
	"pkg.deepin.com/service/lib/log"
)

var store repo.Curd

func InitCurd(curd repo.Curd) {
	store = curd
}

func QueryUser(c *gin.Context) {
	name := c.Param("name")
	if strings.TrimSpace(name) == "" {
		response.Fail(c, gin.H{"data": "data is nil"}, "query failed")
		return
	}
	user := &models.User{Name: name}
	data, err := store.QueryUser(user)
	if err != nil {
		log.Debug("log.DEBUG" + err.Error())
		return
	}
	response.Success(c, gin.H{"data": data}, "query suceed")
}

func CreateUser(c *gin.Context) {
	name := c.PostForm("name")
	telephone := c.PostForm("telephone")
	email := c.PostForm("email")
	password := c.PostForm("password")

	user := &models.User{
		Name:      name,
		Telephone: telephone,
		Email:     email,
		Password:  password,
	}
	err := store.CreateUser(user)
	if err != nil {
		log.Debug("log.DEBUG" + err.Error())
		return
	}
	response.Success(c, gin.H{"data": "create user"}, "create suceed")
}

func DeleteUser(c *gin.Context) {
	name := c.Param("name")
	if strings.TrimSpace(name) == "" {
		response.Fail(c, gin.H{"data": "data is nil"}, "delete failed")
		return
	}
	user := &models.User{Name: name}
	err := store.DeleteUser(user)
	if err != nil {
		log.Debug("log.DEBUG" + err.Error())
		return
	}
	response.Success(c, gin.H{"data": name}, "delete suceed")

}

func UpdateUser(c *gin.Context) {
	name := c.Param("name")
	if strings.TrimSpace(name) == "" {
		response.Fail(c, gin.H{"data": "data is nil"}, "Update failed")
		return
	}
	user := &models.User{Name: name}
	updateUser := &models.User{}
	n, err := store.UpdateUser(user, updateUser)
	if err != nil {
		log.Debug("log.DEBUG" + err.Error())
		return
	}
	response.Success(c, gin.H{"len(data)": n}, "Update suceed")
}
