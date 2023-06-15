package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/nmluci/stellar-payment-lite/internal/model"
)

func (r *repository) FindAccountByUserID(ctx context.Context, custID int64) (res []*model.Account, err error) {
	stmt, args, err := squirrel.Select("id", "customer_id", "account_no", "account_type", "card_number", "cvv", "pin", "balance").From("accounts").
		Where(squirrel.And{squirrel.Eq{"customer_id": custID}}).ToSql()
	if err != nil {
		r.logger.Error().Err(err).Msg("squirrel err")
		return
	}

	res = []*model.Account{}
	rows, err := r.mariaDB.QueryxContext(ctx, stmt, args...)
	if err != nil {
		r.logger.Error().Err(err).Msg("sql err")
		return
	}

	for rows.Next() {
		temp := &model.Account{}
		if err = rows.StructScan(temp); err != nil {
			r.logger.Error().Err(err).Msg("sql map err")
			return
		}

		res = append(res, temp)
	}

	return
}

func (r *repository) FindAccountByID(ctx context.Context, accountID int64) (res *model.Account, err error) {
	stmt, args, err := squirrel.Select("id", "customer_id", "account_no", "account_type", "card_number", "cvv", "pin", "balance").From("accounts").
		Where(squirrel.And{squirrel.Eq{"id": accountID}}).ToSql()
	if err != nil {
		r.logger.Error().Err(err).Msg("squirrel err")
		return
	}

	res = &model.Account{}
	err = r.mariaDB.QueryRowxContext(ctx, stmt, args...).StructScan(res)
	if err != nil {
		r.logger.Error().Err(err).Msg("sql err")
		return
	}

	return
}

func (r *repository) InsertAccount(ctx context.Context, payload *model.Account) (err error) {
	stmt, args, err := squirrel.Insert("accounts").Columns("customer_id", "account_no", "account_type", "card_number", "cvv", "pin", "balance", "created_at", "updated_at").Values(
		payload.CustomerID, payload.AccountNo, payload.AccountType, payload.CardNumber, payload.CVV, payload.PIN, payload.Balance, time.Now(), time.Now(),
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

func (r *repository) UpdateAccount(ctx context.Context, payload *model.Account) (err error) {
	tx, err := r.mariaDB.BeginTx(ctx, nil)
	if err != nil {
		r.logger.Error().Err(err).Msg("tx err")
		return
	}
	defer tx.Rollback()

	err = r.UpdateAccountTx(ctx, tx, payload)
	if err != nil {
		return
	}

	if err = tx.Commit(); err != nil {
		r.logger.Error().Err(err).Msg("tx err")
		return
	}

	return
}

func (r *repository) UpdateAccountTx(ctx context.Context, tx *sql.Tx, payload *model.Account) (err error) {
	stmt, args, err := squirrel.Update("accounts").SetMap(map[string]interface{}{
		"pin":        squirrel.Expr("coalesce(nullif(?, ''), pin)", payload.PIN),
		"balance":    squirrel.Expr("coalesce(nullif(?, ''), balance)", payload.Balance),
		"updated_at": time.Now(),
	}).Where(squirrel.Eq{"id": payload.ID}).ToSql()
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

func (r *repository) UpdateAccountBalanceTx(ctx context.Context, tx *sql.Tx, payload *model.Account) (err error) {
	stmt, args, err := squirrel.Update("accounts").SetMap(map[string]interface{}{
		"balance":    squirrel.Expr("(balance - ?)", payload.Balance),
		"updated_at": time.Now(),
	}).Where(squirrel.Eq{"id": payload.ID}).ToSql()
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
