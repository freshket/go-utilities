package lxg

import (
	"context"
)

var labelsCtxKey = &struct{ bool }{}

func CtxWithLabels(ctx context.Context, labels ...label) context.Context {
	if len(labels) == 0 {
		return ctx
	}

	fields, _ := ctx.Value(labelsCtxKey).([]field)

	for _, l := range labels {
		fields = append(fields, field(l))
	}

	return context.WithValue(ctx, labelsCtxKey, fields)
}

func labelsFromCtx(ctx context.Context) []field {
	fields, _ := ctx.Value(labelsCtxKey).([]field)

	return fields
}
