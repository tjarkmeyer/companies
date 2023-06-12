package configs

import (
	"github.com/tjarkmeyer/golang-toolkit/config/v1"
	"github.com/tjarkmeyer/golang-toolkit/database/v1"
)

// DatabaseConfig - stores DB configuration
var DatabaseConfig database.DBConfig

// DataConnectionConfig - stores DB connection configuration
var DataConnectionConfig database.DataConnectionConf

// LoadPostgresConfig - loads the postgres DB (form environment variables)
func LoadPostgresConfig() {
	var cDbConf database.DBConfig
	var cData database.DataConnectionConf

	config.Process(&cDbConf)
	config.Process(&cData)

	DatabaseConfig = cDbConf
	DataConnectionConfig = cData
}
