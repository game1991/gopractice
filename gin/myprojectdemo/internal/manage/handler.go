package manage

import (
	"mydemo/internal/pkg/models"
	repo "mydemo/internal/repositories"
	"mydemo/response"

	"github.com/gin-gonic/gin"
	"pkg.deepin.com/service/lib/log"
)

var (
	store repo.CURD
)

func QueryUser(c *gin.Context) {
	user:=&models.User{}
	data,err:=store.QueryUser(user)
	if err!=nil{
		log.Debug("log.DEBUG"+err.Error())
	}
    response.Success(c,gin.H{"data":data},"query sucess")
}

func CreateUser(c *gin.Context)  {
	
}

func DeleteUser(c *gin.Context)  {
	
}

func UpdateUser(c *gin.Context)  {
	
}
