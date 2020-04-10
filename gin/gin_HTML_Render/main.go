package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//HTML Render

func main() {
	engine := gin.Default()
	//加载静态模板文件
	engine.LoadHTMLGlob("templates/**/**")
	//设置静态文件的目录
	//第一个参数是代码里面使用的路径,第二个路径是实际保存静态文件的路径
	engine.Static("/dsb", "./statics")
	engine.GET("/login", loginHandler)
	engine.GET("/index", indexHandler)
	engine.Run(":9999")

}

func loginHandler(ctx *gin.Context) {
	//ctx.Request.FormValue("")
	ctx.HTML(http.StatusOK, "posts/login.html", gin.H{
		"msg": "呵呵",
	})
}

func indexHandler(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "users/index.html", gin.H{
		"msg": "哈哈",
	})
}
