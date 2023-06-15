package service

import (
	"context"

	"github.com/nmluci/stellar-payment-lite/internal/model"
	"github.com/nmluci/stellar-payment-lite/internal/util/ctxutil"
	"github.com/nmluci/stellar-payment-lite/pkg/dto"
	"github.com/nmluci/stellar-payment-lite/pkg/errs"
)

func (s *service) GetAccountByUser(ctx context.Context) (res []*dto.AccountResponse, err error) {
	ctxmeta := ctxutil.GetUserCTX(ctx)

	data, err := s.repository.FindAccountByUserID(ctx, ctxmeta.CustomerID)
	if err != nil {
		s.logger.Error().Err(err).Send()
		return
	} else if data == nil {
		return nil, errs.ErrNotFound
	}

	res = []*dto.AccountResponse{}

	for _, v := range data {
		res = append(res, &dto.AccountResponse{
			AccountID:   v.ID,
			AccountNo:   v.AccountNo,
			AccountType: v.AccountType,
			CardNumber:  v.CardNumber,
			Balance:     v.Balance,
		})
	}

	return
}

func (s *service) CreateAccount(ctx context.Context, payload *dto.AccountRequest) (err error) {
	ctxmeta := ctxutil.GetUserCTX(ctx)

	if payload.CustomerID == 0 {
		if ctxmeta == nil {
			return errs.ErrBadRequest
		}

		payload.CustomerID = ctxmeta.CustomerID
	}

	err = s.repository.InsertAccount(ctx, &model.Account{
		CustomerID:  payload.CustomerID,
		AccountNo:   payload.AccountNo,
		AccountType: payload.AccountType,
		CardNumber:  payload.CardNumber,
		CVV:         "",
		PIN:         "",
		Balance:     100000,
	})

	return
}
