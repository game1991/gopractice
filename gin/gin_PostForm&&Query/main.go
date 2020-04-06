package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)
/*
POST /post?id=1234&page=1 HTTP/1.1
Content-Type: application/x-www-form-urlencoded

name=game&message=this_is_great
*/

func main()  {
	router:=gin.New()
	router.POST("/post", func(ctx *gin.Context) {
		id := ctx.Query("id")
		page:=ctx.DefaultQuery("page","0")
		name := ctx.PostForm("name")
		message := ctx.PostForm("message")
		fmt.Printf("id:%s;page:%s;name:%s;message:%s",id,page,name,message)
	})
	router.Run(":8080")
}
