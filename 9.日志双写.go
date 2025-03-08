package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func main() {
	f2()
}

func f2() {
	cfg := zap.NewDevelopmentConfig()
	cfg.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02")

	file, _ := os.OpenFile("app.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)

	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(cfg.EncoderConfig),
		zapcore.NewMultiWriteSyncer(os.Stdout, file),
		zapcore.InfoLevel,
	)

	logger := zap.New(core, zap.AddCaller())
	logger.Info("xxx")
}

func f1() {
	cfg := zap.NewDevelopmentConfig()
	cfg.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02")

	consoleCore := zapcore.NewCore(
		zapcore.NewConsoleEncoder(cfg.EncoderConfig),
		os.Stdout,
		zapcore.InfoLevel,
	)
	file, _ := os.OpenFile("app.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	fileCore := zapcore.NewCore(
		zapcore.NewConsoleEncoder(cfg.EncoderConfig),
		file,
		zapcore.InfoLevel,
	)

	core := zapcore.NewTee(consoleCore, fileCore)

	logger := zap.New(core, zap.AddCaller())
	logger.Info("xxx")
}
