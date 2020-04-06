package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/testdata/protoexample"
)

//以各种格式渲染

func main() {
	r := gin.Default()
	//1.json
	r.GET("/someJSON", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "someJSON", "status": "200"})
	})
	//2.结构体响应
	r.GET("/someStruct", func(c *gin.Context) {
		var msg struct {
			Name    string
			Message string
			Number  int
		}
		msg.Name = "root"
		msg.Message = "message"
		msg.Number = 123
		c.JSON(200, msg)
	})
	//3.xml
	r.GET("/someXML", func(c *gin.Context) {
		c.XML(200, gin.H{"message": "abc"})
	})
	//4.yaml
	r.GET("/someYAML", func(c *gin.Context) {
		c.YAML(200, gin.H{"name": "game"})
	})
	//5.protobuf格式
	r.GET("/someProtobuf", func(c *gin.Context) {
		resp := []int64{int64(1), int64(2)}
		//定义数据
		label := "label"
		//传Protobuf格式数据
		data := &protoexample.Test{Label: &label, Reps: resp}
		c.ProtoBuf(200, data)
	})
	r.Run(":9000")
}
