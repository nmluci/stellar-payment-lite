package main

import (
	"github.com/nmluci/go-backend/cmd/webservice"
	"github.com/nmluci/go-backend/internal/component"
	"github.com/nmluci/go-backend/internal/config"
)

var (
	buildVer  string = "unknown"
	buildTime string = "unknown"
)

func main() {
	config.Init(buildTime, buildVer)
	conf := config.Get()

	logger := component.NewLogger(component.NewLoggerParams{
		ServiceName: conf.ServiceName,
		PrettyPrint: true,
	})

	webservice.Start(conf, logger)
}