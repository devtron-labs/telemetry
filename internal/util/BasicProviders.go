package util

import (
	"github.com/caarlos0/env"
	"go.uber.org/zap"
	"net/http"
)

var (
	// Logger is the defaut logger
	logger *zap.SugaredLogger
	//FIXME: remove this
	//defer Logger.Sync()
)

// Deprecated: instead calling this method inject logger from wire
func GetLogger() *zap.SugaredLogger {
	return logger
}

type PosthogConfig struct {
	PosthogApiKey     string `env:"POSTHOG_API_KEY" envDefault:""`
	PosthogEndpoint   string `env:"POSTHOG_ENDPOINT" envDefault:"https://app.posthog.com"`
}

func GetPosthogConfig() (*PosthogConfig, error) {
	cfg := &PosthogConfig{}
	err := env.Parse(cfg)
	if err != nil {
		return nil, err
	}
	return cfg, err
}

func init() {
	_, err := zap.NewProduction()
	if err != nil {
		panic("failed to create the default logger: " + err.Error())
	}
}

func NewSugardLogger() *zap.SugaredLogger {
	return logger
}

func NewHttpClient() *http.Client {
	return http.DefaultClient
}
