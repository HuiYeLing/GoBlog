package core

import (
	"Go-blog/global"
	"bytes"
	"fmt"
	"os"
	"path"

	"github.com/sirupsen/logrus"
)

//颜色常量
const (
	red = 31
	yellow = 33
	blue = 36
	gray = 37
)

type LogFormatter struct {

}

func (t *LogFormatter) Format (entry *logrus.Entry) ([]byte,error){ //这个功能的作用是将日志格式化成我们想要的格式
	var levelColor int //根据日志级别设置颜色
	switch entry.Level { //设置日志级别
	case logrus.DebugLevel, logrus.InfoLevel: //debug和info级别的日志使用默认颜色
		levelColor = gray 	//默认颜色
	case logrus.WarnLevel:	//警告级别的日志使用黄色
		levelColor = yellow 
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel://错误级别的日志使用红色
		levelColor = red
	default: //其他级别的日志使用蓝色
		levelColor = blue
	}
	var b *bytes.Buffer //b *bytes.Buffer就是设一个b指针，指向这个字节缓冲区
	//创建一个字节缓冲区,为什么要有这个缓冲区呢？
	//因为logrus的日志格式化器需要返回一个字节缓冲区，里面存放的是格式化后的日志内容
	//如果不使用缓冲区，就会频繁的创建和销毁内存，影响性能
	if entry.Buffer != nil { //如果缓冲区不为空，就使用缓冲区
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{} //否则就创建一个新的缓冲区
	}

	log := global.Config.Logger //获取全局配置的日志对象

	timestamp := entry.Time.Format("2006-01-02 15:04:05") //时间格式化
	if entry.HasCaller() { //entry.HasCaller()是logrus的一个方法，用来判断是否有调用栈信息
		funcVal := entry.Caller.Function //获取函数名
		fileVal := fmt.Sprintf("%s:%d", path.Base(entry.Caller.File), entry.Caller.Line) //获取文件名和行号
			// path.Base(entry.Caller.File), entry.Caller.Line
			// entry.Caller.File - 包含生成日志的文件的完整路径
			// path.Base() 函数 - 从路径中提取文件名部分
		fmt.Fprintf(b,"%s[%s] \x1b[%dm[%s]\x1b[0m %s %s %s\n",log.Prefix, timestamp, levelColor, entry.Level, fileVal, funcVal, entry.Message) 
		//# fmt.Fprintf 日志格式化参数详解
			// `"[%s]"` - 时间戳位置
			// `\x1b[%dm` - ANSI 颜色转义序列的开始部分
			// `[%s]` - 日志级别位置
			// `\x1b[0m` - ANSI 颜色转义序列的结束部分（重置颜色）
			// `%s %s %s\n` - 文件位置、函数名和日志消息的位置，末尾换行符
			//增加一个前缀log.Prefix，表示日志的前缀信息
	}else{
		//没有栈信息，就只打印时间戳、日志级别和日志消息
		fmt.Fprintf(b,"%s[%s] \x1b[%dm[%s]\x1b[0m %s\n", log.Prefix,timestamp, levelColor, entry.Level, entry.Message)
	}
	return b.Bytes(), nil //返回字节缓冲区的内容
	//b.Bytes()是将字节缓冲区转换为字节切片，返回给logrus使用
}

func InitLogger() *logrus.Logger{
	//功能描述：初始化日志对象，设置输出到标准输出，设置是否输出调用栈信息，设置日志格式化器，设置日志级别为debug
	mLog := logrus.New() //创建一个新的logrus实例
	mLog.SetOutput(os.Stdout) //设置输出到标准输出
	mLog.SetReportCaller(global.Config.Logger.ShowLine) //设置是否输出调用栈信息
	mLog.SetFormatter(&LogFormatter{}) //设置日志格式化器
	
	mLog.SetLevel(logrus.DebugLevel) //设置日志级别为debug
	InitDefaultLogger()
	return mLog //返回logrus实例
	//logrus实例就是一个日志对象，可以直接使用
}
func InitDefaultLogger(){ 
	//全局log
	//功能描述：初始化默认日志对象，设置输出到标准输出，设置是否输出调用栈信息，设置日志格式化器，设置日志级别为debug
	logrus.SetOutput(os.Stdout) //设置输出到标准输出
	logrus.SetReportCaller(global.Config.Logger.ShowLine) //设置是否输出调用栈信息
	logrus.SetFormatter(&LogFormatter{}) //设置自定义日志格式化器
	level, err := logrus.ParseLevel(global.Config.Logger.Level) //解析日志级别
	if err != nil {
		level = logrus.InfoLevel //如果解析失败，就使用info级别
	}
	logrus.SetLevel(level) //设置日志级别为debug
}