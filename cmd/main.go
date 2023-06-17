package main

import (
	"net/http"

	"github.com/tjarkmeyer/companies/companies/configs"
	"github.com/tjarkmeyer/companies/companies/internal/v1/controllers"
	"github.com/tjarkmeyer/companies/companies/internal/v1/models"
	"github.com/tjarkmeyer/companies/companies/internal/v1/repositories"
	"github.com/tjarkmeyer/companies/companies/internal/v1/services"
	"github.com/tjarkmeyer/companies/companies/pkg/error_adapters/httpadapter"
	"github.com/tjarkmeyer/companies/companies/pkg/error_adapters/sqladapter"
	"github.com/tjarkmeyer/golang-toolkit/database/v1"
	logger "github.com/tjarkmeyer/golang-toolkit/logger/v1"
	"github.com/tjarkmeyer/golang-toolkit/servers/rest"
	tracing "github.com/tjarkmeyer/golang-toolkit/tracing"
)

func init() {
	configs.LoadAppEnv()
	configs.LoadPostgres()
}

func main() {
	logger := logger.New(configs.AppEnvConfig.Environment, configs.AppEnvConfig.SentryDSN)

	db := database.Connect(configs.DataConnectionConfig, configs.DatabaseConfig)
	err := models.Migration(db)
	if err != nil {
		logger.Info(err.Error())
		logger.Panic("DB migration failed")
	}

	tracer, err := tracing.New(configs.AppEnvConfig.SentryDSN, configs.AppEnvConfig.Environment, configs.AppEnvConfig.AppName)
	if err != nil {
		logger.Info(err.Error())
		logger.Panic("Init tracing failed")
	}

	sqlErrAdapter := sqladapter.New(repositories.ErrorsMap)
	httpErrAdapter := httpadapter.New(http.StatusInternalServerError, httpadapter.AdaptBadRequestError)

	restController := rest.NewController()

	companiesRepository := repositories.NewCompaniesRepository(db, sqlErrAdapter)
	companiesService := services.NewCompaniesService(companiesRepository, logger)
	companiesHandler := controllers.NewCompaniesHandler(companiesService, logger, tracer, httpErrAdapter)
	controllers.NewCompaniesAPIRouter(companiesHandler, restController)

	apiController := restController.CreateControllerByName()

	logger.Info("Server started")

	err = http.ListenAndServe(":8080", apiController)
	if err != nil {
		logger.Info(err.Error())
		logger.Panic("Server crashed")
	}
}
