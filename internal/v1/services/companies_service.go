package services

import (
	"github.com/jinzhu/copier"
	"github.com/tjarkmeyer/companies/companies/internal/v1"
	"github.com/tjarkmeyer/companies/companies/internal/v1/dtos"
	"github.com/tjarkmeyer/companies/companies/internal/v1/models"
	"go.uber.org/zap"
)

// CompaniesService - defines the companies service
type CompaniesService struct {
	repository internal.ICompaniesRepository
	log        *zap.Logger
}

// NewCompaniesService - creates a new companies service
func NewCompaniesService(repository internal.ICompaniesRepository, log *zap.Logger) *CompaniesService {
	return &CompaniesService{
		repository: repository,
		log:        log,
	}
}

// Create - creates a new company
func (s *CompaniesService) Create(companyIn *dtos.CompanyIn) error {
	company := &models.Company{}
	if err := copier.Copy(company, companyIn); err != nil {
		return err
	}
	return s.repository.Create(company)
}

// Update - updates a company
func (s *CompaniesService) Update(companyIn *dtos.CompanyIn) error {
	company := &models.Company{}
	if err := copier.Copy(company, companyIn); err != nil {
		return err
	}
	return s.repository.Update(company)
}

// GetByID - returns a company by ID
func (s *CompaniesService) GetByID(id string) (*models.Company, error) {
	return s.repository.GetByID(id)
}

// DeleteByID - deletes a company by ID
func (s *CompaniesService) DeleteByID(id string) error {
	return s.repository.DeleteByID(id)
}
