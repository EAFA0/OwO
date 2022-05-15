package log

import (
	"go.uber.org/zap"
)

type Filed = zap.Field

var Int = func(key string, val int) Filed { return zap.Int(key, val) }
var Err = func(err error) Filed { return String("errmsg", err.Error()) }
var String = func(key string, val string) Filed { return zap.String(key, val) }

var logger = zap.L()

func Debug(msg string, fileds ...Filed) {
	logger.Debug(msg, fileds...)
}

func Info(msg string, fileds ...Filed) {
	logger.Info(msg, fileds...)
}

func Warn(msg string, fileds ...Filed) {
	logger.Warn(msg, fileds...)
}

func Error(msg string, fileds ...Filed) {
	logger.Error(msg, fileds...)
}

func Panic(msg string, fileds ...Filed) {
	logger.Panic(msg, fileds...)
}
