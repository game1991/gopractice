package mylogger

import (
	"fmt"
	"os"
	"path"
	"time"
)

//往文件里面写日志的相关代码

type FileLogger struct {
	Level       LogLevel
	filePath    string   //日志文件保存的路徑
	fileName    string   //日志文件保存的文件名
	fileObj     *os.File //存储文件
	errfileObj  *os.File //错误日志集合
	maxFileSize int64    //最大文件大小
	lastHour    int      //上一次的切割时间的小时数
}

//NewFileLogger构造函数
func NewFileLogger(levelStr, filePath, fileName string, maxFileSize int64) *FileLogger {
	level, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	fl := &FileLogger{
		Level:       level,
		filePath:    filePath,
		fileName:    fileName,
		maxFileSize: maxFileSize,
	}
	err = fl.initFile() //按照文件路径和文件名打开
	if err != nil {
		panic(err)
	}
	return fl
}

//根据指定的日志文件路径和文件名打开日志
func (f *FileLogger) initFile() error {
	fullFileName := path.Join(f.filePath, f.fileName)
	fileObj, err := os.OpenFile(fullFileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open log file failed,err:%v\n", err)
		return err
	}
	errfileObj, err := os.OpenFile(fullFileName+".err", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open err log file failed,err:%v\n", err)
		return err
	}
	f.fileObj = fileObj
	f.errfileObj = errfileObj
	return nil
}

//开关函数，权限控制 判断是否需要记录改日志
func (f *FileLogger) enable(logLevel LogLevel) bool {
	return logLevel >= f.Level
}

//文件切割前检查文件大小，是否达到切割要求
func (f *FileLogger) checkSize(file *os.File) bool {
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("get file info failed,err:%v\n", err)
		return false
	}
	//如果当前文件大小 大于等于日志文件的最大值，则返回true
	return fileInfo.Size() >= f.maxFileSize
}

//按小时数切割
func (f *FileLogger) checkHour(file *os.File) bool {
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("get file info failed,err:%v\n", err)
		return false
	}
	//在写日志之前检查一下当前时间的小时数和之前保存的是否一致，不一致就要切割
	 return fileInfo.ModTime().Hour() != f.lastHour
}

//文件切割
func (f *FileLogger) spiltFile(file *os.File) (*os.File, error) {
	//需要切割文件
	nowStr := time.Now().Format("20060102150405000")
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("get file info err,err is:%v\n", err)
		return nil, err
	}
	logName := path.Join(f.filePath, fileInfo.Name())      //拿到当前的日志文件完整路径
	newlogName := fmt.Sprintf("%s.bak%s", logName, nowStr) //拼接一个日志文件备份名
	//1、关闭当前日志文件
	file.Close()
	//2、备份一下rename xx.log ---> xx.log.bak202003122359
	os.Rename(logName, newlogName)
	//3、打开一个新的日志文件
	fileObj, err := os.OpenFile(logName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open new log file failed,err:%v\n", err)
		return nil, err
	}
	//4、将打开的新日志对象赋值给 f.fileObj
	return fileObj, nil
}

//记录日志的方法
func (f *FileLogger) log(lv LogLevel, format string, arg ...interface{}) {
	if f.enable(lv) {
		msg := fmt.Sprintf(format, arg...)
		now := time.Now()
		funcName, fileName, lineNo := getInfo(3)
		if f.checkSize(f.fileObj) {
			newFile, err := f.spiltFile(f.fileObj)
			if err != nil {
				fmt.Printf("spiltFile fileObj failed,err:%v\n", err)
				return
			}
			f.fileObj = newFile
		}
		fmt.Fprintf(f.fileObj, "[%s] [%s] [%s--%s:%d] %s\n", now.Format("2006-01-02 15:04:05"), getLogString(lv), fileName, funcName, lineNo, msg)
		if lv >= ERROR {
			//如果要记录的日志大于ERROR级别，则要需要在errfileObj中记录一遍
			if f.checkSize(f.errfileObj) {
				newFile, err := f.spiltFile(f.errfileObj)
				if err != nil {
					fmt.Printf("spiltFile errfileObj failed,err:%v\n", err)
					return
				}
				f.errfileObj = newFile
			}
			fmt.Fprintf(f.errfileObj, "[%s] [%s] [%s--%s:%d] %s\n", now.Format("2006-01-02 15:04:05"), getLogString(lv), fileName, funcName, lineNo, msg)
		}
	}
}

func (f *FileLogger) Debug(format string, arg ...interface{}) {

	f.log(DEBUG, format, arg...)

}

func (f *FileLogger) Info(format string, arg ...interface{}) {

	f.log(INFO, format, arg...)

}

func (f *FileLogger) Warning(format string, arg ...interface{}) {

	f.log(WARNING, format, arg...)

}

func (f *FileLogger) Trace(format string, arg ...interface{}) {

	f.log(TRACE, format, arg...)

}

func (f *FileLogger) Error(format string, arg ...interface{}) {

	f.log(ERROR, format, arg...)

}

func (f *FileLogger) Fatal(format string, arg ...interface{}) {

	f.log(FATAL, format, arg...)

}

func (f *FileLogger) Close() {
	f.fileObj.Close()
	f.errfileObj.Close()
}
