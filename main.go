package main

import (
	"github.com/astaxie/beego"
	"os/exec"
	"shell/logger"
	"log"
	"github.com/axgle/mahonia"
)

var myLog *log.Logger

func init() {
	//日志输出位置
	logOut := beego.AppConfig.String("log")
	var fileOut, consoleOut, allOut bool
	if logOut == "file" {
		fileOut = true
	} else if logOut == "console" {
		consoleOut = true
	} else {
		allOut = true
	}
	logger.Config = &logger.LogConfig{
		ConsoleOut: consoleOut,
		FileOut:    fileOut,
		AllOut:     allOut,
		Level:      "DEBUG",
		FilePath:   "./cmd.log",
	}
	myLog = logger.NewLog()
}

//配置文件读取，执行脚本
func main() {
	run := beego.AppConfig.String("run")
	profile := beego.AppConfig.String("profile")
	myLog.Println("命令为：", run)
	myLog.Println("运行环境为：", profile)
	var cmd *exec.Cmd
	if profile=="windows"{
		cmd = exec.Command("cmd","/c",run)
	}else{
		cmd = exec.Command("/bin/bash","-c",run)
	}

	bytes, err := cmd.CombinedOutput()
	if err != nil {
		myLog.Println("run err：", err)
	}
	s := string(bytes)
	enc := mahonia.NewEncoder("GBK")
	output := enc.ConvertString(s)
	myLog.Println(output)
}
