package core

import (
	"Go-blog/config" //配置文件路径
	"Go-blog/global" 
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

//InitConf 读取yaml配置文件，需要下载yaml包
//命令：go get gopkg.in/yaml.v2
func InitConf() {
	const configFile = "settings.yaml" //要读取的配置文件路径
	c := &config.Config{} //读取配置文件，使用指针类型指向配置文件内容
	yamlConf, err := ioutil.ReadFile(configFile) //读取配置文件
	if err != nil { //读取配置文件失败
		panic(fmt.Errorf("get yamlConf error: %s", err))	
	}
	err = yaml.Unmarshal(yamlConf, c) //解析配置文件
	if err != nil { //解析配置文件失败
		log.Fatal("config Inint Unmarshal: %v", err) //解析配置文件失败,%v表示错误信息
	}
	log.Println("config yamlFile load Init success") //解析配置文件成功
	// fmt.Println(c)//打印配置文件内容
	global.Config = c //将配置文件内容赋值给全局变量Config
}