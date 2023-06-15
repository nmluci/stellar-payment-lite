package handler

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/nmluci/stellar-payment-lite/internal/util/echttputil"
	"github.com/nmluci/stellar-payment-lite/pkg/dto"
	"github.com/nmluci/stellar-payment-lite/pkg/errs"
)

type AuthLoginHandler func(context.Context, *dto.AuthLoginPayload) (*dto.AuthResponse, error)

func HandleAuthLogin(handler AuthLoginHandler) echo.HandlerFunc {
	return func(c echo.Context) error {
		req := &dto.AuthLoginPayload{}
		if err := c.Bind(req); err != nil {
			return echttputil.WriteErrorResponse(c, errs.ErrBrokenUserReq)
		}

		res, err := handler(c.Request().Context(), req)
		if err != nil {
			return echttputil.WriteErrorResponse(c, err)
		}

		return echttputil.WriteSuccessResponse(c, res)
	}
}
