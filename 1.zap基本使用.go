package main

import (
	"go.uber.org/zap"
)

func dev() {
	logger, _ := zap.NewDevelopment()
	logger.Debug("this is dev debug log")
	logger.Info("this is dev info log")
	logger.Warn("this is dev warn log")
	logger.Error("this is dev error log")
	logger.Panic("this is dev panic log")
	logger.Fatal("this is dev fatal log")
}

func example() {
	logger := zap.NewExample()
	logger.Debug("this is example debug log")
	logger.Info("this is example info log")
	logger.Warn("this is example warn log")
	logger.Error("this is example error log")
	logger.Panic("this is example panic log")
	logger.Fatal("this is example fatal log")
}

func prod() {
	logger, _ := zap.NewProduction()
	logger.Debug("this is prod debug log")
	logger.Info("this is prod info log")
	logger.Warn("this is prod warn log")
	logger.Error("this is prod error log")
	logger.Panic("this is prod panic log")
	logger.Fatal("this is prod fatal log")
}

func main() {
	//dev()
	//example()
	prod()
}
