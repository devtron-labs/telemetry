package telemetry

import (
	"github.com/devtron-labs/telemetry/internal/util"
	"go.uber.org/zap"
	"net/http"
)

type TelemetryService interface {
	GetByAPIKey() (string, error)
}

type TelemetryServiceImpl struct {
	logger        *zap.SugaredLogger
	client        *http.Client
	posthogConfig *util.PosthogConfig
}

func NewTelemetryServiceImpl(logger *zap.SugaredLogger, client *http.Client, posthogConfig *util.PosthogConfig) *TelemetryServiceImpl {
	serviceImpl := &TelemetryServiceImpl{
		logger:        logger,
		client:        client,
		posthogConfig: posthogConfig,
	}
	return serviceImpl
}

func (impl *TelemetryServiceImpl) GetByAPIKey() (string, error) {
	apiKey := impl.posthogConfig.PosthogApiKey
	return apiKey, nil
}
