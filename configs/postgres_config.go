package configs

import (
	"github.com/tjarkmeyer/golang-toolkit/config/v1"
	"github.com/tjarkmeyer/golang-toolkit/database/v1"
)

// DatabaseConfig - stores DB configuration
var DatabaseConfig database.Config

// DataConnectionConfig - stores DB connection configuration
var DataConnectionConfig database.ConnectionConfig

// LoadPostgres - loads the postgres DB (form environment variables)
func LoadPostgres() {
	var cDbConf database.Config
	var cData database.ConnectionConfig

	config.Process(&cDbConf)
	config.Process(&cData)

	DatabaseConfig = cDbConf
	DataConnectionConfig = cData
}
