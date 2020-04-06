package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func sayHello(c *gin.Context) {
	//c.JSON(200, gin.H{
	//	"message": "hello",
	//})
	name := c.Param("name")
	action := c.Param("action")
	c.String(http.StatusOK, fmt.Sprintf(name+" is "+action))

}

func main() {

	//默认模式,中间包括中间件logger,recovery
	r := gin.Default()
	r.GET("/hello/:name/*action", sayHello)

	r.GET("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"massage": "GET",
		})
	})

	r.POST("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"massage": "POST",
		})
	})

	r.PUT("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"massage": "PUT",
		})
	})

	r.DELETE("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"massage": "DELETE",
		})
	})

	r.Run() //监听在本地8080端启动服务
}
