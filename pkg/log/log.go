package log

import (
	"go.uber.org/zap"
)

type Filed = zap.Field

var Int = func(key string, val int) Filed { return zap.Int(key, val) }
var Err = func(err error) Filed { return String("msg", err.Error()) }
var String = func(key string, val string) Filed { return zap.String(key, val) }

var logger = zap.L()

func Debug(msg string, fields ...Filed) {
	logger.Debug(msg, fields...)
}

func Info(msg string, fields ...Filed) {
	logger.Info(msg, fields...)
}

func Warn(msg string, fields ...Filed) {
	logger.Warn(msg, fields...)
}

func Error(msg string, fields ...Filed) {
	logger.Error(msg, fields...)
}

func Panic(msg string, fields ...Filed) {
	logger.Panic(msg, fields...)
}
