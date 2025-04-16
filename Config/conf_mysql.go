package config

type Mysql struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	DB       string `yaml:"db"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Loglevel string `yaml:"log_level"` //日志级别，debug竞赛输出全部sql，dev，release就是线上环境
}