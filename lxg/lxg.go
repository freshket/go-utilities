package lxg

import (
	"context"
	"errors"
	"os"

	"go.elastic.co/ecszap"
	"go.uber.org/zap"
)

type logger struct {
	disabled bool
	zl       *zap.Logger
}

var loggerinstance *logger

func getlogger() *logger {
	if loggerinstance != nil {
		return loggerinstance
	}

	logname := os.Getenv("LOG_NAME")
	logdisabled := os.Getenv("LOG_DISABLED") == "true"
	loglvenabled := zap.DebugLevel

	switch os.Getenv("LOG_LEVEL") {
	case "INFO":
		loglvenabled = zap.InfoLevel
	case "WARN":
		loglvenabled = zap.WarnLevel
	case "ERROR":
		loglvenabled = zap.ErrorLevel
	case "PANIC":
		loglvenabled = zap.PanicLevel
	case "FATAL":
		loglvenabled = zap.FatalLevel
	}

	zl := zap.New(
		ecszap.NewCore(
			ecszap.NewDefaultEncoderConfig(),
			os.Stdout,
			loglvenabled,
		),
	).Named(logname)

	return &logger{
		disabled: logdisabled,
		zl:       zl,
	}
}

func fieldsFromCtx(ctx context.Context) []field {
	var fields []field

	fields = append(fields, labelsFromCtx(ctx)...)
	fields = append(fields, correlationIDFieldFromCtx(ctx))

	return fields
}

func Debug(msg string, fields ...field) {
	getlogger().zl.Debug(msg, convertFieldsToZapFields(fields)...)
}

func DebugCtx(ctx context.Context, msg string, fields ...field) {
	Debug(msg, append(fieldsFromCtx(ctx), fields...)...)
}

func Info(msg string, fields ...field) {
	getlogger().zl.Info(msg, convertFieldsToZapFields(fields)...)
}

func InfoCtx(ctx context.Context, msg string, fields ...field) {
	Info(msg, append(fieldsFromCtx(ctx), fields...)...)
}

func Warn(msg string, fields ...field) {
	getlogger().zl.Warn(msg, convertFieldsToZapFields(fields)...)
}

func WarnCtx(ctx context.Context, msg string, fields ...field) {
	Warn(msg, append(fieldsFromCtx(ctx), fields...)...)
}

func Error(msg string, err error, fields ...field) {
	zfields := append(
		convertFieldsToZapFields(fields),
		zap.Error(err),
	)

	getlogger().zl.Error(msg, zfields...)
}

func ErrorCtx(ctx context.Context, msg string, err error, fields ...field) {
	Error(msg, err, append(fieldsFromCtx(ctx), fields...)...)
}

func Error2(msg string, fields ...field) {
	zfields := append(
		convertFieldsToZapFields(fields),
		zap.Error(errors.New(msg)),
	)

	getlogger().zl.Error(msg, zfields...)
}

func Error2Ctx(ctx context.Context, msg string, fields ...field) {
	Error2(msg, append(fieldsFromCtx(ctx), fields...)...)
}

func Close() {
	_ = getlogger().zl.Sync()
}
