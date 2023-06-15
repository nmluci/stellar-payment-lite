package service

import (
	"context"

	"github.com/nmluci/stellar-payment-lite/internal/commonkey"
	"github.com/nmluci/stellar-payment-lite/internal/model"
	"github.com/nmluci/stellar-payment-lite/internal/util/ctxutil"
	"github.com/nmluci/stellar-payment-lite/internal/util/timeutil"
	"github.com/nmluci/stellar-payment-lite/pkg/dto"
	"github.com/nmluci/stellar-payment-lite/pkg/errs"
)

func (s *service) GetCustomerByID(ctx context.Context, param *dto.CustomerQueryParams) (res *dto.CustomerResponse, err error) {
	ctxMeta := ctxutil.GetUserCTX(ctx)
	if ctxMeta.RoleID != commonkey.ROLE_ADMIN && ctxMeta.CustomerID != param.CustomerID {
		return nil, errs.ErrNoAccess
	}

	data, err := s.repository.FindCustomerByID(ctx, param.CustomerID)
	if err != nil {
		s.logger.Error().Err(err).Send()
		return
	}

	if data == nil {
		return nil, errs.ErrNotFound
	}

	res = &dto.CustomerResponse{
		CustomerID: data.ID,
		LegalName:  data.LegalName,
		Address:    data.Address,
		Phone:      data.Phone,
		Birthplace: data.Birthplace,
		Birthdate:  timeutil.FormatDate(data.Birthdate),
		NIK:        data.NIK,
		Occupation: data.Occupation,
		KTPUrl:     data.KTPUrl,
	}

	return
}

func (s *service) UpdateCustomer(ctx context.Context, param *dto.CustomerQueryParams, payload *dto.CustomerPayload) (err error) {
	ctxMeta := ctxutil.GetUserCTX(ctx)
	if ctxMeta.RoleID != commonkey.ROLE_ADMIN && ctxMeta.CustomerID != param.CustomerID {
		return errs.ErrNoAccess
	}

	cust := &model.Customer{
		ID:         param.CustomerID,
		LegalName:  payload.LegalName,
		Address:    payload.Address,
		Phone:      payload.Phone,
		Birthplace: payload.Birthplace,
		Birthdate:  timeutil.ParseDate(payload.Birthdate),
		NIK:        payload.NIK,
		Occupation: payload.Occupation,
		KTPUrl:     payload.KTPUrl,
	}

	err = s.repository.UpdateCustomer(ctx, cust)
	if err != nil {
		s.logger.Error().Err(err).Send()
		return
	}

	return
}
