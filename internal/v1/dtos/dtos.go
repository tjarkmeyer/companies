package dtos

import (
	"github.com/gofrs/uuid"
)

// CompanyIn - the incoming company
type CompanyIn struct {
	ID       *string    `json:"id"`
	Name     *string    `json:"name"`
	Address  AddressIn  `json:"address"`
	Industry IndustryIn `json:"industry"`
}

// AddressIn - the incoming address
type AddressIn struct {
	Street   *string `json:"street"`
	Postcode *string `json:"postcode"`
	City     *string `json:"city"`
}

// IndustryIn - the incoming industry
type IndustryIn struct {
	ID           *uuid.UUID `json:"id" gorm:"<-:create"`
	Name         *string    `json:"name"`
	MarketValue  *float64   `json:"marketValue"`
	Co2Footprint *float64   `json:"co2Footprint"`
}
