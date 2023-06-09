package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewDevInfoZapLogger() (*zap.Logger, error) {
	a := zap.NewDevelopmentConfig()
	return a.Build()
}
func NewDevErrorZapLogger() (*zap.Logger, error) {
	a := zap.NewDevelopmentConfig()
	a.Level.SetLevel(zapcore.ErrorLevel)
	return a.Build()
}
