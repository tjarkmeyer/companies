package internal

import (
	"github.com/tjarkmeyer/companies/companies/internal/v1/dtos"
	"github.com/tjarkmeyer/companies/companies/internal/v1/models"
)

// ICompaniesRepository - defines the companies repository
type ICompaniesRepository interface {
	Create(*models.Company) error
	Update(*models.Company) error
	GetByID(string) (*models.Company, error)
	DeleteByID(string) error
}

// ICompaniesService - defines the companies service
type ICompaniesService interface {
	Create(*dtos.CompanyIn) error
	Update(*dtos.CompanyIn) error
	GetByID(string) (*models.Company, error)
	DeleteByID(string) error
}
