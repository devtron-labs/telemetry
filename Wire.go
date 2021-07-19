//+build wireinject

package main

import (
	"github.com/devtron-labs/telemetry/api"
	"github.com/devtron-labs/telemetry/internal/logger"
	"github.com/devtron-labs/telemetry/internal/util"
	"github.com/devtron-labs/telemetry/pkg/telemetry"
	"github.com/google/wire"
)

func InitializeApp() (*App, error) {
	wire.Build(
		NewApp,
		api.NewMuxRouter,
		logger.NewSugardLogger,
		logger.NewHttpClient,
		api.NewRestHandlerImpl,
		wire.Bind(new(api.RestHandler), new(*api.RestHandlerImpl)),
		telemetry.NewTelemetryServiceImpl,
		wire.Bind(new(telemetry.TelemetryService), new(*telemetry.TelemetryServiceImpl)),
		util.GetPosthogConfig,
		telemetry.GetOptOutConfig,
	)
	return &App{}, nil
}
