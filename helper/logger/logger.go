package logger

import (
	"github.com/Mstch/naruto/conf"
	"log"
	"os"
)

type LEVEL int

const (
	DEBUG = iota
	INFO
	WARNING
	ERROR
)

var (
	Level   = INFO
	debug   = &log.Logger{}
	info    = &log.Logger{}
	warning = &log.Logger{}
	err     = &log.Logger{}
)

func init() {
	c := conf.Conf

	switch c.LogLevel {
	case "DEBUG":
		Level = DEBUG
	case "INFO":
		Level = INFO
	case "WARNING":
		Level = WARNING
	case "ERROR":
		Level = ERROR
	default:
		log.Println("日志等级未找到,将配置为INFO")
		Level = INFO
	}

	debug.SetFlags(log.LstdFlags)
	info.SetFlags(log.LstdFlags)
	warning.SetFlags(log.LstdFlags)
	err.SetFlags(log.LstdFlags)

	debug.SetOutput(os.Stdout)
	info.SetOutput(os.Stdout)
	warning.SetOutput(os.Stderr)
	err.SetOutput(os.Stderr)

	debug.SetPrefix("[DEBUG]")
	info.SetPrefix("[INFO]")
	warning.SetPrefix("[WARNING]")
	err.SetPrefix("[ERROR]")

	Debug("初始化logger成功")
}

func Debug(format string, args ...interface{}) {
	if Level == DEBUG {
		debug.Printf(format, args...)
	}
}

func Info(format string, args ...interface{}) {
	if Level <= INFO {
		info.Printf(format, args...)
	}
}

func Warning(format string, args ...interface{}) {
	if Level <= WARNING {
		warning.Printf(format, args...)
	}
}

func Error(format string, args ...interface{}) {
	if Level <= ERROR {
		err.Printf(format, args...)
	}
}

func Fatal(format string, e error, args ...interface{}) {
	if Level <= ERROR {
		err.Printf(format, args...)
	}
	panic(e)
}
