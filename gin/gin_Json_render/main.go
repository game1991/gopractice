package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type user struct {
	Name     string `json:"姓名"`
	Password string `json:"pwd"`
}

func main() {
	router := gin.Default()
	router.GET("index", indexHandler)
	router.Run(":8888")

}

func indexHandler(c *gin.Context) {
	//两种方法:1、自己拼接
	/*
		c.JSON(http.StatusOK,gin.H{
			"msg":"hello",
		})
	*/
	//2、使用结构体
	var u user
	u.Name = "甘大神在此..."
	u.Password = "110110"
	c.JSON(http.StatusOK, u)
}
