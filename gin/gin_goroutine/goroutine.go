package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func main() {
	r := gin.Default()
	//1.异步,在启动新的goroutine时,不能直接使用gin.context,而需要使用它的副本,否则会报错
	r.GET("/long_async", func(c *gin.Context) {
		//做个副本处理异步
		copyContext := c.Copy()
		go func() {
			time.Sleep(3 * time.Second)
			log.Println("异步执行:", copyContext.Request.URL.Path)
		}()
	})
	//2.同步执行
	r.GET("//long_sync", func(c *gin.Context) {
		time.Sleep(time.Second * 3)
		log.Println("同步执行:", c.Request.URL.Path)
	})
	r.Run(":9090")
}
