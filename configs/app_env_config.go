package configs

import (
	"github.com/tjarkmeyer/golang-toolkit/config/v1"
)

// AppEnvConfig - to store global app configuration
var AppEnvConfig AppEnvironmentConfig

// AppEnvironmentConfig - the global app configuration
type AppEnvironmentConfig struct {
	AppName     string `default:"companies" envconfig:"APP_NAME"`
	Environment string `default:"development" envconfig:"APP_ENV"`
	SentryDSN   string `default:"" envconfig:"SENTRY_DSN"`
}

// LoadAppEnvConfig - loads the app configuration (form environment variables)
func LoadAppEnvConfig() {
	var appEnvConfig AppEnvironmentConfig
	config.Process(&appEnvConfig)
	AppEnvConfig = appEnvConfig
}
