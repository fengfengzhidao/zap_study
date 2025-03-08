package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	//logger, _ := zap.NewDevelopment()
	cfg := zap.NewDevelopmentConfig()
	cfg.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02")
	logger, _ := cfg.Build()

	logger.Debug("this is dev debug log")
	logger.Info("this is dev info log")
	logger.Warn("this is dev warn log")
	logger.Error("this is dev error log")
}
