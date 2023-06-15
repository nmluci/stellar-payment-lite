package webservice

import (
	"sync"

	"github.com/labstack/echo/v4"
	"github.com/nmluci/stellar-payment-lite/cmd/webservice/router"
	"github.com/nmluci/stellar-payment-lite/internal/component"
	"github.com/nmluci/stellar-payment-lite/internal/config"
	"github.com/nmluci/stellar-payment-lite/internal/repository"
	"github.com/nmluci/stellar-payment-lite/internal/service"
	"github.com/rs/zerolog"
)

func Start(conf *config.Config, logger zerolog.Logger) {
	db, err := component.InitMariaDB(&component.InitMariaDBParams{
		Conf:   &conf.MariaDBConfig,
		Logger: logger,
	})

	if err != nil {
		logger.Fatal().Msgf("failed to initialize maria db: %+v", err)
	}

	// redis, err := component.InitRedis(&component.InitRedisParams{
	// 	Conf:   &conf.RedisConfig,
	// 	Logger: logger,
	// })

	// if err != nil {
	// 	logger.Fatalf("%s initalizing redis: %+v", logTagStartWebservice, err)
	// }

	ec := echo.New()
	ec.HideBanner = true
	ec.HidePort = true

	repo := repository.NewRepository(&repository.NewRepositoryParams{
		Logger:  logger,
		MariaDB: db,
		// Redis: redis,
	})

	service := service.NewService(&service.NewServiceParams{
		Logger:     logger,
		Repository: repo,
	})

	router.Init(&router.InitRouterParams{
		Logger:  logger,
		Service: service,
		Ec:      ec,
		Conf:    conf,
	})

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		logger.Info().Msgf("starting service, listening to: %s", conf.ServiceAddress)

		if err := ec.Start(conf.ServiceAddress); err != nil {
			logger.Error().Msgf("starting service, cause: %+v", err)
		}
	}()

	wg.Wait()
}
