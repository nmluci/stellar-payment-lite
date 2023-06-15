package handler

import (
	"context"

	"github.com/labstack/echo/v4"

	"github.com/nmluci/stellar-payment-lite/internal/util/echttputil"
	"github.com/nmluci/stellar-payment-lite/internal/util/structutil"
	"github.com/nmluci/stellar-payment-lite/pkg/dto"
	"github.com/nmluci/stellar-payment-lite/pkg/errs"
)

type RegisterUserHandler func(context.Context, *dto.UserRegistrationPayload) error

func HandleRegisterUser(handler RegisterUserHandler) echo.HandlerFunc {
	return func(c echo.Context) error {
		req := &dto.UserRegistrationPayload{}
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

type GetUserDetailByIDHandler func(context.Context, *dto.UserQueryParams) (*dto.UserResponse, error)

func HandleGetUserDetailByID(handler GetUserDetailByIDHandler) echo.HandlerFunc {
	return func(c echo.Context) error {
		req := &dto.UserQueryParams{}
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

type UpdateUserDetailHandler func(context.Context, *dto.UserQueryParams, *dto.UserPayload) error

func HandleUpdateUserDetail(handler UpdateUserDetailHandler) echo.HandlerFunc {
	return func(c echo.Context) error {
		params := &dto.UserQueryParams{
			UserID: structutil.StringToInt64(c.Param("userID")),
		}

		req := &dto.UserPayload{}
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
