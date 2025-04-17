package global

import (
	"Go-blog/config" //配置文件路径

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var (
	Config *config.Config //全局配置变量，指向配置文件内容 第一个Config是变量名，*第二个Config是类型名，指向配置文件内容
	DB *gorm.DB 			//全局数据库连接变量，指向数据库连接对象
	Log *logrus.Logger 	//全局日志对象，指向日志对象
)