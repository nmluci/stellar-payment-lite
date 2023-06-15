package service

import (
	"context"

	"github.com/nmluci/stellar-payment-lite/internal/model"
	"github.com/nmluci/stellar-payment-lite/internal/util/timeutil"
	"github.com/nmluci/stellar-payment-lite/pkg/dto"
	"github.com/nmluci/stellar-payment-lite/pkg/errs"
)

func (s *service) GetMerchants(ctx context.Context) (res []*dto.MerchantResponse, err error) {
	data, err := s.repository.FindMerchants(ctx)
	if err != nil {
		s.logger.Error().Err(err).Send()
		return
	}

	res = []*dto.MerchantResponse{}
	for _, v := range data {
		res = append(res, &dto.MerchantResponse{
			ID:           v.ID,
			Name:         v.Name,
			Address:      v.Address,
			Phone:        v.Phone,
			MerchantCode: v.MerchantCode,
		})
	}

	return
}

func (s *service) CreateMerchant(ctx context.Context, payload *dto.MerchantRequest) (err error) {
	err = s.repository.InsertMerchant(ctx, &model.Merchant{
		Name:         payload.Name,
		Address:      payload.Address,
		Phone:        payload.Phone,
		MerchantCode: payload.MerchantCode,
	})
	if err != nil {
		s.logger.Error().Err(err).Send()
		return
	}

	return
}

func (s *service) FindMerchantSettlements(ctx context.Context, payload *dto.MerchantQueryParams) (res []*dto.MerchantSettlement, err error) {
	if exists, err := s.repository.FindMerchantByID(ctx, payload.MerchantID); err != nil {
		s.logger.Error().Err(err).Send()
		return nil, err
	} else if exists == nil {
		return nil, errs.ErrNotFound
	}

	data, err := s.repository.FindSettlementByMerchantID(ctx, payload.MerchantID)
	if err != nil {
		s.logger.Error().Err(err).Send()
		return
	}

	res = []*dto.MerchantSettlement{}
	for _, v := range data {
		res = append(res, &dto.MerchantSettlement{
			ID:           v.ID,
			TrxID:        v.TrxID,
			TrxDatetime:  timeutil.FormatLocaltime(v.TrxDatetime),
			MerchantID:   v.MerchantID,
			MerchantName: v.MerchantName,
			Nominal:      v.Nominal,
			Status:       v.Status,
		})
	}

	return
}
