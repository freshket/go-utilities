package logger

import (
	"encoding/json"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type field struct {
	key string
	val interface{}
}

type AllowedType interface {
	string | *string | int | *int | bool | *bool
}

func param(key string, val interface{}) field {
	return field{"param." + key, val}
}

func Param[T AllowedType](key string, val T) field {
	return param(key, val)
}

func ParamJson(key string, val interface{}) field {
	b, _ := json.Marshal(val)

	return param(key, string(b))
}

type label field

func Label[T AllowedType](key string, val T) label {
	return label{"label." + key, val}
}

func convertFieldsToZapFields(fields []field) []zapcore.Field {
	var zfields []zapcore.Field

	for _, p := range fields {
		zfields = append(zfields, zap.Any(p.key, p.val))
	}

	return zfields
}
