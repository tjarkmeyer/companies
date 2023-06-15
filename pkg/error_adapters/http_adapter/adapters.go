package http_adapter

import (
	"net/http"

	"github.com/tjarkmeyer/companies/companies/internal/v1/repositories"
	"github.com/tjarkmeyer/companies/companies/pkg/errors"
)

type ErrorToHttpCodeAdapter func(err error) (code int)

func AdaptNotFoundError(err error) (code int) {
	if err == repositories.ErrNotFound {
		return http.StatusNotFound
	}

	return code
}

func AdaptBadRequestError(err error) (code int) {
	switch err {
	case errors.ErrInvalidCompanyName,
		errors.ErrInvalidIndustryName,
		errors.ErrInvalidCo2Footprint,
		errors.ErrInvalidUUID:
		return http.StatusBadRequest
	case repositories.ErrAlreadyExist:
		return http.StatusConflict
	case repositories.ErrNotFound:
		return http.StatusNotFound
	default:
		return code
	}
}
