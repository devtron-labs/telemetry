package telemetry

import (
	"github.com/caarlos0/env"
	"github.com/devtron-labs/telemetry/internal/util"
	"go.uber.org/zap"
	"net/http"
	"strings"
)

type TelemetryService interface {
	GetByAPIKey() (string, error)
	CheckForOptOut(ucid string) (bool, error)
}

type OptOutConfig struct {
	OptOutClients string `env:"OPT_OUT_CLIENTS" envDefault:""`
}

func GetOptOutConfig() (*OptOutConfig, error) {
	cfg := &OptOutConfig{}
	err := env.Parse(cfg)
	if err != nil {
		return nil, err
	}
	return cfg, err
}

type TelemetryServiceImpl struct {
	logger        *zap.SugaredLogger
	client        *http.Client
	posthogConfig *util.PosthogConfig
	optOutConfig  *OptOutConfig
}

func NewTelemetryServiceImpl(logger *zap.SugaredLogger, client *http.Client, posthogConfig *util.PosthogConfig,
	optOutConfig *OptOutConfig) *TelemetryServiceImpl {
	serviceImpl := &TelemetryServiceImpl{
		logger:        logger,
		client:        client,
		posthogConfig: posthogConfig,
		optOutConfig:  optOutConfig,
	}
	return serviceImpl
}

func (impl *TelemetryServiceImpl) GetByAPIKey() (string, error) {
	apiKey := impl.posthogConfig.PosthogApiKey
	return apiKey, nil
}

func (impl *TelemetryServiceImpl) CheckForOptOut(ucid string) (bool, error) {
	isOptOut := false
	optOutClients := impl.optOutConfig.OptOutClients
	optOutClientsArray := strings.Split(optOutClients, ",")
	for _, optOutClientId := range optOutClientsArray {
		if ucid == optOutClientId {
			isOptOut = true
			return isOptOut, nil
		}
	}
	return isOptOut, nil
}
