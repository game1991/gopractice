package mylogger

import (
	"fmt"
	"time"
)

//往终端上输出日志

//实现开关控制，设置级别权限
type ConsoleLogger struct {
	Level LogLevel
}

//构造函数
func NewConsoleLogger(levelStr string) ConsoleLogger {
	level, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	return ConsoleLogger{
		Level: level,
	}
}

//开关函数，权限控制
func (c ConsoleLogger) enable(logLevel LogLevel) bool {
	return logLevel >= c.Level
}

//显示详细信息的函数
func (c ConsoleLogger)log(lv LogLevel, format string,arg ...interface{}) {
	if c.enable(lv) {
		msg := fmt.Sprintf(format, arg...)
		now := time.Now()
		funcName, fileName, lineNo := getInfo(3)
		fmt.Printf("[%s] [%s] [%s--%s:%d] %s\n", now.Format("2006-01-02 15:04:05"), getLogString(lv), fileName, funcName, lineNo, msg)
	}
}

func (c ConsoleLogger) Debug(format string,arg ...interface{}) {

		c.log(DEBUG, format,arg...)

}

func (c ConsoleLogger) Info(format string,arg ...interface{}) {

		c.log(INFO, format,arg...)

}

func (c ConsoleLogger) Warning(format string,arg ...interface{}) {

		c.log(WARNING, format,arg...)

}

func (c ConsoleLogger) Trace(format string,arg ...interface{}) {

		c.log(TRACE, format,arg...)

}

func (c ConsoleLogger) Error(format string,arg ...interface{}) {

		c.log(ERROR, format,arg...)

}

func (c ConsoleLogger) Fatal(format string,arg ...interface{}) {

		c. log(FATAL, format,arg...)

}
