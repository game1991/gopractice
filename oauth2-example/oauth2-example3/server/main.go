package main

import "github.com/gin-gonic/gin"

func main() {
	engine:=gin.Default()
	router:=engine.Group("oauth2")


	router.Any("auth", func(ctx *gin.Context) {
		ctx.Request.ParseForm()
		direct := oauth2.Authorize("abce", ctx.Request.Form)
		fmt.Println(direct)
		ctx.Redirect(http.StatusFound, direct)
	})
	router.POST("token", func(ctx *gin.Context) {
		state, info := oauth2.Token(ctx.Request)
		ctx.JSON(state, info)
	})
	g.Run(":8888")

}