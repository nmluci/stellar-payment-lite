package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/nmluci/stellar-payment-lite/internal/model"
)

func (r *repository) FindMerchants(ctx context.Context) (res []*model.Merchant, err error) {
	stmt, args, err := squirrel.Select("id", "name", "address", "phone", "merchant_code").From("merchants").ToSql()
	if err != nil {
		r.logger.Error().Err(err).Msg("squirrel err")
		return
	}

	res = []*model.Merchant{}
	rows, err := r.mariaDB.QueryxContext(ctx, stmt, args...)
	if err != nil {
		r.logger.Error().Err(err).Msg("sql err")
		return
	}

	for rows.Next() {
		temp := &model.Merchant{}
		if err = rows.StructScan(temp); err != nil {
			r.logger.Error().Err(err).Msg("sql err")
			return
		}

		res = append(res, temp)
	}

	return
}

func (r *repository) FindMerchantByID(ctx context.Context, merchantID int64) (res *model.Merchant, err error) {
	stmt, args, err := squirrel.Select("id", "name", "address", "phone", "merchant_code").From("merchants").Where(squirrel.And{
		squirrel.Eq{"id": merchantID},
	}).ToSql()
	if err != nil {
		r.logger.Error().Err(err).Msg("squirrel err")
		return
	}

	res = &model.Merchant{}
	err = r.mariaDB.QueryRowxContext(ctx, stmt, args...).StructScan(res)
	if err != nil && err != sql.ErrNoRows {
		r.logger.Error().Err(err).Msg("sql err")
		return
	} else if err == sql.ErrNoRows {
		return nil, nil
	}

	return
}

func (r *repository) InsertMerchant(ctx context.Context, payload *model.Merchant) (err error) {
	stmt, args, err := squirrel.Insert("merchants").Columns("name", "address", "phone", "merchant_code", "created_at", "updated_at").Values(
		payload.Name, payload.Address, payload.Phone, payload.MerchantCode, time.Now(), time.Now(),
	).ToSql()
	if err != nil {
		r.logger.Error().Err(err).Msg("squirrel err")
		return
	}

	_, err = r.mariaDB.ExecContext(ctx, stmt, args...)
	if err != nil {
		r.logger.Error().Err(err).Msg("sql err")
		return
	}

	return
}
