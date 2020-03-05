// reference: https://www.mountedthoughts.com/golang-logger-interface/

package logger

import "time"

// A global variable so that log functions can be directly accessed
var log Logger

// Fields Type to pass when we want to call WithFields for structured logging
type Fields = map[string]interface{}

const (
	//DebugLevel has verbose message
	DebugLevel = "debug"
	//InfoLevel is default log level
	InfoLevel = "info"
	//WarnLevel is for logging messages about possible issues
	WarnLevel = "warn"
	//ErrorLevel is for logging errors
	ErrorLevel = "error"
	//FatalLevel is for logging fatal messages. The sytem shutsdown after logging the message.
	FatalLevel = "fatal"
)

// LogFunc logs a message.
type LogFunc func(args ...interface{})

// LogFormatFunc uses fmt.Sprintf to log a templated message.
type LogFormatFunc func(format string, args ...interface{})

// Logger is our contract for the logger
type Logger interface {
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
	Panicf(format string, args ...interface{})

	Debug(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})
	Panic(args ...interface{})

	WithFields(keyValues Fields) Logger
}

// Configuration stores the config for the logger
// For some loggers there can only be one level across writers, for such the level of Console is picked by default
type Configuration struct {
	IsProduction bool
	Level        string
	CallerSkip   int
	Fields       Fields
}

func init() {
	err := SetLoggerConfig(Configuration{})
	if err != nil {
		panic(err)
	}
}

// SetLoggerConfig sets logger configuration
func SetLoggerConfig(config Configuration) error {
	logger, err := newZapLogger(config)
	if err != nil {
		return err
	}
	log = logger
	SetLogger(logger)
	return nil
}

// GetLogger 获取当前使用的 Logger
func GetLogger() Logger {
	return log
}

// Package funcs
var (
	Debugf LogFormatFunc
	Infof  LogFormatFunc
	Warnf  LogFormatFunc
	Errorf LogFormatFunc
	Fatalf LogFormatFunc
	Panicf LogFormatFunc

	Debug LogFunc
	Info  LogFunc
	Warn  LogFunc
	Error LogFunc
	Fatal LogFunc
	Panic LogFunc

	WithFields func(keyValues Fields) Logger
)

// SetLogger 设置 Logger
func SetLogger(l Logger) {
	Debugf = l.Debugf
	Infof = l.Infof
	Warnf = l.Warnf
	Errorf = l.Errorf
	Fatalf = l.Fatalf
	Panicf = l.Panicf
	Debug = l.Debug
	Info = l.Info
	Warn = l.Warn
	Error = l.Error
	Fatal = l.Fatal
	Panic = l.Panic
	WithFields = l.WithFields
}

// HTTPTimestampFormat 打印 HTTP 日志使用的时间格式
var HTTPTimestampFormat = "2006/01/02 - 15:04:05"

// LogHTTP 记录 HTTP 请求
func LogHTTP(timestamp time.Time, statusCode int, clientIP, method, path string, latency time.Duration, errorMessage string) {

	if latency > time.Minute {
		// Truncate in a golang < 1.8 safe way
		latency = latency - latency%time.Second
	}

	log.Infof("%v - %3d - %13v - %15s - %-7s %s - %s", timestamp.Format(HTTPTimestampFormat), statusCode, latency, clientIP, method, path, errorMessage)
}
