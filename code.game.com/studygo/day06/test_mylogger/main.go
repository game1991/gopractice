package main

import (
	"code.game.com/studygo/day06/mylogger"
)

var log mylogger.Logger //声明一个全局的接口变量

//测试我们自己的日志库
func main()  {
	//log=mylogger.NewConsoleLogger("info") //终端日志实例
	log=mylogger.NewFileLogger("info","./","wenjianceshi.log",10*1024*518) //文件日志实例
	for {
		log.Debug("这是一条Debug日志")
		log.Trace("这是一条Trace日志")
		log.Info("这是一条Info日志")
		log.Warning("这是一条Warning日志")
		id := 110
		name := "大帅比"
		log.Error("这是一条Error日志,id:%d,name:%s", id, name)
		log.Fatal("这是一条Fatal日志")
		//ime.Sleep(time.Second)
	}
}
