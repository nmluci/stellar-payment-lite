package handler

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/nmluci/stellar-payment-lite/internal/util/echttputil"
	"github.com/nmluci/stellar-payment-lite/pkg/dto"
	"github.com/nmluci/stellar-payment-lite/pkg/errs"
)

type GetMerchantHandler func(context.Context) ([]*dto.MerchantResponse, error)

func HandleGetMerchant(handler GetMerchantHandler) echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := handler(c.Request().Context())
		if err != nil {
			return echttputil.WriteErrorResponse(c, err)
		}

		return echttputil.WriteSuccessResponse(c, res)
	}
}

type CreateMerchantHandler func(context.Context, *dto.MerchantRequest) error

func HandleCreateMerchant(handler CreateMerchantHandler) echo.HandlerFunc {
	return func(c echo.Context) error {
		req := &dto.MerchantRequest{}
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

type FindMerchantSettlementsHandler func(context.Context, *dto.MerchantQueryParams) ([]*dto.MerchantSettlement, error)

func HandleFindMerchantSettlements(handler FindMerchantSettlementsHandler) echo.HandlerFunc {
	return func(c echo.Context) error {
		req := &dto.MerchantQueryParams{}
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
