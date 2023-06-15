package router

import (
	"github.com/labstack/echo/v4"
	ecMiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/nmluci/go-backend/cmd/webservice/handler"
	"github.com/nmluci/go-backend/internal/config"
	"github.com/nmluci/go-backend/internal/middleware"
	"github.com/nmluci/go-backend/internal/service"
	"github.com/rs/zerolog"
)

type InitRouterParams struct {
	Logger  zerolog.Logger
	Service service.Service
	Ec      *echo.Echo
	Conf    *config.Config
}

func Init(params *InitRouterParams) {
	params.Ec.Use(ecMiddleware.CORS(), middleware.ServiceVersioner)
	params.Ec.GET(PingPath, handler.HandlePing(params.Service.Ping))
}
