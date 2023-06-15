package repository

import (
	"context"
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/nmluci/stellar-payment-lite/internal/indto"
)

func (r *repository) FindUserRoleByID(ctx context.Context, userID int64) (res *indto.UserRole, err error) {
	stmt, args, err := squirrel.Select("u.id user_id", "u.customer_id", "u.username", "u.role_id").From("users u").Where(squirrel.And{
		squirrel.Eq{"u.id": userID},
		squirrel.Eq{"u.deleted_at": nil},
	}).ToSql()
	if err != nil {
		r.logger.Error().Err(err).Send()
		return
	}

	res = &indto.UserRole{}
	err = r.mariaDB.QueryRowxContext(ctx, stmt, args...).StructScan(res)
	if err != nil && err != sql.ErrNoRows {
		r.logger.Error().Err(err).Send()
		return nil, err
	} else if err == sql.ErrNoRows {
		return nil, nil
	}

	return
}
