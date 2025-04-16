package config

type System struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
	Env string `yaml:"env"` //环境变量，dev，test，release
}