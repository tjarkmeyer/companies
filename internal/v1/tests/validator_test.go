package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tjarkmeyer/companies/companies/internal/v1/controllers"
	"github.com/tjarkmeyer/companies/companies/pkg/errors"
	"github.com/tjarkmeyer/golang-toolkit/utils"
)

func Test_Validator(t *testing.T) {
	validator := controllers.NewValidator()

	assert.Equal(t, validator.ValidateCompanyIn(CompanyInOk), nil)
	// invalid company name
	invalidCompanyName := *CompanyInOk
	invalidCompanyName.Name = utils.P[string](string260Char)
	assert.Equal(t, validator.ValidateCompanyIn(&invalidCompanyName), errors.ErrInvalidCompanyName)
	// invalid industry name
	invalidIndustryName := *CompanyInOk
	invalidIndustryName.Industry.Name = utils.P[string](string260Char)
	assert.Equal(t, validator.ValidateCompanyIn(&invalidIndustryName), errors.ErrInvalidIndustryName)
	// invalid co2 footprint
	invalidIndustryFootprint := *CompanyInOk
	invalidIndustryFootprint.Industry.Co2Footprint = utils.P[float64](float64(-2))
	assert.Equal(t, validator.ValidateCompanyIn(&invalidIndustryFootprint), errors.ErrInvalidCo2Footprint)
}
