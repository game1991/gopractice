package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

//cmd控制台可以使用curl命令测试
/*
curl http://127.0.0.1:8080/v1/login?name=zhangsan
curl http://127.0.0.1:8080/v2/submit -X POST -d{"name":"game"}
*/

func main() {
	router := gin.Default()
	//路由组1,处理get请求
	v1 := router.Group("/v1")
	//{}是书写规范
	{
		v1.GET("/login", login)
		v1.GET("/submit", submit)
	}
	//路由组2,处理post请求
	v2 := router.Group("/v2")
	{
		v2.POST("/login", login)
		v2.POST("/submit", submit)
	}
	router.Run(":8080")
}

func login(c *gin.Context) {
	name := c.DefaultQuery("name", "jack")
	c.String(200, fmt.Sprintf("hello %s\n", name))
}

func submit(c *gin.Context) {
	name := c.DefaultQuery("name", "lily")
	c.String(200, fmt.Sprintf("hello %s\n", name))
}
