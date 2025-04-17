package config

import "strconv"

type Mysql struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Config   string `yaml:"config"` //高级配置，例如charset，连接参数，charset=utf8mb4&parseTime=True&loc=Local
	DB       string `yaml:"db"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Loglevel string `yaml:"log_level"` //日志级别，debug竞赛输出全部sql，dev，release就是线上环境
}

func (m Mysql) Dsn() string {
	return m.User + ":" + m.Password + "@tcp(" + m.Host + ":" + strconv.Itoa(m.Port) + ")/" + m.DB + "?" + m.Config
	//返回连接字符串：用户名:密码@tcp(主机:端口)/数据库?高级配置
}