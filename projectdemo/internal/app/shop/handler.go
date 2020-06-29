package shop

import (
	"fmt"
	"net/http"

	"projectdemo/internal/pkg/models"
	repo "projectdemo/internal/repositories"
	"projectdemo/response"

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
		err = c.AbortWithError(http.StatusBadRequest, err)
		response.Fail(c, err, "create shop failed")
		return
	}
	response.Success(c, shop, "create shop success")
}
