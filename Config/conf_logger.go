package config

type Logger struct {
	Host         string `yaml:"host"`
	Prefix       string `yaml:"prefix"`
	Director     string `yaml:"director"`
	ShowLine     bool   `yaml:"show_line"` //是否显示行号
	LogInConsole bool   `yaml:"log_in_console"`//是否显示打印的路径
	Level 	  string `yaml:"level"` //日志级别，debug竞赛输出全部sql，dev，release就是线上环境
}