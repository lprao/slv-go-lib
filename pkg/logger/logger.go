package logger

import (
	"context"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger struct
type Log struct {
	sugarLogger *zap.SugaredLogger
}

// Logger interface methods
type Logger interface {
	InitLogger()
	NewContext(context.Context) context.Context
	Infof(template string, args ...interface{})
	Debugf(template string, args ...interface{})
	Warnf(template string, args ...interface{})
	Errorf(template string, args ...interface{})
	Fatalf(template string, args ...interface{})
}

// NewLogger
func NewLogger() *Log {
	return &Log{}
}

// Intialize the sugar logger
func (l *Log) InitLogger() {
	zapEncoderConfig := zap.NewDevelopmentEncoderConfig()
	zapEncoderConfig.LevelKey = "Severity"
	zapEncoderConfig.TimeKey = "Timestamp"
	zapEncoderConfig.MessageKey = "Message"
	zapEncoderConfig.CallerKey = "Method"
	zapEncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	zapEncoder := zapcore.NewJSONEncoder(zapEncoderConfig)
	core := zapcore.NewCore(zapEncoder, zapcore.AddSync(os.Stdout), zap.NewAtomicLevelAt(zapcore.InfoLevel))

	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	l.sugarLogger = logger.Sugar()
	l.sugarLogger.Sync()
}

// Infof logger
func (l *Log) Infof(template string, args ...interface{}) {
	l.sugarLogger.Infof(template, args...)
}

// Debugf logger
func (l *Log) Debugf(template string, args ...interface{}) {
	l.sugarLogger.Debugf(template, args...)
}

// Warnf logger
func (l *Log) Warnf(template string, args ...interface{}) {
	l.sugarLogger.Warnf(template, args...)
}

// Errorf logger
func (l *Log) Errorf(template string, args ...interface{}) {
	l.sugarLogger.Errorf(template, args...)
}

// Fatalf logger
func (l *Log) Fatalf(template string, args ...interface{}) {
	l.sugarLogger.Fatalf(template, args...)
}
