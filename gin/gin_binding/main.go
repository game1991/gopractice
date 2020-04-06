package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Login struct {
	//binding修饰的字段,若接收值为空值,则报错,是必须字段
	User     string `form:"username" json:"user" uri:"user" xml:"username" binding:"required"`
	Password string `form:"password" json:"pwd" uri:"password" xml:"pwd" binding:"required"`
}

func main() {
	//创建路由
	r := gin.Default()

	/*
		//JSON绑定
		r.POST("/loginJSON", func(c *gin.Context) {
			//声明要接收的变量
			var json Login
			if err:=c.ShouldBindJSON(&json);err!=nil{
				//返回错误信息,gin.H封装了生成json数据的工具
				c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
				return
			}
			//判断用户名密码是否正确
			if json.User!="root"||json.Password!="admin"{
				c.JSON(http.StatusBadRequest,gin.H{"status":"304"})
				return
			}
		    c.JSON(http.StatusOK,gin.H{"status":"200"})
		})
		r.Run(":8000")
	*/

	/*
		//FORM绑定
		r.POST("/loginFORM", func(c *gin.Context) {
			//声明要接收的变量
			var form Login
			//Bind()默认解析并绑定form格式
			//根据请求头中的content-type自动推断
			if err:=c.Bind(&form);err!=nil{
				//返回错误信息,gin.H封装了生成json数据的工具
				c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
				return
			}
			//判断用户名密码是否正确
			if form.User!="root"||form.Password!="admin"{
				c.JSON(http.StatusBadRequest,gin.H{"status":"304"})
				return
			}
		    c.JSON(http.StatusOK,gin.H{"status":"200"})
		})
		r.Run(":8000")
	*/

	//URI绑定
	r.GET("/:user/:password", func(c *gin.Context) {
		//声明要接收的变量
		var login Login
		if err := c.ShouldBindUri(&login); err != nil {
			//返回错误信息,gin.H封装了生成json数据的工具
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		//判断用户名密码是否正确
		if login.User != "root" || login.Password != "admin" {
			c.JSON(http.StatusBadRequest, gin.H{"status": "304"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "200"})
	})
	r.Run(":8000")

}
