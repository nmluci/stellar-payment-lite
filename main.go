package main

import (
	"github.com/nmluci/stellar-payment-lite/cmd/webservice"
	"github.com/nmluci/stellar-payment-lite/internal/component"
	"github.com/nmluci/stellar-payment-lite/internal/config"
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
