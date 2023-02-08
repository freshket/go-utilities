package logger

import (
	"context"
)

var correlationIDCtxKey = &struct{ bool }{}

func CtxWithCorrelationID(ctx context.Context, correlationID string) context.Context {
	if len(correlationID) == 0 {
		return ctx
	}

	return context.WithValue(ctx, correlationIDCtxKey, correlationID)
}

func correlationIDFieldFromCtx(ctx context.Context) field {
	correlationID, _ := ctx.Value(correlationIDCtxKey).(string)

	return field{"correlation_id", correlationID}
}
