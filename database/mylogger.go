package database

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
	"time"
)

type MyLogger struct {
	logrus.Entry
	logger.Config
	logLevel logrus.Level
}

func NewMyLogger(logLevel logrus.Level, l *logrus.Logger) *MyLogger {
	return &MyLogger{
		Entry:    *l.WithField("module", "gorm"),
		logLevel: logLevel,
		Config: logger.Config{
			SlowThreshold: time.Second, // 慢查询阈值
			LogLevel:      getLogLevel(logLevel),
			Colorful:      true,
		},
	}
}

func getLogLevel(l logrus.Level) logger.LogLevel {
	switch l {
	case logrus.DebugLevel, logrus.TraceLevel, logrus.InfoLevel:
		return logger.Info
	case logrus.WarnLevel:
		return logger.Warn
	case logrus.ErrorLevel:
		return logger.Error
	case logrus.FatalLevel:
		return logger.Silent
	default:
		return logger.Silent
	}
}

func (l *MyLogger) LogMode(level logger.LogLevel) logger.Interface {
	newLogger := *l
	newLogger.LogLevel = level
	return &newLogger
}

func (l *MyLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	switch l.logLevel {
	case logrus.TraceLevel:
		l.Tracef(msg, data...)
	case logrus.DebugLevel:
		l.Debugf(msg, data...)
	default:
		l.Infof(msg, data...)
	}
}

func (l *MyLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	l.Warnf(msg, data...)
}

func (l *MyLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	l.Errorf(msg, data...)
}

func (l *MyLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if l.LogLevel <= 0 {
		return
	}

	elapsed := time.Since(begin)
	switch {
	case err != nil && l.LogLevel >= logger.Error:
		sql, rows := fc()
		if rows == -1 {
			l.WithError(err).Errorf("%s [%.3fms] [rows:%v] %s", utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, "-", sql)
		} else {
			l.WithError(err).Errorf("%s [%.3fms] [rows:%v] %s", utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	case elapsed > l.SlowThreshold && l.SlowThreshold != 0 && l.LogLevel >= logger.Warn:
		sql, rows := fc()
		slowLog := fmt.Sprintf("SLOW SQL >= %v", l.SlowThreshold)
		if rows == -1 {
			l.Warnf("%s [%.3fms] [rows:%v] %s %s", utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, "-", slowLog, sql)
		} else {
			l.Warnf("%s [%.3fms] [rows:%v] %s %s", utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, rows, slowLog, sql)
		}
	case l.LogLevel >= logger.Info:
		sql, rows := fc()
		if rows == -1 {
			l.Infof("%s [%.3fms] [rows:%v] %s", utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, "-", sql)
		} else {
			l.Infof("%s [%.3fms] [rows:%v] %s", utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	}
}
