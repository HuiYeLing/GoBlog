package core

import (
	"Go-blog/global"
	"fmt"
	"log"
	"time"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

//用于初始化数据库连接

func InitGorm() *gorm.DB{
	if global.Config.Mysql.Host == "" {
		log.Println("未配置mysql,取消gorm连接")
		return nil
	}
	dsn := global.Config.Mysql.Dsn() //获取连接字符串

	var mysqlLogger logger.Interface
	if global.Config.System.Env == "debug"{ //判断当前环境是否是debug模式，如果是debug模式，就输出sql语句

		mysqlLogger = logger.Default.LogMode(logger.Info) //开发模式，输出sql语句
	}else{
		mysqlLogger = logger.Default.LogMode(logger.Error)//只打印错误的sql语句
	}
	db, err := gorm.Open(mysql.Open(dsn),&gorm.Config{
		Logger:mysqlLogger, 
	})
	if err!=nil{
		log.Fatal(fmt.Sprintf("[%s]mysql连接失败",dsn))//globakl.LOG是全局的日志对象，输出日志
	}
	sqlDB, _ := db.DB() //获取底层的sql.DB对象
	sqlDB.SetMaxIdleConns(10) //setMaxIdleConns是设置连接池中最大空闲连接数，默认值是0，表示不限制
	sqlDB.SetMaxOpenConns(100) //setMaxOpenConns是设置最大连接数，默认值是0，表示不限制
	sqlDB.SetConnMaxLifetime(time.Hour * 4) //连接最大的存活时间，如果连接超过这个时间，就会被关闭
	return db //返回连接对象
}

