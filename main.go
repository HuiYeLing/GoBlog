package main

import (
	"Go-blog/core"
	"Go-blog/global"
	// "fmt"

	// "github.com/sirupsen/logrus"
)

func main() {
	core.InitConf() //读取配置文件

	global.DB = core.InitGorm() //初始化数据库连接
	global.Log = core.InitLogger() //初始化日志对象
	global.Log.Warnln("嘻嘻嘻")
	global.Log.Error("嘻嘻嘻")
	global.Log.Infof("嘻嘻嘻")

	// logrus.Warnln("嘻嘻嘻")
	// logrus.Error("嘻嘻嘻")
	// logrus.Infof("嘻嘻嘻")
	// fmt.Println(global.DB)
}