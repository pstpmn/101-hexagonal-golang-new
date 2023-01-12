package pkg

import (
	"os"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ILogger interface {
	Log(msg string)
	Error(msg string)
	ErrorPanic(msg string)
}

type logger struct {
	error *zap.Logger
	log   *zap.Logger
}

func (l logger) Log(msg string) {
	l.log.Info(msg)
}

func (l logger) Error(msg string) {
	l.error.Error(msg)
}

func (l logger) ErrorPanic(msg string) {
	l.error.Panic(msg)
}

func NewLogger() ILogger {

	// initialize the rotator
	logErrorFile := "../temp/logs/errors/%Y-%m-%d.log"
	logFile := "../temp/logs/out/%Y-%m-%d.log"

	rotatorErr, err := rotatelogs.New(
		logErrorFile,
		rotatelogs.WithMaxAge(60*24*time.Hour),
		rotatelogs.WithRotationTime(time.Hour))
	if err != nil {
		panic(err)
	}

	rotatorLog, err := rotatelogs.New(
		logFile,
		rotatelogs.WithMaxAge(60*24*time.Hour),
		rotatelogs.WithRotationTime(time.Hour))
	if err != nil {
		panic(err)
	}

	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	fileEncoder := zapcore.NewJSONEncoder(config)
	consoleEncoder := zapcore.NewConsoleEncoder(config)
	defaultLogLevel := zapcore.InfoLevel

	log := zapcore.AddSync(rotatorLog)
	error := zapcore.AddSync(rotatorErr)

	core := zapcore.NewTee(
		zapcore.NewCore(fileEncoder, log, defaultLogLevel),
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), defaultLogLevel),
	)

	errorCore := zapcore.NewTee(
		zapcore.NewCore(fileEncoder, error, defaultLogLevel),
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), defaultLogLevel),
	)

	zapLogger := zap.New(core)
	zapError := zap.New(errorCore)
	return logger{log: zapLogger, error: zapError}
}
