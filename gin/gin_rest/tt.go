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
	//Gin 中的 URL 参数解析的两种方式，分别是路径中的参数解析和查询字符串的参数解析,
	//下面的例子是路径解析,c.Query("key")是字符串解析,他是c.Request.URL.Query().Get("key")的简写
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

	r.GET("/redirect", func(c *gin.Context) {
		//支持内部和外部重定向
		c.Redirect(http.StatusMovedPermanently, "http://www.baidu.com")
	})

	r.Run() //监听在本地8080端启动服务
}
