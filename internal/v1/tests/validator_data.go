package tests

import (
	"github.com/gofrs/uuid"
	"github.com/tjarkmeyer/companies/companies/internal/v1/dtos"
	"github.com/tjarkmeyer/golang-toolkit/utils"
)

const string260Char = "Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata san."

var (
	CompanyInOk = &dtos.CompanyIn{
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
)
