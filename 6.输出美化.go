package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	Bule   = "\033[34m"
	Yellow = "\033[33m"
	Red    = "\033[31m"
	Reset  = "\033[0m"
)

func encodeLevel(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	switch level {
	case zapcore.InfoLevel:
		enc.AppendString(Bule + "INFO" + Reset)
	case zapcore.WarnLevel:
		enc.AppendString(Yellow + "WARN" + Reset)
	case zapcore.ErrorLevel, zapcore.DPanicLevel, zapcore.PanicLevel, zapcore.FatalLevel:
		enc.AppendString(Red + "ERROR" + Reset)
	default:
		enc.AppendString(level.String())
	}
}

func main() {
	cfg := zap.NewDevelopmentConfig()
	cfg.EncoderConfig.EncodeLevel = encodeLevel
	logger, _ := cfg.Build()
	logger.Debug("this is dev debug log")
	logger.Info("this is dev info log")
	logger.Warn("this is dev warn log")
	logger.Error("this is dev error log")

	//fmt.Printf("\033[31mthis is 红色\n\033[0m")
	//fmt.Printf("\033[32mthis is 绿色\n\033[0m")
	//fmt.Printf("\033[33mthis is 黄色\n\033[0m")
	//fmt.Printf("\033[34mthis is 蓝色\n\033[0m")
	//fmt.Printf("\033[35mthis is 紫色\n\033[0m")
	//fmt.Printf("\033[36mthis is 青色\n\033[0m")
}
