package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	router.POST("/form", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")
		//表单参数设置默认值
		types := c.DefaultPostForm("type", "anonymous")
		//多选框
		hobbys := c.PostFormArray("hobby")
		c.String(http.StatusOK,
			fmt.Sprintf("type is %s, username is %s ,password is %s,hobby is %v",
				types, username, password, hobbys))
		c.String(200, fmt.Sprintf("\n"))
		c.String(200, fmt.Sprintf("\n"))
		c.JSON(200, gin.H{
			"status":   "posted",
			"username": username,
			"password": password,
			"hobbys":   hobbys,
		})
	})
	router.Run(":8080")
}
