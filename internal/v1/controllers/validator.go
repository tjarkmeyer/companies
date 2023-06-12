package controllers

import (
	"unicode/utf8"

	"github.com/tjarkmeyer/companies/companies/internal/v1/dtos"
	"github.com/tjarkmeyer/companies/companies/pkg/errors"
)

type CompanyValidationParams struct {
	CompanyNameLen  int
	IndustryNameLen int
	Co2FootprintMin float64
}

type IValidator interface {
	ValidateCompanyIn(req *dtos.CompanyIn) (err error)
}

type Validator struct {
	validationParams CompanyValidationParams
}

func NewValidator() *Validator {
	return &Validator{
		validationParams: CompanyValidationParams{
			CompanyNameLen:  256,
			IndustryNameLen: 256,
			Co2FootprintMin: 0,
		},
	}
}

func (v *Validator) ValidateCompanyIn(req *dtos.CompanyIn) (err error) {
	if req.Name == nil || utf8.RuneCountInString(*req.Name) > v.validationParams.CompanyNameLen {
		return errors.ErrInvalidCompanyName
	}

	if req.Industry.Name == nil || utf8.RuneCountInString(*req.Industry.Name) > v.validationParams.IndustryNameLen {
		return errors.ErrInvalidIndustryName
	}

	if req.Industry.Co2Footprint != nil && *req.Industry.Co2Footprint < v.validationParams.Co2FootprintMin {
		return errors.ErrInvalidCo2Footprint
	}

	return nil
}
