// reference: https://www.mountedthoughts.com/golang-logger-interface/

package logger

// A global variable so that log functions can be directly accessed
var log Logger

//Fields Type to pass when we want to call WithFields for structured logging
type Fields map[string]interface{}

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

//Logger is our contract for the logger
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
}

func init() {
	err := SetLoggerConfig(Configuration{})
	if err != nil {
		panic(err)
	}
}

//SetLoggerConfig sets logger configuration
func SetLoggerConfig(config Configuration) error {
	logger, err := newZapLogger(config)
	if err != nil {
		return err
	}
	log = logger
	return nil
}

func Debugf(format string, args ...interface{}) {
	log.Debugf(format, args...)
}

func Infof(format string, args ...interface{}) {
	log.Infof(format, args...)
}

func Warnf(format string, args ...interface{}) {
	log.Warnf(format, args...)
}

func Errorf(format string, args ...interface{}) {
	log.Errorf(format, args...)
}

func Fatalf(format string, args ...interface{}) {
	log.Fatalf(format, args...)
}

func Panicf(format string, args ...interface{}) {
	log.Panicf(format, args...)
}

func Debug(args ...interface{}) {
	log.Debug(args...)
}

func Info(args ...interface{}) {
	log.Info(args...)
}

func Warn(args ...interface{}) {
	log.Warn(args...)
}

func Error(args ...interface{}) {
	log.Error(args...)
}

func Fatal(args ...interface{}) {
	log.Fatal(args...)
}

func Panic(args ...interface{}) {
	log.Panic(args...)
}

func WithFields(keyValues Fields) Logger {
	return log.WithFields(keyValues)
}
