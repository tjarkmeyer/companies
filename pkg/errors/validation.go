package errors

import "errors"

var (
	ErrInvalidCompanyName  = errors.New("invalid company name")
	ErrInvalidIndustryName = errors.New("invalid industry name")
	ErrInvalidCo2Footprint = errors.New("invalid co2 footprint")
	ErrInvalidUUID         = errors.New("invalid uuid")
)
