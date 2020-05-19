package user

import (
	"strings"

	"ganlei.github.com/gopractice/gin/myprojectdemo/internal/pkg/models"
	repo "ganlei.github.com/gopractice/gin/myprojectdemo/internal/repositories"
	"ganlei.github.com/gopractice/gin/myprojectdemo/response"

	"log"

	"github.com/gin-gonic/gin"
)

var store repo.Curd

//接口初始化
func InitCurd(curd repo.Curd) {
	store = curd
}

//QueryUser 查询所有用户
func QueryAllUser(c *gin.Context) {
	var user models.User
	data, err := store.UserList(&user)
	if err != nil {
		log.Fatal("log.Fatal" + err.Error())
		return
	}
	if len(data) < 0 {
		response.Fail(c, data, "query result is nil")
		return
	}
	response.Success(c, data, "query success")

}

//CreateUser 创建用户
func CreateUser(c *gin.Context) {
	name := c.PostForm("name")
	telephone := c.PostForm("telephone")
	email := c.PostForm("email")
	password := c.PostForm("password")
	shopParam := c.PostForm("shop")
	shopList := strings.Split(shopParam, ",")
	shops := make([]models.Shop, 0)
	for _, v := range shopList {
		tmp := models.Shop{Name: v}
		shops = append(shops, tmp)
	}
	user := &models.User{
		Name:      name,
		Telephone: telephone,
		Email:     email,
		Password:  password,
		Shops:     shops,
	}
	err := store.CreateUser(user)
	if err != nil {
		log.Fatal("log.Fatal" + err.Error())
		return
	}
	response.Success(c, user, "create suceed")
}

//DeleteUser 删除用户
func DeleteUser(c *gin.Context) {
	name := c.Param("name")
	if strings.TrimSpace(name) == "" {
		response.Fail(c, "param error", "delete failed")
		return
	}
	user := &models.User{Name: name}
	err := store.DeleteUser(user)
	if err != nil {
		log.Fatal("log.Fatal" + err.Error())
		return
	}
	response.Success(c, user, "delete suceed")

}

//UpdateUser 更新用户
func UpdateUser(c *gin.Context) {
	name := c.Param("name")
	if strings.TrimSpace(name) == "" {
		response.Fail(c, gin.H{"data": "data is nil"}, "Update failed")
		return
	}
	user := &models.User{Name: name}
	updateUser := &models.User{}
	_, err := store.UpdateUser(user, updateUser)
	if err != nil {
		log.Fatal("log.Fatal" + err.Error())
		return
	}
	response.Success(c, updateUser, "Update suceed")
}

//查询用户购买商品信息通过姓名
func QueryUserByName(c *gin.Context) {
	name := c.Param("name")
	if strings.TrimSpace(name) == "" {
		response.Fail(c, "param error", "query failed")
		return
	}
	user := models.User{Name: name}
	data, err := store.QueryUserByName(&user)
	if err != nil {
		log.Fatal("log.Fatal" + err.Error())
		return
	}
	if data == nil {
		response.Fail(c, data, "query result is nil")
		return
	}
	response.Success(c, data, "query success")

}
