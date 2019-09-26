package logger

import (
	"log"
	"os"
	"strconv"
	"io"
	"fmt"
)

var (
	Log    *log.Logger //输出
	Config *LogConfig  //日志输出配置
)

type LogConfig struct {
	ConsoleOut bool   //是否只是控制台输出 默认false
	FileOut    bool   //是否只是文件形式输出 默认false
	AllOut     bool   //是否控制台和文件都输出 默认true
	Level      string //log info level
	FilePath   string //输出的文件路径
}

//创建一个新的日志输出配置
func NewLog() *log.Logger {
	if Config.AllOut {
		file, err := os.OpenFile(Config.FilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalln("日志文件创建失败：", err)
		}
		//defer file.Close()
		Log = log.New(io.MultiWriter(file, os.Stderr),
			strconv.Itoa(os.Getpid())+" "+Config.Level+" ",
			log.Ldate|log.Lmicroseconds|log.Lshortfile)
	}

	if Config.ConsoleOut {
		Log = log.New(os.Stdout,
			strconv.Itoa(os.Getpid())+" "+Config.Level+" ",
			log.Ldate|log.Lmicroseconds|log.Lshortfile)
	}

	if Config.FileOut {
		file, err := os.OpenFile(Config.FilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalln("日志文件创建失败：", err)
		}
		//defer file.Close()
		Log = log.New(file,
			strconv.Itoa(os.Getpid())+" "+Config.Level+" ",
			log.Ldate|log.Lmicroseconds|log.Lshortfile)
	}
	return Log
}

func Info(msg string) {
	Log.Println(msg)
}

func Err(v ...interface{}) {
	Log.Fatalln(fmt.Sprint(v...))
}
