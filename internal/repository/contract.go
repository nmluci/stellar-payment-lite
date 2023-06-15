package repository

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
	"github.com/nmluci/stellar-payment-lite/internal/indto"
	"github.com/nmluci/stellar-payment-lite/internal/model"
	"github.com/rs/zerolog"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	FindUserRoleByID(ctx context.Context, userID int64) (res *indto.UserRole, err error)

	// Users
	FindUserByUsername(ctx context.Context, username string) (res *indto.UserDetail, err error)
	FindUserByID(ctx context.Context, id int64) (res *indto.UserDetail, err error)
	InsertUser(ctx context.Context, usr *model.User, cust *model.Customer) (err error)
	UpdateUser(ctx context.Context, usr *model.User) (err error)

	// Customers
	FindCustomerByID(ctx context.Context, id int64) (res *model.Customer, err error)
	UpdateCustomer(ctx context.Context, payload *model.Customer) (err error)

	// Accounts
	FindAccountByUserID(ctx context.Context, userID int64) (res []*model.Account, err error)
	FindAccountByID(ctx context.Context, accountID int64) (res *model.Account, err error)
	InsertAccount(ctx context.Context, payload *model.Account) (err error)
	UpdateAccount(ctx context.Context, payload *model.Account) (err error)

	// Merchants
	FindMerchants(ctx context.Context) (res []*model.Merchant, err error)
	FindMerchantByID(ctx context.Context, merchantID int64) (res *model.Merchant, err error)
	InsertMerchant(ctx context.Context, payload *model.Merchant) (err error)

	// Transactions
	FindTransactionByAccountID(ctx context.Context, accountID int64) (res []*indto.TranasctionHistory, err error)
	InsertTransactionP2P(ctx context.Context, payload *model.Transaction) (err error)
	InsertTransactionMerchant(ctx context.Context, payload *model.Transaction) (err error)

	// Settlements
	FindSettlementByMerchantID(ctx context.Context, merchantID int64) (res []*indto.SettlementDetail, err error)
}

type repository struct {
	mariaDB *sqlx.DB
	redis   *redis.Client
	logger  zerolog.Logger
	conf    *repositoryConfig
}

type repositoryConfig struct {
}

type NewRepositoryParams struct {
	Logger  zerolog.Logger
	MariaDB *sqlx.DB
	MongoDB *mongo.Database
	Redis   *redis.Client
}

func NewRepository(params *NewRepositoryParams) Repository {
	return &repository{
		logger:  params.Logger,
		conf:    &repositoryConfig{},
		mariaDB: params.MariaDB,
		redis:   params.Redis,
	}
}
