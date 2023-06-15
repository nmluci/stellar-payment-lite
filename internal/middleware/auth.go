package middleware

import (
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/nmluci/stellar-payment-lite/internal/commonkey"
	"github.com/nmluci/stellar-payment-lite/internal/service"
	"github.com/nmluci/stellar-payment-lite/internal/util/ctxutil"
	"github.com/nmluci/stellar-payment-lite/internal/util/echttputil"
	"github.com/nmluci/stellar-payment-lite/pkg/errs"
)

func AuthorizationMiddleware(svc service.Service) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			var err error
			header := c.Request().Header
			ctx := c.Request().Context()

			token := header.Get("authorization")
			if token == "" {
				return echttputil.WriteErrorResponse(c, errs.ErrNoAccess)
			}

			splittedToken := strings.Split(token, " ")
			if len(splittedToken) != 2 {
				return echttputil.WriteErrorResponse(c, errs.ErrInvalidCred)
			}

			if splittedToken[0] != "Bearer" {
				return echttputil.WriteErrorResponse(c, errs.ErrInvalidCred)
			}

			accessToken := splittedToken[1]
			user, err := svc.FindUserByAccessToken(ctx, accessToken)
			if err != nil {
				return echttputil.WriteErrorResponse(c, errs.ErrInvalidCred)
			} else if user == nil {
				return echttputil.WriteErrorResponse(c, errs.ErrNoAccess)
			}

			c.SetRequest(c.Request().Clone(ctxutil.WrapCtx(ctx, commonkey.SCOPE_CTX_KEY, user)))
			return next(c)
		}
	}
}
