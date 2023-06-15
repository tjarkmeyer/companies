package tests

import (
	"net/http"
	"net/http/httptest"

	sentryhttp "github.com/getsentry/sentry-go/http"
	"github.com/tjarkmeyer/companies/companies/internal/v1"
	"github.com/tjarkmeyer/companies/companies/internal/v1/controllers"
	"github.com/tjarkmeyer/companies/companies/internal/v1/services"
	"github.com/tjarkmeyer/companies/companies/pkg/error_adapters/http_adapter"
	logger "github.com/tjarkmeyer/golang-toolkit/logger/sentry"
	"github.com/tjarkmeyer/golang-toolkit/servers/rest"
)

const jsonContentType = "application/json; charset=utf-8"

func prepareServer(repositroy internal.ICompaniesRepository) *httptest.Server {
	restController := rest.NewRestController()
	logger := logger.InitLogger("testing", "")
	adapter := http_adapter.New(http.StatusInternalServerError, http_adapter.AdaptNotFoundError, http_adapter.AdaptBadRequestError)
	service := services.NewCompaniesService(repositroy, logger)

	handler := controllers.NewCompaniesHandler(service, logger, &sentryhttp.Handler{}, adapter)
	controllers.NewCompaniessAPIRouter(handler, restController)

	return httptest.NewServer(restController.CreateRestControllerByName())
}
