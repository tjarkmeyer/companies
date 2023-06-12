package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

// Migration - database auto migration
func Migration(db *gorm.DB) error {
	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")
	return db.AutoMigrate(Industry{}, Company{})
}

// Company - the company DB model
type Company struct {
	ID         string    `json:"id" gorm:"<-:create"`
	Name       string    `json:"name" gorm:"not null;size:256"`
	Address    Address   `json:"address" gorm:"type:jsonb"`
	IndustryID uuid.UUID `json:"industryId" gorm:"type:uuid"`
	Industry   Industry  `json:"industry" gorm:"foreignKey:IndustryID;references:id;"`
	CreatedAt  int64     `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt  int64     `json:"updatedAt" gorm:"autoUpdateTime"`
}

// Address - the address model
type Address struct {
	Street   string `json:"street"`
	Postcode string `json:"postcode"`
	City     string `json:"city"`
}

func (a *Address) Value() (v driver.Value, err error) {
	return json.Marshal(a)
}

func (a *Address) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("invalid data")
	}

	return json.Unmarshal(b, &a)
}

// Industry - the industry DB model
type Industry struct {
	ID           uuid.UUID `json:"id" gorm:"<-:create;type:uuid;default:uuid_generate_v4()"`
	Name         string    `json:"name" gorm:"not null; size:256"`
	MarketValue  float64   `json:"marketValue"`
	Co2Footprint float64   `json:"co2Footprint"`
	CreatedAt    int64     `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt    int64     `json:"updatedAt" gorm:"autoUpdateTime"`
}
