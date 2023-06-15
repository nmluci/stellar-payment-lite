package service

import (
	"context"

	"github.com/nmluci/go-backend/internal/repository"
	"github.com/nmluci/go-backend/internal/worker"
	"github.com/nmluci/go-backend/pkg/dto"
	"github.com/nmluci/gostellar"
	"github.com/rs/zerolog"
)

type Service interface {
	Ping() (pingResponse dto.PublicPingResponse)
	AuthenticateSession(ctx context.Context, token string) (access context.Context, err error)
	AuthenticateService(ctx context.Context, name string) (access context.Context, err error)
}

type service struct {
	logger     zerolog.Logger
	conf       *serviceConfig
	repository repository.Repository
	stellarRPC *gostellar.StellarRPC
	goStellar  *gostellar.GoStellar
	fileWorker *worker.WorkerManager
}

type serviceConfig struct {
}

type NewServiceParams struct {
	Logger     zerolog.Logger
	Repository repository.Repository
	StellarRPC *gostellar.StellarRPC
	GoStellar  *gostellar.GoStellar
	FileWorker *worker.WorkerManager
}

func NewService(params *NewServiceParams) Service {
	return &service{
		logger:     params.Logger,
		conf:       &serviceConfig{},
		repository: params.Repository,
		stellarRPC: params.StellarRPC,
		goStellar:  params.GoStellar,
		fileWorker: params.FileWorker,
	}
}
