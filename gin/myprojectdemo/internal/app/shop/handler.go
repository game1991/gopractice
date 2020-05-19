package shop

import (
	"fmt"
	"net/http"

	"ganlei.github.com/gopractice/gin/myprojectdemo/internal/pkg/models"
	repo "ganlei.github.com/gopractice/gin/myprojectdemo/internal/repositories"

	"ganlei.github.com/gopractice/gin/myprojectdemo/response"
	"github.com/gin-gonic/gin"
)

var store repo.Shoper

//接口初始化
func InitShop(shop repo.Shoper) {
	store = shop
}

func CreateShop(c *gin.Context) {
	name := c.PostForm("name")
	fmt.Println(name)
	shop := models.Shop{Name: name}
	if err := store.CreateShop(&shop); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	response.Success(c, gin.H{"CreateShop": "success"}, "success")
}
