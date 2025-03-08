package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/buffer"
	"go.uber.org/zap/zapcore"
	"os"
)

// 自定义 Encoder
type levelEncoder struct {
	zapcore.Encoder
	errFile *os.File
}

func (e *levelEncoder) EncodeEntry(entry zapcore.Entry, fields []zapcore.Field) (*buffer.Buffer, error) {
	// 先调用原始的 EncodeEntry 方法生成日志行
	buf, err := e.Encoder.EncodeEntry(entry, fields)
	if err != nil {
		return nil, err
	}
	switch entry.Level {
	case zapcore.ErrorLevel:
		if e.errFile == nil {
			file, _ := os.OpenFile("err.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
			e.errFile = file
		}
		e.errFile.WriteString(buf.String())

	}

	return buf, nil
}
func main() {
	// 使用 zap 的 NewDevelopmentConfig 快速配置
	cfg := zap.NewDevelopmentConfig()
	cfg.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05") // 替换时间格式化方式
	// 创建自定义的 Encoder
	encoder := &levelEncoder{
		Encoder: zapcore.NewConsoleEncoder(cfg.EncoderConfig), // 使用 Console 编码器
	}
	// 创建 Core
	core := zapcore.NewCore(
		encoder,                    // 使用自定义的 Encoder
		zapcore.AddSync(os.Stdout), // 输出到控制台
		zapcore.DebugLevel,         // 设置日志级别
	)

	// 创建 Logger
	logger := zap.New(core, zap.AddCaller())

	logger.Info("this is info")
	logger.Warn("this is warn")
	logger.Error("this is error1")
	logger.Error("this is error2")
}
