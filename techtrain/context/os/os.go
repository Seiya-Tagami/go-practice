package os

import (
	"context"
	"errors"
)

type contextKey string

var osKey contextKey

func ContextWithOS(ctx context.Context, os string) context.Context {
	return context.WithValue(ctx, osKey, os)
}

func OSFromContext(ctx context.Context) (string, error) {
	os, ok := ctx.Value(osKey).(string)
	if !ok {
		return "", errors.New("os not found")
	}
	return os, nil
}
