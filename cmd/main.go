package main

import (
	"net/http"

	"github.com/tjarkmeyer/companies/companies/configs"
	"github.com/tjarkmeyer/companies/companies/internal/v1/controllers"
	"github.com/tjarkmeyer/companies/companies/internal/v1/models"
	"github.com/tjarkmeyer/companies/companies/internal/v1/repositories"
	"github.com/tjarkmeyer/companies/companies/internal/v1/services"
	"github.com/tjarkmeyer/companies/companies/pkg/error_adapters/http_adapter"
	"github.com/tjarkmeyer/companies/companies/pkg/error_adapters/sql_adapter"
	"github.com/tjarkmeyer/golang-toolkit/database/v1"
	logger "github.com/tjarkmeyer/golang-toolkit/logger/sentry"
	tracing "github.com/tjarkmeyer/golang-toolkit/sentry-tracing"
	"github.com/tjarkmeyer/golang-toolkit/servers/rest"
)

func init() {
	configs.LoadAppEnvConfig()
	configs.LoadPostgresConfig()
}

func main() {
	logger := logger.InitLogger(configs.AppEnvConfig.Environment, configs.AppEnvConfig.SentryDSN)

	db := database.Connect(configs.DataConnectionConfig, configs.DatabaseConfig)
	err := models.Migration(db)
	if err != nil {
		logger.Info(err.Error())
		logger.Panic("DB migration failed")
	}

	tracer, err := tracing.InitSentry(configs.AppEnvConfig.SentryDSN, configs.AppEnvConfig.Environment, configs.AppEnvConfig.AppName)
	if err != nil {
		logger.Info(err.Error())
		logger.Panic("Init tracing failed")
	}

	httpErrAdapater := http_adapter.New(http.StatusInternalServerError, http_adapter.AdaptNotFoundError, http_adapter.AdaptBadRequestError)
	sqlErrAdapater := sql_adapter.New(repositories.ErrorsMap)

	restController := rest.NewRestController()

	companiesRepository := repositories.NewCompaniesRepository(db, sqlErrAdapater)
	companiesService := services.NewCompaniesService(companiesRepository, logger)
	companiesHandler := controllers.NewCompaniesHandler(companiesService, logger, tracer, httpErrAdapater)
	controllers.NewCompaniessAPIRouter(companiesHandler, restController)

	apiController := restController.CreateRestControllerByName()

	logger.Info("Server started")

	err = http.ListenAndServe(":8080", apiController)
	if err != nil {
		logger.Info(err.Error())
		logger.Panic("Server crashed")
	}
}
