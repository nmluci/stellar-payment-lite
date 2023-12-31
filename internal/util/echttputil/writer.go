package echttputil

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nmluci/stellar-payment-lite/pkg/dto"
	"github.com/nmluci/stellar-payment-lite/pkg/errs"
)

func WriteSuccessResponse(ec echo.Context, data interface{}) error {
	return ec.JSON(http.StatusOK, dto.BaseResponse{
		Data:   data,
		Errors: nil,
	})
}

func WriteErrorResponse(ec echo.Context, err error) error {
	errResp := errs.GetErrorResp(err)
	return ec.JSON(errResp.Status, dto.BaseResponse{
		Data:   nil,
		Errors: errResp,
	})
}

func WriteFileAttachment(ec echo.Context, path string, filename string) error {
	return ec.Attachment(path, filename)
}
