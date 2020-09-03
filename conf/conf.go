package conf

import (
	"errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"strconv"
)

type Config struct {
	Id              uint32            `yaml:"id"`
	ClientPort      uint32            `yaml:"client_port"`
	LaunchSize      uint32            `yaml:"launch_size"`
	LogLevel        string            `yaml:"log_level"`
	Members         map[uint32]string `yaml:"members"`
}

var Conf *Config

func init() {
	Conf = &Config{
		Id: 1,
		LaunchSize: 3,
		LogLevel:   "INFO",
	}
	ymlFile, err := ioutil.ReadFile("conf.yml")
	if err == nil {
		_ = yaml.Unmarshal(ymlFile, Conf)
	}
	id, err := strconv.Atoi(os.Getenv("NARUTO_ID"))
	if err == nil {
		Conf.Id = uint32(id)
	}
	if Conf.Id == 0 {
		panic(errors.New("id not set or set to zero"))
	}
	clientPort, err := strconv.Atoi(os.Getenv("NARUTO_CLIENT_PORT"))
	if err == nil {
		Conf.ClientPort = uint32(clientPort)
	}
	launchSize, err := strconv.Atoi(os.Getenv("NARUTO_LAUNCH_SIZE"))
	if err == nil {
		Conf.LaunchSize = uint32(launchSize)
	}
	logLevel := os.Getenv("NARUTO_LOG_LEVEL")
	if logLevel != "" {
		Conf.LogLevel = logLevel
	}
}
