package service

import (
	"context"

	"github.com/nmluci/stellar-payment-lite/internal/indto"
	"github.com/nmluci/stellar-payment-lite/internal/repository"
	"github.com/nmluci/stellar-payment-lite/pkg/dto"
	"github.com/rs/zerolog"
)

type Service interface {
	Ping() (pingResponse dto.PublicPingResponse)

	// Auth
	FindUserByAccessToken(ctx context.Context, at string) (res *indto.UserRole, err error)
	AuthLogin(ctx context.Context, payload *dto.AuthLoginPayload) (res *dto.AuthResponse, err error)
	AuthRefreshToken(ctx context.Context, payload *dto.AuthRefreshTokenPayload) (res *dto.AuthResponse, err error)

	// Users
	RegisterUser(ctx context.Context, payload *dto.UserRegistrationPayload) (err error)
	GetUserDetailByID(ctx context.Context, params *dto.UserQueryParams) (res *dto.UserResponse, err error)
	UpdateUserByID(ctx context.Context, params *dto.UserQueryParams, payload *dto.UserPayload) (err error)

	// Customer
	GetCustomerByID(ctx context.Context, param *dto.CustomerQueryParams) (res *dto.CustomerResponse, err error)
	UpdateCustomer(ctx context.Context, param *dto.CustomerQueryParams, payload *dto.CustomerPayload) (err error)
}

type service struct {
	logger     zerolog.Logger
	conf       *serviceConfig
	repository repository.Repository
}

type serviceConfig struct {
}

type NewServiceParams struct {
	Logger     zerolog.Logger
	Repository repository.Repository
}

func NewService(params *NewServiceParams) Service {
	return &service{
		logger:     params.Logger,
		conf:       &serviceConfig{},
		repository: params.Repository,
	}
}
