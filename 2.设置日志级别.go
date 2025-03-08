package main

import "go.uber.org/zap"

func main() {
	//logger, _ := zap.NewDevelopment()
	cfg := zap.NewDevelopmentConfig()
	cfg.Level = zap.NewAtomicLevelAt(zap.ErrorLevel)
	logger, _ := cfg.Build()

	logger.Debug("this is dev debug log")
	logger.Info("this is dev info log")
	logger.Warn("this is dev warn log")
	logger.Error("this is dev error log")
}
