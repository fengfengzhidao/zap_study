![](http://image.fengfengzhidao.com/fengfeng_110920250307225752.png?key=fengfengbuzhidao)

之前讲过logrus日志库的使用

但是后面logrus库官方停止更新了，然后看了下zap的stars非常高

![](http://image.fengfengzhidao.com/fengfeng_110920250308113155.png?key=fengfengbuzhidao)

并且zap的功能也很强大，配置项非常多

但是也正是因为配置项很多，导致做一些简单配置的时候，往往不知道怎么搞

所以本课程就把zap的常用功能讲解一下，帮助大家更好的使用此库



其实这些日志库，功能上都是大差不差的

一般学会以下功能即可

1. 时间格式化
2. 输出美化
3. 预设字段
4. 使用全局日志
5. 日志双写
6. 日志分片



安装

```Go
go get -u go.uber.org/zap
```



## 基本使用

`zap`库的使用与其他的日志库非常相似。

先创建一个`logger`，然后调用各个级别的方法记录日志（`Debug/Info/Error/Warn`）。

```Go
func dev() {
  logger, _ := zap.NewDevelopment()
  logger.Info("dev this is info")
  logger.Warn("dev this is warn")
  logger.Error("dev this is error")
}

func test() {
  logger := zap.NewExample()
  logger.Info("exam this is info")
  logger.Warn("exam this is warn")
  logger.Error("exam this is error")
}

func prod() {
  logger, _ := zap.NewProduction()
  logger.Info("prod this is info")
  logger.Warn("prod this is warn")
  logger.Error("prod this is error")
}
```



dev模式下，日志格式是text格式，并且warn和error会有栈信息

example模式下，格式是json，并且字段只有level和msg

prod模式下，格式也是json，多一个时间和函数位置字段，生产环境上用json格式的日志更方便排查



![](http://image.fengfengzhidao.com/fengfeng_110920250307233631.png?key=fengfengbuzhidao)



## 设置日志级别

```Go
// 使用 zap 的 NewDevelopmentConfig 快速配置
cfg := zap.NewDevelopmentConfig()
cfg.Level = zap.NewAtomicLevelAt(zap.WarnLevel)
// 创建 logger
logger, _ := cfg.Build()
logger.Debug("this is dev debug log")
logger.Info("this is dev info log")
logger.Warn("this is dev warn log")
logger.Error("this is dev error log")
logger.Fatal("this is dev fatal log")
```



## 时间格式化

默认的时间要么是带了时区的，要么就是时间戳，不太美观

如果只是想把时间变成我们喜欢的，可以使用如下配置

```Go
// 使用 zap 的 NewDevelopmentConfig 快速配置
cfg := zap.NewDevelopmentConfig()
cfg.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05") // 替换时间格式化方式

// 创建 logger
logger, _ := cfg.Build()

logger.Info("dev this is info")
logger.Warn("dev this is warn")
logger.Error("dev this is error")
```



![](http://image.fengfengzhidao.com/fengfeng_110920250307234105.png?key=fengfengbuzhidao)



## 标准日志和Sugar日志

标准的`*zap.Logger`只有这几个方法

如果想输出格式化字符串，还得在里面套函数，比较麻烦

![](http://image.fengfengzhidao.com/fengfeng_110920250308003730.png?key=fengfengbuzhidao)

所以可以用Sugar方法得到一个加强版实例

这里面方法就比较多了，常用的就是写格式化字符串哪些方法

![](http://image.fengfengzhidao.com/fengfeng_110920250308003911.png?key=fengfengbuzhidao)



## 结构化日志

`zap`支持通过`Field`的形式记录结构化日志，方便分析和查询

```Go
logger, _ := zap.NewDevelopment()
logger.Info("this is info",
  zap.String("username", "admin"),
  zap.Int("user_id", 42),
  zap.Bool("active", true),
)
```



![](http://image.fengfengzhidao.com/fengfeng_110920250308004251.png?key=fengfengbuzhidao)



## 输出美化

info，warn，error显示不同的颜色，看起来好看些

变色的关键  颜色控制字符

```Go
  fmt.Printf("\033[31mthis is 红色\n\033[0m")
  fmt.Printf("\033[32mthis is 绿色\n\033[0m")
  fmt.Printf("\033[33mthis is 黄色\n\033[0m")
  fmt.Printf("\033[34mthis is 蓝色\n\033[0m")
  fmt.Printf("\033[35mthis is 紫色\n\033[0m")
  fmt.Printf("\033[36mthis is 青色\n\033[0m")
```





```Go
// 定义颜色
const (
  colorRed    = "\033[31m"
  colorGreen  = "\033[32m"
  colorYellow = "\033[33m"
  colorBlue   = "\033[34m"
  colorReset  = "\033[0m"
)

// 自定义 EncodeLevel
func coloredLevelEncoder(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
  switch level {
  case zapcore.DebugLevel:
    enc.AppendString(colorBlue + "DEBUG" + colorReset)
  case zapcore.InfoLevel:
    enc.AppendString(colorGreen + "INFO" + colorReset)
  case zapcore.WarnLevel:
    enc.AppendString(colorYellow + "WARN" + colorReset)
  case zapcore.ErrorLevel, zapcore.DPanicLevel, zapcore.PanicLevel, zapcore.FatalLevel:
    enc.AppendString(colorRed + "ERROR" + colorReset)
  default:
    enc.AppendString(level.String()) // 默认行为
  }
}
func dev() {
  // 使用 zap 的 NewDevelopmentConfig 快速配置
  cfg := zap.NewDevelopmentConfig()
  cfg.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05") // 替换时间格式化方式
  cfg.EncoderConfig.EncodeLevel = coloredLevelEncoder
  // 创建 logger
  logger, _ := cfg.Build()

  logger.Info("dev this is info")
  logger.Warn("dev this is warn")
  logger.Error("dev this is error")
}
```



![](http://image.fengfengzhidao.com/fengfeng_110920250307234812.png?key=fengfengbuzhidao)



## 加日志前缀

一般如果有多个项目日志，可能会在日志的前面加上项目的名称

```Go
// 定义前缀
const logPrefix = "[MyApp] "

// 自定义 Encoder
type prefixedEncoder struct {
  zapcore.Encoder
}

func (e *prefixedEncoder) EncodeEntry(entry zapcore.Entry, fields []zapcore.Field) (*buffer.Buffer, error) {
  // 先调用原始的 EncodeEntry 方法生成日志行
  buf, err := e.Encoder.EncodeEntry(entry, fields)
  if err != nil {
    return nil, err
  }

  // 在日志行的最前面添加前缀
  logLine := buf.String()
  buf.Reset()
  buf.AppendString(logPrefix + logLine)

  return buf, nil
}
func dev() {
  // 使用 zap 的 NewDevelopmentConfig 快速配置
  cfg := zap.NewDevelopmentConfig()
  cfg.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05") // 替换时间格式化方式
  // 创建自定义的 Encoder
  encoder := &prefixedEncoder{
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

  logger.Info("dev this is info")
  logger.Warn("dev this is warn")
  logger.Error("dev this is error")
}
```



![](http://image.fengfengzhidao.com/fengfeng_110920250307235658.png?key=fengfengbuzhidao)



## 全局日志

因为zap推崇的还是以对象的形式使用日志

但是有些时候想要在应用程序的任何地方都可以直接使用的日志实例，那么可以用到全局日志

```Go
// 初始化全局日志
func initLogger() {
  // 使用 zap 的 NewDevelopmentConfig 快速配置
  cfg := zap.NewDevelopmentConfig()
  cfg.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05") // 替换时间格式化方式
  // 创建 Logger
  logger, _ := cfg.Build()
  zap.ReplaceGlobals(logger)
}

func dev() {
  zap.L().Info("dev this is info")
  zap.L().Warn("dev this is warn")
  zap.L().Error("dev this is error")
  zap.S().Infof("dev this is info %s", "xxx")
  zap.S().Warnf("dev this is warn %s", "xxx")
  zap.S().Errorf("dev this is error %s", "xxx")
}
```



L方法返回的是标准zap实例，S方法返回的是superZap的实例，superZap主要多了模板字符串方法



![](http://image.fengfengzhidao.com/fengfeng_110920250308000449.png?key=fengfengbuzhidao)



## 日志双写

常见的：控制台和日志文件双写

使用zapcore.NewTee可以组合多个core实例

```Go
// 初始化全局日志
func initLogger() {
  // 使用 zap 的 NewDevelopmentConfig 快速配置
  cfg := zap.NewDevelopmentConfig()
  cfg.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05") // 替换时间格式化方式
  // 创建 Core
  consoleCore := zapcore.NewCore(
    zapcore.NewConsoleEncoder(cfg.EncoderConfig),
    zapcore.AddSync(os.Stdout), // 输出到控制台
    zapcore.DebugLevel,         // 设置日志级别
  )
  file, _ := os.OpenFile("app.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
  fileCore := zapcore.NewCore(
    zapcore.NewConsoleEncoder(cfg.EncoderConfig),
    zapcore.AddSync(file), // 输出到文件
    zapcore.DebugLevel,    // 设置日志级别
  )
  core := zapcore.NewTee(consoleCore, fileCore)
  // 创建 Logger
  logger := zap.New(core, zap.AddCaller())
  zap.ReplaceGlobals(logger)
}

```



或者

使用 `zapcore.NewMultiWriteSyncer`

```Go
// 初始化全局日志
func initLogger() {
  // 使用 zap 的 NewDevelopmentConfig 快速配置
  cfg := zap.NewDevelopmentConfig()
  cfg.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05") // 替换时间格式化方式
  file, _ := os.OpenFile("app.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
  writeSyncer := zapcore.NewMultiWriteSyncer(
    zapcore.AddSync(os.Stdout),
    zapcore.AddSync(file),
  )
  // 创建 Core
  core := zapcore.NewCore(
    zapcore.NewConsoleEncoder(cfg.EncoderConfig),
    zapcore.AddSync(writeSyncer), // 输出到控制台
    zapcore.DebugLevel,         // 设置日志级别
  )
  // 创建 Logger
  logger := zap.New(core, zap.AddCaller())
  zap.ReplaceGlobals(logger)
}
```



## 日志切片

一般情况下，按照时间分片，每天的日志放到一个日志文件里面去

需要自定义Write方法

```Go
// 自定义日志写入器
type dynamicLogWriter struct {
  mu         sync.Mutex
  currentDay string
  file       *os.File
  logDir     string
}

func (w *dynamicLogWriter) Write(p []byte) (n int, err error) {
  w.mu.Lock()
  defer w.mu.Unlock()

  // 检查是否需要切换到新的日志文件
  currentDay := time.Now().Format("2006-01-02")
  if currentDay != w.currentDay {
    // 关闭当前日志文件
    if w.file != nil {
      w.file.Close()
    }

    // 创建新的日志文件
    if err := os.MkdirAll(w.logDir, 0755); err != nil {
      return 0, err
    }
    filePath := w.logDir + "/app-" + currentDay + ".log"
    file, err := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
    if err != nil {
      return 0, err
    }
    w.file = file
    w.currentDay = currentDay
  }

  // 写入日志
  return w.file.Write(p)
}

// 初始化全局日志
func initLogger() {
  // 使用 zap 的 NewDevelopmentConfig 快速配置
  cfg := zap.NewDevelopmentConfig()
  cfg.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05") // 替换时间格式化方式
  // 创建 Logger
  writer := &dynamicLogWriter{
    logDir: "logs",
  }
  // 创建 Core
  consoleCore := zapcore.NewCore(
    zapcore.NewConsoleEncoder(cfg.EncoderConfig),
    zapcore.AddSync(os.Stdout), // 输出到控制台
    zapcore.DebugLevel,         // 设置日志级别
  )
  fileCore := zapcore.NewCore(
    zapcore.NewConsoleEncoder(cfg.EncoderConfig),
    zapcore.AddSync(writer), // 输出到文件
    zapcore.DebugLevel,      // 设置日志级别
  )
  core := zapcore.NewTee(consoleCore, fileCore)
  // 创建 Logger
  logger := zap.New(core, zap.AddCaller())
  zap.ReplaceGlobals(logger)
}
```



## 日志按照level分片

把error的日志单独分出来，放到一个文件里面去

```Go
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

```



## 项目中的日志的使用

把以上功能做一个整合

1. 可设置级别
2. 时间格式化
3. 输出美化
4. 日志前缀
5. 日志双写
6. 日志时间分片、单独把error的日志分出来



```go
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


```





以上只是zap的基本使用，掌握之后再去探索zap的高级功能，就会更加的得心应手！





## 参考文档

go-zap日志框架 [https://blog.csdn.net/weixin_37375983/article/details/139625517](https://blog.csdn.net/weixin_37375983/article/details/139625517)

zap日志时间格式化 [https://blog.csdn.net/test1280/article/details/117266333](https://blog.csdn.net/test1280/article/details/117266333)

zap使用 [https://blog.csdn.net/qq_43466788/article/details/144973711](https://blog.csdn.net/qq_43466788/article/details/144973711)



