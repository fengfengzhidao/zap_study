package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/buffer"
	"go.uber.org/zap/zapcore"
	"os"
)

type myEncoder struct {
	zapcore.Encoder
}

func (m myEncoder) EncodeEntry(entry zapcore.Entry, fields []zapcore.Field) (*buffer.Buffer, error) {
	buff, err := m.Encoder.EncodeEntry(entry, fields)
	if err != nil {
		return buff, err
	}

	data := buff.String()
	buff.Reset()
	buff.AppendString("[myApp] " + data)

	return buff, nil
}

func main() {

	core := zapcore.NewCore(
		myEncoder{
			Encoder: zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig()),
		},
		os.Stdout,
		zapcore.InfoLevel,
	)

	logger := zap.New(core, zap.AddCaller())
	logger.Info("this is info log")
}
