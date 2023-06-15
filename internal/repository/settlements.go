package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/nmluci/stellar-payment-lite/internal/indto"
	"github.com/nmluci/stellar-payment-lite/internal/model"
)

func (r *repository) FindSettlementByMerchantID(ctx context.Context, merchantID int64) (res []*indto.SettlementDetail, err error) {
	stmt, args, err := squirrel.Select("s.id settlement_id", "t.id trx_id", "t.trx_datetime", "s.merchant_id", "m.name merchant_name", "s.nominal", "s.status").
		From("settlements s").LeftJoin("transactions t on s.transaction_id = t.id").LeftJoin("merchants m on s.merchant_id = m.id").Where(squirrel.And{
		squirrel.Eq{"s.merchant_id": merchantID},
	}).ToSql()
	if err != nil {
		r.logger.Error().Err(err).Msg("squirrel err")
		return
	}

	res = []*indto.SettlementDetail{}
	rows, err := r.mariaDB.QueryxContext(ctx, stmt, args...)
	if err != nil {
		r.logger.Error().Err(err).Msg("sql err")
		return
	}

	for rows.Next() {
		temp := &indto.SettlementDetail{}
		if err = rows.StructScan(temp); err != nil {
			r.logger.Error().Err(err).Msg("sql map err")
			return
		}

		res = append(res, temp)
	}

	return
}

func (r *repository) InsertTransactionMerchant(ctx context.Context, payload *model.Transaction) (err error) {
	tx, err := r.mariaDB.BeginTx(ctx, nil)
	if err != nil {
		r.logger.Error().Err(err).Msg("tx err")
		return
	}
	defer tx.Rollback()

	err = r.UpdateAccountBalanceTx(ctx, tx, &model.Account{ID: payload.AccountID, Balance: payload.Nominal + payload.TransactionFee})
	if err != nil {
		r.logger.Error().Err(err).Send()
		return
	}

	trxID, err := r.InsertTransactionTx(ctx, tx, payload)
	if err != nil {
		r.logger.Error().Err(err).Send()
		return
	}

	settlements := &model.Settlement{
		TransactionID: trxID,
		MerchantID:    payload.RecipientID,
		Nominal:       payload.Nominal,
		Status:        1,
	}
	err = r.InsertSettlementTx(ctx, tx, settlements)
	if err != nil {
		r.logger.Error().Err(err).Send()
		return
	}

	if err = tx.Commit(); err != nil {
		r.logger.Error().Err(err).Msg("tx err")
	}

	return
}

func (r *repository) InsertSettlementTx(ctx context.Context, tx *sql.Tx, payload *model.Settlement) (err error) {
	stmt, args, err := squirrel.Insert("settlements").Columns("transaction_id", "merchant_id", "nominal", "status", "created_at", "updated_at").Values(
		payload.TransactionID, payload.MerchantID, payload.Nominal, payload.Status, time.Now(), time.Now(),
	).ToSql()
	if err != nil {
		r.logger.Error().Err(err).Msg("squirrel err")
		return
	}

	_, err = tx.ExecContext(ctx, stmt, args...)
	if err != nil {
		r.logger.Error().Err(err).Msg("sql err")
		return
	}

	return
}
