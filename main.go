package main

import (
	"Go-blog/core"
	"fmt"
	"Go-blog/global"
)

func main() {
	core.InitConf() //读取配置文件
	fmt.Println(global.Config) //打印配置文件内容
}