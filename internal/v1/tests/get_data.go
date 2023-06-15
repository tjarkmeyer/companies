package tests

import (
	"github.com/tjarkmeyer/companies/companies/internal/v1/models"
)

const defaultCompanyID = "my_company"

var (
	GetCompanyRepo = models.Company{
		ID:   defaultCompanyID,
		Name: *CreateCompanyRequestOk.Name,
		Address: models.Address{
			Street:   *CreateCompanyRequestOk.Address.Street,
			Postcode: *CreateCompanyRequestOk.Address.Postcode,
			City:     *CreateCompanyRequestOk.Address.City,
		},
		Industry: models.Industry{
			ID:           *CreateCompanyRequestOk.Industry.ID,
			Name:         *CreateCompanyRequestOk.Industry.Name,
			MarketValue:  *CreateCompanyRequestOk.Industry.MarketValue,
			Co2Footprint: *CreateCompanyRequestOk.Industry.Co2Footprint,
			CreatedAt:    DefaultTime.Unix(),
			UpdatedAt:    DefaultTime.Unix(),
		},
		CreatedAt: DefaultTime.Unix(),
		UpdatedAt: DefaultTime.Unix(),
	}
)
