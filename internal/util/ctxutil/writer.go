package ctxutil

import (
	"context"

	"github.com/nmluci/stellar-payment-lite/internal/commonkey"
	"github.com/nmluci/stellar-payment-lite/internal/indto"
)

func WrapCtx(ctx context.Context, key commonkey.CtxKey, val interface{}) context.Context {
	return context.WithValue(ctx, key, val)
}

func GetCtx[T any](ctx context.Context, key commonkey.CtxKey) (res T, ok bool) {
	res, ok = ctx.Value(key).(T)
	return
}

func GetUserCTX(ctx context.Context) (res *indto.UserRole) {
	res, ok := GetCtx[*indto.UserRole](ctx, commonkey.SCOPE_CTX_KEY)
	if !ok {
		panic("invalid data")
	}

	return
}
