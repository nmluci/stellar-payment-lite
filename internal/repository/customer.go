package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/nmluci/stellar-payment-lite/internal/model"
)

func (r *repository) FindCustomerByID(ctx context.Context, id int64) (res *model.Customer, err error) {
	stmt, args, err := squirrel.Select("id", "legal_name", "address", "phone", "birthplace", "birthdate", "nik", "occupation", "ktp_url").From("customers").Where(squirrel.And{
		squirrel.Eq{"id": id},
		squirrel.Eq{"deleted_at": nil},
	}).ToSql()
	if err != nil {
		r.logger.Error().Err(err).Msg("squirrel err")
		return
	}

	res = &model.Customer{}
	err = r.mariaDB.QueryRowxContext(ctx, stmt, args...).StructScan(res)
	if err != nil && err != sql.ErrNoRows {
		r.logger.Error().Err(err).Msg("sql err")
		return
	} else if err == sql.ErrNoRows {
		return nil, nil
	}

	return
}

func (r *repository) InsertCustomerTx(ctx context.Context, tx *sql.Tx, payload *model.Customer) (customerID int64, err error) {
	stmt, args, err := squirrel.Insert("customers").Columns("legal_name", "address", "phone", "birthplace", "birthdate", "nik", "occupation", "ktp_url", "created_at", "updated_at").Values(
		payload.LegalName, payload.Address, payload.Phone, payload.Birthplace, payload.Birthdate, payload.NIK, payload.Occupation, payload.KTPUrl, time.Now(), time.Now(),
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

	customerID, err = res.LastInsertId()
	if err != nil {
		r.logger.Error().Err(err).Msg("sql lid err")
		return
	}

	return
}

func (r *repository) UpdateCustomer(ctx context.Context, payload *model.Customer) (err error) {
	stmt, args, err := squirrel.Update("customers").SetMap(map[string]interface{}{
		"legal_name": squirrel.Expr("coalesce(nullif(?, ''), legal_name)", payload.LegalName),
		"address":    squirrel.Expr("coalesce(nullif(?, ''), address)", payload.Address),
		"phone":      squirrel.Expr("coalesce(nullif(?, ''), phone)", payload.Phone),
		"birthplace": squirrel.Expr("coalesce(nullif(?, ''), birthplace)", payload.Birthplace),
		"birthdate":  squirrel.Expr("coalesce(nullif(?, ''), birthdate)", payload.Birthdate),
		"nik":        squirrel.Expr("coalesce(nullif(?, ''), nik)", payload.NIK),
		"occupation": squirrel.Expr("coalesce(nullif(?, ''), occupation)", payload.Occupation),
		"ktp_url":    squirrel.Expr("coalesce(nullif(?, ''), ktp_url)", payload.KTPUrl),
		"updated_at": time.Now(),
	}).Where(squirrel.Eq{"id": payload.ID}).ToSql()
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
