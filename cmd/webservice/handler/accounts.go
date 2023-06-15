package handler

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/nmluci/stellar-payment-lite/internal/util/echttputil"
	"github.com/nmluci/stellar-payment-lite/pkg/dto"
	"github.com/nmluci/stellar-payment-lite/pkg/errs"
)

type GetAccountByUserHandler func(context.Context) ([]*dto.AccountResponse, error)

func HandleGetAccountByUser(handler GetAccountByUserHandler) echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := handler(c.Request().Context())
		if err != nil {
			return echttputil.WriteErrorResponse(c, err)
		}

		return echttputil.WriteSuccessResponse(c, res)
	}
}

type CreateAccountHandler func(context.Context, *dto.AccountRequest) error

func HandleCreateAccount(handler CreateAccountHandler) echo.HandlerFunc {
	return func(c echo.Context) error {
		req := &dto.AccountRequest{}
		if err := c.Bind(req); err != nil {
			return echttputil.WriteErrorResponse(c, errs.ErrBrokenUserReq)
		}

		err := handler(c.Request().Context(), req)
		if err != nil {
			return echttputil.WriteErrorResponse(c, err)
		}

		return echttputil.WriteSuccessResponse(c, nil)
	}
}
