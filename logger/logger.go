package logger

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
	"strings"
)

type ConfigOptions struct {
	Level      string
	Filename   string
	MaxSize    int
	MaxBackups int
	MaxAge     int
}

func NewLogger(options *ConfigOptions) *logrus.Logger {

	logFile := &lumberjack.Logger{
		Filename:   options.Filename,   // 日志文件的路径和文件名
		MaxSize:    options.MaxSize,    // 每个日志文件的最大大小（以 MB 为单位）
		MaxBackups: options.MaxBackups, // 保留的旧日志文件的最大数目
		MaxAge:     options.MaxAge,     // 保留的旧日志文件的最大天数
		LocalTime:  true,               // 使用本地时间
		Compress:   true,               // 是否压缩旧日志文件
	}

	l := logrus.New()
	l.SetOutput(io.MultiWriter(logFile, os.Stdout))
	l.SetLevel(getLevel(options.Level))

	return l
}

func getLevel(levelString string) logrus.Level {
	var level logrus.Level
	upLevelString := strings.ToUpper(levelString)
	switch upLevelString {
	case "TRACE":
		level = logrus.TraceLevel
		break
	case "DEBUG":
		level = logrus.DebugLevel
		break
	case "INFO":
		level = logrus.InfoLevel
		break
	case "WARN":
		level = logrus.WarnLevel
		break
	case "ERROR":
		level = logrus.ErrorLevel
		break
	case "FATAL":
		level = logrus.FatalLevel
		break
	case "PANIC":
		level = logrus.PanicLevel
		break
	default:
		level = logrus.ErrorLevel
	}

	return level
}
