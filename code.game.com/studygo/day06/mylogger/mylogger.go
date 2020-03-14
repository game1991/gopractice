package mylogger

import (
	"errors"
	"fmt"
	"path"
	"runtime"
	"strings"
)

//自定义一个日志库
type LogLevel uint16

//Logger接口
type Logger interface {
	Debug(format string, arg ...interface{})
	Trace(format string, arg ...interface{})
	Info(format string, arg ...interface{})
	Warning(format string, arg ...interface{})
	Error(format string, arg ...interface{})
	Fatal(format string, arg ...interface{})
}


const (
	UNKNOWN LogLevel = iota
	DEBUG
	TRACE
	INFO
	WARNING
	ERROR
	FATAL
)



func parseLogLevel(s string) (LogLevel, error) {
	s = strings.ToLower(s)
	switch s {
	case "debug":
		return DEBUG, nil
	case "trace":
		return TRACE, nil
	case "info":
		return INFO, nil
	case "warning":
		return WARNING, nil
	case "error":
		return ERROR, nil
	case "fatal":
		return FATAL, nil
	default:
		err := errors.New("无效的日志级别")
		return UNKNOWN, err
	}
}

func getLogString(lv LogLevel) string {
	switch lv {
	case DEBUG:
		return "DEBUG"
	case TRACE:
		return "TRACE"
	case INFO:
		return "INFO"
	case WARNING:
		return "WARNING"
	case ERROR:
		return "ERROR"
	case FATAL:
		return "FATAL"
	default:
		return "DEBUG"
	}
}

func getInfo(skip int) (funcName, fileName string, lineNo int) {
	pc, file, lineNo, ok := runtime.Caller(skip)
	if !ok {
		fmt.Printf("runtime.Caller() failed\n")
		return
	}
	/*
		fmt.Println(file) // test_mylogger/mylogger.go
		fmt.Println(line) // 50
	*/
	funcName = runtime.FuncForPC(pc).Name()
	funcName=strings.Split(funcName,".")[1]
	fileName = path.Base(file)
	return
}
