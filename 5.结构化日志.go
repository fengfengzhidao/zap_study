package main

import "go.uber.org/zap"

func main() {
	logger, _ := zap.NewDevelopment()
	logger.Info("this is log")
	logger.Info("this is log",
		zap.String("name", "fengfeng"),
		zap.Int("age", 12),
		zap.Bool("isGender", true),
	)
}
