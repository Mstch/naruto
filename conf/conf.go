package conf

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"strconv"
)

type Config struct {
	Port       int32  `yaml:"port"`
	UPort      int32  `yaml:"uport"`
	LaunchSize int32  `yaml:"launch_size"`
	LogLevel   string `yaml:"log_level"`
}

var Conf *Config

func init() {
	//TODO 文件配置
	Conf = &Config{
		Port:       1000,
		UPort:      8848,
		LaunchSize: 3,
		LogLevel:   "INFO",
	}
	ymlFile, err := ioutil.ReadFile("conf.yml")
	if err == nil {
		_ = yaml.Unmarshal(ymlFile, Conf)
	}
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err == nil {
		Conf.Port = int32(port)
	}
	launchSize, err := strconv.Atoi(os.Getenv("LAUNCH_SIZE"))
	if err == nil {
		Conf.LaunchSize = int32(launchSize)
	}
	logLevel := os.Getenv("LOG_LEVEL")
	if logLevel != "" {
		Conf.LogLevel = logLevel
	}
}
