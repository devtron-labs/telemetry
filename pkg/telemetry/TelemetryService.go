package telemetry

import (
	"github.com/devtron-labs/telemetry/internal/util"
	"go.uber.org/zap"
	"net/http"
	"strings"
)

type TelemetryService interface {
	GetByAPIKey() (string, error)
	CheckWhitelist(ucid string) (bool, error)
}

const devtronUCIDPrefix = "devtron"

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

func (impl *TelemetryServiceImpl) CheckWhitelist(ucid string) (bool, error) {
	//todo - whitelisted ids
	var whitelistedUcids []string
	if strings.Contains(ucid, devtronUCIDPrefix) {
		return true, nil
	} else if contains(whitelistedUcids, ucid) {
		return true, nil
	}
	return false, nil
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
