package tests

import (
	"net/http"
	"net/http/httptest"

	sentryhttp "github.com/getsentry/sentry-go/http"
	"github.com/tjarkmeyer/companies/companies/internal/v1"
	"github.com/tjarkmeyer/companies/companies/internal/v1/controllers"
	"github.com/tjarkmeyer/companies/companies/internal/v1/services"
	"github.com/tjarkmeyer/companies/companies/pkg/error_adapters/httpadapter"
	logger "github.com/tjarkmeyer/golang-toolkit/logger/v1"
	"github.com/tjarkmeyer/golang-toolkit/servers/rest"
)

const jsonContentType = "application/json; charset=utf-8"

func prepareServer(repositroy internal.ICompaniesRepository) *httptest.Server {
	restController := rest.NewController()
	logger := logger.New("testing", "")
	adapter := httpadapter.New(http.StatusInternalServerError, httpadapter.AdaptBadRequestError)
	service := services.NewCompaniesService(repositroy, logger)

	handler := controllers.NewCompaniesHandler(service, logger, &sentryhttp.Handler{}, adapter)
	controllers.NewCompaniesAPIRouter(handler, restController)

	return httptest.NewServer(restController.CreateControllerByName())
}
