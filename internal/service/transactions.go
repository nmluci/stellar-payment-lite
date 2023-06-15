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

func (s *service) GetTransactionHistoryByUser(ctx context.Context, params *dto.AccountQueryParams) (res []*dto.TransactionResponse, err error) {
	data, err := s.repository.FindTransactionByAccountID(ctx, params.AccountID)
	if err != nil {
		s.logger.Error().Err(err).Send()
		return nil, err
	}

	res = []*dto.TransactionResponse{}
	for _, v := range data {
		res = append(res, &dto.TransactionResponse{
			ID:            v.ID,
			AccountID:     v.AccountID,
			AccountOwner:  v.AccountOwner,
			RecipientID:   v.RecipientID,
			RecipientName: v.RecipientName,
			TrxType:       commonkey.TRX_TYPE_MAP[v.TrxType],
			TrxDatetime:   timeutil.FormatDate(v.TrxDatetime),
			Nominal:       v.Nominal,
			TrxFee:        v.TrxFee,
		})
	}

	return
}

func (s *service) CreateTransactionP2P(ctx context.Context, params *dto.TransactionRequest) (err error) {
	ctxmeta := ctxutil.GetUserCTX(ctx)

	accMeta, err := s.repository.FindAccountByID(ctx, params.AccountID)
	if err != nil {
		s.logger.Error().Err(err).Send()
		return
	}

	if accMeta.CustomerID != ctxmeta.CustomerID {
		return errs.ErrNoAccess
	}

	if params.Nominal > (accMeta.Balance * 0.9) {
		s.logger.Error().Float32("has", accMeta.Balance).Float32("90%% balance", accMeta.Balance*0.9).Float32("required", params.Nominal).Msg("not enough balance")
		return errs.ErrBadRequest
	}

	if exists, err := s.repository.FindAccountByID(ctx, params.RecipientID); err != nil {
		s.logger.Error().Err(err).Send()
		return err
	} else if exists == nil {
		s.logger.Error().Msg("recipient id not found")
		return errs.ErrBadRequest
	}

	params.TransactionFee = params.Nominal * 0.1
	err = s.repository.InsertTransactionP2P(ctx, &model.Transaction{
		AccountID:      params.AccountID,
		RecipientID:    params.RecipientID,
		TrxType:        0,
		TrxDatetime:    timeutil.ParseLocaltime(params.TrxDatetime),
		TrxStatus:      1,
		Nominal:        params.Nominal,
		TransactionFee: params.TransactionFee,
	})
	if err != nil {
		s.logger.Error().Err(err).Send()
		return
	}

	return
}

func (s *service) CreateTransactionMerchant(ctx context.Context, params *dto.TransactionRequest) (err error) {
	ctxmeta := ctxutil.GetUserCTX(ctx)

	accMeta, err := s.repository.FindAccountByID(ctx, params.AccountID)
	if err != nil {
		s.logger.Error().Err(err).Send()
		return
	}

	if accMeta.CustomerID != ctxmeta.CustomerID {
		return errs.ErrNoAccess
	}

	if params.Nominal > (accMeta.Balance * 0.9) {
		s.logger.Error().Float32("has", accMeta.Balance).Float32("90%% balance", accMeta.Balance*0.9).Float32("required", params.Nominal).Msg("not enough balance")
		return errs.ErrBadRequest
	}

	if exists, err := s.repository.FindMerchantByID(ctx, params.RecipientID); err != nil {
		s.logger.Error().Err(err).Send()
		return err
	} else if exists == nil {
		s.logger.Error().Msg("merchant id not found")
		return errs.ErrBadRequest
	}

	params.TransactionFee = params.Nominal * 0.17
	err = s.repository.InsertTransactionMerchant(ctx, &model.Transaction{
		AccountID:      params.AccountID,
		RecipientID:    params.RecipientID,
		TrxType:        1,
		TrxDatetime:    timeutil.ParseLocaltime(params.TrxDatetime),
		TrxStatus:      1,
		Nominal:        params.Nominal,
		TransactionFee: params.TransactionFee,
	})
	if err != nil {
		s.logger.Error().Err(err).Send()
		return
	}

	return
}
