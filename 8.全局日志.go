package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func InitLogger() {
	cfg := zap.NewDevelopmentConfig()
	cfg.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02")
	logger, _ := cfg.Build()

	zap.ReplaceGlobals(logger)
}

func main() {
	InitLogger()
	zap.L().Info("全局日志方法")
	zap.S().Infof("全局日志方法")
	xxx()
}

func xxx() {
	zap.L().Info("xxx")
}
