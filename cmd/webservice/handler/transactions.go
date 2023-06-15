package handler

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/nmluci/stellar-payment-lite/internal/util/echttputil"
	"github.com/nmluci/stellar-payment-lite/pkg/dto"
	"github.com/nmluci/stellar-payment-lite/pkg/errs"
)

type GetTransactionHistoryByUserHandler func(context.Context, *dto.AccountQueryParams) ([]*dto.TransactionResponse, error)

func HandleGetTransactionHistoryByUser(handler GetTransactionHistoryByUserHandler) echo.HandlerFunc {
	return func(c echo.Context) error {
		req := &dto.AccountQueryParams{}
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

type CreateTransactionP2PHandler func(context.Context, *dto.TransactionRequest) error

func HandleCreateTransactionP2P(handler CreateTransactionP2PHandler) echo.HandlerFunc {
	return func(c echo.Context) error {
		req := &dto.TransactionRequest{}
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

type CreateTransactionMerchantHandler func(context.Context, *dto.TransactionRequest) error

func HandleCreateTransactionMerchant(handler CreateTransactionMerchantHandler) echo.HandlerFunc {
	return func(c echo.Context) error {
		req := &dto.TransactionRequest{}
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
