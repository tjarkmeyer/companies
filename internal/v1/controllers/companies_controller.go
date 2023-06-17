package controllers

import (
	"encoding/json"
	"net/http"

	sentryhttp "github.com/getsentry/sentry-go/http"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/tjarkmeyer/companies/companies/internal/v1"
	"github.com/tjarkmeyer/companies/companies/internal/v1/dtos"
	"github.com/tjarkmeyer/companies/companies/pkg/error_adapters/httpadapter"
	"github.com/tjarkmeyer/golang-toolkit/httpencoder"
	"github.com/tjarkmeyer/golang-toolkit/servers/rest"
	"go.uber.org/zap"
)

// CompaniesHandler - defines the companies handler
type CompaniesHandler struct {
	service      internal.ICompaniesService
	log          *zap.Logger
	tracing      *sentryhttp.Handler
	errorAdapter httpadapter.IErrorAdapter
	validator    IValidator
	encoder      httpencoder.IHttpEncoder
}

// NewCompaniesHandler - new companies handler
func NewCompaniesHandler(service internal.ICompaniesService, log *zap.Logger, tracing *sentryhttp.Handler, errAdapater httpadapter.IErrorAdapter) *CompaniesHandler {
	return &CompaniesHandler{
		service:      service,
		log:          log,
		tracing:      tracing,
		errorAdapter: errAdapater,
		validator:    NewValidator(),
		encoder:      httpencoder.New(),
	}
}

// NewCompaniesAPIRouter - creates a rest new companies API router
func NewCompaniesAPIRouter(h *CompaniesHandler, rd *rest.Definitions) {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(h.tracing.Handle)

	r.With(MiddlewareRoleCheck).Post("/", h.create)
	r.With(MiddlewareRoleCheck).Put("/", h.update)
	r.Get("/{companyID}", h.getByID)
	r.With(MiddlewareRoleCheck).Delete("/{companyID}", h.deleteByID)

	controllerDefinition := &rest.Definition{Controller: r, Name: "companies"}

	rd.AddController(controllerDefinition)
}

func (h *CompaniesHandler) create(w http.ResponseWriter, req *http.Request) {
	companyIn := &dtos.CompanyIn{}
	if err := json.NewDecoder(req.Body).Decode(companyIn); err != nil {
		h.log.Error("[ERROR] Decode http request body", zap.Error(err))
		h.encoder.EncodeFailed(w, http.StatusBadRequest)
		return
	}

	if err := h.validator.ValidateCompanyIn(companyIn); err != nil {
		h.log.Error("[ERROR] Validate request vody", zap.Error(err))
		h.encoder.EncodeFailed(w, h.errorAdapter.AdaptToHttpCode(err))
		return
	}

	h.log.Debug("[START] Create /companies", zap.Any("companyIn", companyIn))

	if err := h.service.Create(companyIn); err != nil {
		h.log.Error("[ERROR] Something bad happend while working on the request", zap.Error(err))
		h.encoder.EncodeFailed(w, h.errorAdapter.AdaptToHttpCode(err))
		return
	}

	h.log.Debug("[DONE] Create /companies", zap.Any("companyIn", companyIn))

	h.encoder.EncodeSuccesful(w, http.StatusCreated)
}

func (h *CompaniesHandler) update(w http.ResponseWriter, req *http.Request) {
	companyIn := &dtos.CompanyIn{}
	if err := json.NewDecoder(req.Body).Decode(companyIn); err != nil {
		h.log.Error("[ERROR] Decode http request body", zap.Error(err))
		h.encoder.EncodeFailed(w, http.StatusBadRequest)
		return
	}

	if err := h.validator.ValidateCompanyIn(companyIn); err != nil {
		h.log.Error("[ERROR] Validate request vody", zap.Error(err))
		h.encoder.EncodeFailed(w, h.errorAdapter.AdaptToHttpCode(err))
		return
	}

	h.log.Debug("[START] Update /companies", zap.Any("companyIn", companyIn))

	if err := h.service.Update(companyIn); err != nil {
		h.log.Error("[ERROR] Something bad happend while working on the request", zap.Error(err))
		h.encoder.EncodeFailed(w, h.errorAdapter.AdaptToHttpCode(err))
		return
	}

	h.log.Debug("[DONE] Update /companies", zap.Any("companyIn", companyIn))

	h.encoder.EncodeSuccesful(w, http.StatusOK)
}

func (h *CompaniesHandler) getByID(w http.ResponseWriter, req *http.Request) {
	companyID := chi.URLParam(req, "companyID")

	h.log.Debug("[START] Get /companies/{companyID}", zap.String("companyID", companyID))

	result, err := h.service.GetByID(companyID)

	if err != nil {
		h.log.Error("[ERROR] Something bad happend while working on the request", zap.Error(err))
		h.encoder.EncodeFailed(w, h.errorAdapter.AdaptToHttpCode(err))
		return
	}

	h.log.Debug("[DONE] Get /companies/{companyID}", zap.String("companyID", companyID))

	h.encoder.EncodeJson(result, w, http.StatusOK)
}

func (h *CompaniesHandler) deleteByID(w http.ResponseWriter, req *http.Request) {
	companyID := chi.URLParam(req, "companyID")

	h.log.Debug("[START] DELETE /companies/{companyID}", zap.String("companyID", companyID))

	if err := h.service.DeleteByID(companyID); err != nil {
		h.log.Error("[ERROR] Something bad happend while working on the request", zap.Error(err))
		h.encoder.EncodeFailed(w, h.errorAdapter.AdaptToHttpCode(err))
		return
	}

	h.log.Debug("[DONE] DELETE /companies/{companyID}", zap.String("companyID", companyID))

	h.encoder.EncodeSuccesful(w, http.StatusOK)
}
