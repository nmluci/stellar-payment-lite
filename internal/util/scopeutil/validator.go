package scopeutil

import (
	"context"

	"github.com/nmluci/stellar-payment-lite/internal/commonkey"
	"github.com/nmluci/stellar-payment-lite/internal/indto"
	"github.com/nmluci/stellar-payment-lite/internal/util/ctxutil"
)

func ValidateScope(ctx context.Context, roles ...int64) (ok bool) {
	userMeta, ok := ctxutil.GetCtx[*indto.UserRole](ctx, commonkey.SCOPE_CTX_KEY)
	if !ok {
		return false
	}

	for _, v := range roles {
		if userMeta.RoleID == v {
			return true
		}
	}

	return false
}
