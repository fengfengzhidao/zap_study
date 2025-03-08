package main

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/buffer"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

// logEncoder 时间分片和level分片同时做
type logEncoder struct {
	zapcore.Encoder
	errFile     *os.File
	file        *os.File
	currentDate string
}

const (
	BuleColor   = "\033[34m"
	YellowColor = "\033[33m"
	RedColor    = "\033[31m"
	ResetColor  = "\033[0m"
)

func myEncodeLevel(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	switch level {
	case zapcore.InfoLevel:
		enc.AppendString(BuleColor + "INFO" + ResetColor)
	case zapcore.WarnLevel:
		enc.AppendString(YellowColor + "WARN" + ResetColor)
	case zapcore.ErrorLevel, zapcore.DPanicLevel, zapcore.PanicLevel, zapcore.FatalLevel:
		enc.AppendString(RedColor + "ERROR" + ResetColor)
	default:
		enc.AppendString(level.String())
	}
}
func (e *logEncoder) EncodeEntry(entry zapcore.Entry, fields []zapcore.Field) (*buffer.Buffer, error) {
	// 先调用原始的 EncodeEntry 方法生成日志行
	buff, err := e.Encoder.EncodeEntry(entry, fields)
	if err != nil {
		return nil, err
	}
	data := buff.String()
	buff.Reset()
	buff.AppendString("[myApp] " + data)
	data = buff.String()
	// 时间分片
	now := time.Now().Format("2006-01-02")
	if e.currentDate != now {
		os.MkdirAll(fmt.Sprintf("logs/%s", now), 0666)
		// 时间不同，先创建目录
		name := fmt.Sprintf("logs/%s/out.log", now)
		file, _ := os.OpenFile(name, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
		e.file = file
		e.currentDate = now
	}

	switch entry.Level {
	case zapcore.ErrorLevel:
		if e.errFile == nil {
			name := fmt.Sprintf("logs/%s/err.log", now)
			file, _ := os.OpenFile(name, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
			e.errFile = file
		}
		e.errFile.WriteString(buff.String())
	}

	if e.currentDate == now {
		e.file.WriteString(data)
	}
	return buff, nil
}

func InitLog() *zap.Logger {
	// 使用 zap 的 NewDevelopmentConfig 快速配置
	cfg := zap.NewDevelopmentConfig()
	cfg.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05") // 替换时间格式化方式
	cfg.EncoderConfig.EncodeLevel = myEncodeLevel
	// 创建自定义的 Encoder
	encoder := &logEncoder{
		Encoder: zapcore.NewConsoleEncoder(cfg.EncoderConfig), // 使用 Console 编码器
	}
	// 创建 Core
	core := zapcore.NewCore(
		encoder,                    // 使用自定义的 Encoder
		zapcore.AddSync(os.Stdout), // 输出到控制台
		zapcore.InfoLevel,          // 设置日志级别
	)
	// 创建 Logger
	logger := zap.New(core, zap.AddCaller())

	zap.ReplaceGlobals(logger)
	return logger
}

func main() {

	logger := InitLog()
	logger.Info("this is info")
	logger.Warn("this is warn")
	logger.Error("this is error1")
	logger.Error("this is error2")
	zap.S().Infof("%s xxx", "fengfeng")
}
