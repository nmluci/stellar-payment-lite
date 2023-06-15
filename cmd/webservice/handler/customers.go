package handler

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/nmluci/stellar-payment-lite/internal/util/echttputil"
	"github.com/nmluci/stellar-payment-lite/internal/util/structutil"
	"github.com/nmluci/stellar-payment-lite/pkg/dto"
	"github.com/nmluci/stellar-payment-lite/pkg/errs"
)

type GetCustomerDetailByIDHandler func(context.Context, *dto.CustomerQueryParams) (*dto.CustomerResponse, error)

func HandleGetCustomerDetailByID(handler GetCustomerDetailByIDHandler) echo.HandlerFunc {
	return func(c echo.Context) error {
		req := &dto.CustomerQueryParams{}
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

type UpdateCustomerDetailHandler func(context.Context, *dto.CustomerQueryParams, *dto.CustomerPayload) error

func HandleUpdateCustomerDetail(handler UpdateCustomerDetailHandler) echo.HandlerFunc {
	return func(c echo.Context) error {
		params := &dto.CustomerQueryParams{
			CustomerID: structutil.StringToInt64(c.Param("customerID")),
		}

		req := &dto.CustomerPayload{}
		if err := c.Bind(req); err != nil {
			return echttputil.WriteErrorResponse(c, errs.ErrBrokenUserReq)
		}

		err := handler(c.Request().Context(), params, req)
		if err != nil {
			return echttputil.WriteErrorResponse(c, err)
		}

		return echttputil.WriteSuccessResponse(c, nil)
	}
}
