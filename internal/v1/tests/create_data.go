package tests

import (
	"time"

	"github.com/gofrs/uuid"
	"github.com/tjarkmeyer/companies/companies/internal/v1/dtos"
	"github.com/tjarkmeyer/companies/companies/internal/v1/models"
	"github.com/tjarkmeyer/golang-toolkit/utils"
)

const defaultUUIDString = "4f6b2a36-7429-4855-89a8-84cb64b76fef"

var (
	defaultUUID = uuid.FromStringOrNil(defaultUUIDString)
	DefaultTime = time.Date(2022, 1, 1, 1, 1, 1, 1, time.UTC)

	CreateCompanyRequestOk = &dtos.CompanyIn{
		Name: utils.P[string]("Company Name"),
		Address: dtos.AddressIn{
			Street:   utils.P[string]("Example street 1"),
			Postcode: utils.P[string]("12345"),
			City:     utils.P[string]("City"),
		},
		Industry: dtos.IndustryIn{
			ID:           utils.P[uuid.UUID](defaultUUID),
			Name:         utils.P[string]("retail"),
			MarketValue:  utils.P[float64](float64(1000)),
			Co2Footprint: utils.P[float64](float64(5000)),
		},
	}

	CreateCompanyRepo = models.Company{
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
		},
	}
)
