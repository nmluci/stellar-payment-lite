package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/nmluci/stellar-payment-lite/internal/indto"
	"github.com/nmluci/stellar-payment-lite/internal/model"
)

func (r *repository) FindTransactionByAccountID(ctx context.Context, accountID int64) (res []*indto.TranasctionHistory, err error) {
	stmt, args, err := squirrel.Select("t.id trx_id", "t.account_id", "c1.legal_name account_owner",
		"t.recipient_id", "coalesce(m2.name, c2.legal_name) recipient_name", "t.trx_type", "t.trx_datetime", "t.nominal", "t.transaction_fee trx_fee").From("transactions t").
		LeftJoin("accounts a1 on t.account_id = a1.id").LeftJoin("customers c1 on a1.customer_id = c1.id").
		LeftJoin("merchants m2 on t.recipient_id = m2.id and t.trx_type = 1").
		LeftJoin("accounts a2 on t.recipient_id = a2.id").LeftJoin("customers c2 on a2.customer_id = c2.id").
		Where(squirrel.Or{
			squirrel.Eq{"t.account_id": accountID},
			squirrel.And{
				squirrel.Eq{"t.recipient_id": accountID},
				squirrel.Eq{"t.trx_type": 0},
			},
		}).ToSql()
	if err != nil {
		r.logger.Error().Err(err).Msg("squirrel err")
		return
	}

	res = []*indto.TranasctionHistory{}
	rows, err := r.mariaDB.QueryxContext(ctx, stmt, args...)
	if err != nil {
		r.logger.Error().Err(err).Msg("sql err")
		return
	}

	for rows.Next() {
		temp := &indto.TranasctionHistory{}
		if err = rows.StructScan(temp); err != nil {
			r.logger.Error().Err(err).Msg("sql map err")
			return
		}

		res = append(res, temp)
	}

	return
}

func (r *repository) InsertTransactionP2P(ctx context.Context, payload *model.Transaction) (err error) {
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

	err = r.UpdateAccountBalanceTx(ctx, tx, &model.Account{ID: payload.RecipientID, Balance: -payload.Nominal})
	if err != nil {
		r.logger.Error().Err(err).Send()
		return
	}

	_, err = r.InsertTransactionTx(ctx, tx, payload)
	if err != nil {
		r.logger.Error().Err(err).Send()
		return
	}

	if err = tx.Commit(); err != nil {
		r.logger.Error().Err(err).Msg("tx err")
		return
	}

	return
}

func (r *repository) InsertTransactionTx(ctx context.Context, tx *sql.Tx, payload *model.Transaction) (trxID int64, err error) {
	stmt, args, err := squirrel.Insert("transactions").Columns("account_id", "recipient_id", "trx_type", "trx_datetime", "trx_status",
		"nominal", "transaction_fee", "created_at", "updated_at").Values(
		payload.AccountID, payload.RecipientID, payload.TrxType, payload.TrxDatetime, payload.TrxStatus,
		payload.Nominal, payload.TransactionFee, time.Now(), time.Now(),
	).ToSql()
	if err != nil {
		r.logger.Error().Err(err).Msg("squirrel err")
		return
	}

	res, err := tx.ExecContext(ctx, stmt, args...)
	if err != nil {
		r.logger.Error().Err(err).Msg("sql err")
		return
	}

	trxID, err = res.LastInsertId()
	if err != nil {
		r.logger.Error().Err(err).Msg("sql lid err")
		return
	}

	return
}
