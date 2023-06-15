package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/nmluci/stellar-payment-lite/internal/indto"
	"github.com/nmluci/stellar-payment-lite/internal/model"
)

func (r *repository) FindUserByUsername(ctx context.Context, username string) (res *indto.UserDetail, err error) {
	stmt, args, err := squirrel.Select("u.id user_id", "u.username", "c.legal_name", "u.password password", "u.role_id role_id").From("users u").
		LeftJoin("customers c on u.customer_id = c.id").Where(squirrel.And{
		squirrel.Eq{"u.username": username},
		squirrel.Eq{"u.deleted_at": nil},
	}).ToSql()
	if err != nil {
		r.logger.Error().Err(err).Send()
		return
	}

	res = &indto.UserDetail{}
	err = r.mariaDB.QueryRowxContext(ctx, stmt, args...).StructScan(res)
	if err != nil && err != sql.ErrNoRows {
		r.logger.Error().Err(err).Send()
		return
	} else if err == sql.ErrNoRows {
		return nil, nil
	}

	return
}

func (r *repository) FindUserByID(ctx context.Context, id int64) (res *indto.UserDetail, err error) {
	stmt, args, err := squirrel.Select("u.id user_id", "c.id customer_id", "u.username", "c.legal_name", "u.password password", "u.role_id role_id").From("users u").
		LeftJoin("customers c on u.customer_id = c.id").Where(squirrel.And{
		squirrel.Eq{"u.id": id},
		squirrel.Eq{"u.deleted_at": nil},
	}).ToSql()
	if err != nil {
		r.logger.Error().Err(err).Send()
		return
	}

	res = &indto.UserDetail{}
	err = r.mariaDB.QueryRowxContext(ctx, stmt, args...).StructScan(res)
	if err != nil && err != sql.ErrNoRows {
		r.logger.Error().Err(err).Send()
		return
	} else if err == sql.ErrNoRows {
		return nil, nil
	}

	return
}

func (r *repository) InsertUser(ctx context.Context, usr *model.User, cust *model.Customer) (err error) {
	tx, err := r.mariaDB.BeginTx(ctx, nil)
	if err != nil {
		r.logger.Error().Err(err).Send()
		return
	}
	defer tx.Rollback()

	cid, err := r.InsertCustomerTx(ctx, tx, cust)
	if err != nil {
		r.logger.Error().Err(err).Send()
		return
	}

	usr.CustomerID = cid
	err = r.InsertUserTx(ctx, tx, usr)
	if err != nil {
		r.logger.Error().Err(err).Send()
		return
	}

	if err = tx.Commit(); err != nil {
		r.logger.Error().Err(err).Send()
		return
	}

	return
}

func (r *repository) InsertUserTx(ctx context.Context, tx *sql.Tx, usr *model.User) (err error) {
	stmt, args, err := squirrel.Insert("users").Columns("role_id", "customer_id", "username", "password", "created_at", "updated_at").Values(
		usr.RoleID, usr.CustomerID, usr.Username, usr.Password, time.Now(), time.Now()).ToSql()
	if err != nil {
		r.logger.Error().Err(err).Send()
		return
	}

	_, err = tx.ExecContext(ctx, stmt, args...)
	if err != nil {
		r.logger.Error().Err(err).Send()
		return
	}

	return
}

func (r *repository) UpdateUser(ctx context.Context, usr *model.User) (err error) {
	stmt, args, err := squirrel.Update("users").SetMap(map[string]interface{}{
		"password": squirrel.Expr("coalesce(nullif(?, ''), password)", usr.Password),
	}).Where(squirrel.And{
		squirrel.Eq{"id": usr.UserID},
	}).ToSql()
	if err != nil {
		r.logger.Error().Err(err).Send()
		return
	}

	_, err = r.mariaDB.ExecContext(ctx, stmt, args...)
	if err != nil {
		r.logger.Error().Err(err).Send()
		return
	}

	return
}
