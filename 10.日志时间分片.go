package main

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"sync"
	"time"
)

type myWriter struct {
	file        *os.File
	mutex       sync.Mutex
	currentDate string
}

func (m *myWriter) Write(b []byte) (n int, err error) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	now := time.Now().Format("2006-01-02")
	if m.currentDate == now {
		return m.file.Write(b)
	}
	if m.file != nil {
		m.file.Close()
	}
	name := fmt.Sprintf("logs/%s.log", now)
	file, _ := os.OpenFile(name, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	m.file = file
	m.currentDate = now
	return file.Write(b)
}

func main() {
	cfg := zap.NewDevelopmentConfig()
	cfg.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02")

	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(cfg.EncoderConfig),
		zapcore.NewMultiWriteSyncer(os.Stdout, zapcore.AddSync(&myWriter{})),
		zapcore.InfoLevel,
	)

	logger := zap.New(core, zap.AddCaller())
	for i := 0; i < 10; i++ {
		logger.Sugar().Infof("this is %d log", i)
		time.Sleep(time.Second)
	}
}
