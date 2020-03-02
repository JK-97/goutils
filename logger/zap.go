package logger

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type zapLogger struct {
	*zap.SugaredLogger
}

func getEncoder(isJSON bool) zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	if isJSON {
		return zapcore.NewJSONEncoder(encoderConfig)
	}
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getZapLevel(level string) zapcore.Level {
	switch level {
	case InfoLevel:
		return zapcore.InfoLevel
	case WarnLevel:
		return zapcore.WarnLevel
	case DebugLevel:
		return zapcore.DebugLevel
	case ErrorLevel:
		return zapcore.ErrorLevel
	case FatalLevel:
		return zapcore.FatalLevel
	default:
		return zapcore.InfoLevel
	}
}

func newZapLogger(config Configuration) (Logger, error) {
	var zapConfig zap.Config

	if config.IsProduction {
		zapConfig = zap.NewProductionConfig()
	} else {
		zapConfig = zap.NewDevelopmentConfig()
	}
	if config.Level != "" {
		zapConfig.Level = zap.NewAtomicLevelAt(getZapLevel(config.Level))
	}

	logger, err := zapConfig.Build()
	if err != nil {
		return nil, err
	}
	return &zapLogger{
		SugaredLogger: logger.Sugar(),
	}, err
}

func (l *zapLogger) WithFields(fields Fields) Logger {
	var f = make([]interface{}, 0)
	for k, v := range fields {
		f = append(f, k)
		f = append(f, v)
	}
	newLogger := l.SugaredLogger.With(f...)
	return &zapLogger{newLogger}
}

// 常用颜色
const (
	TextBlack = iota + 30
	TextRed
	TextGreen
	TextYellow
	TextBlue
	TextMagenta
	TextCyan
	TextWhite
)

// SetColor 设置颜色
func SetColor(msg string, conf, bg, text int) string {
	return fmt.Sprintf("%c[%d;%d;%dm%s%c[0m", 0x1B, conf, bg, text, msg, 0x1B)
}
