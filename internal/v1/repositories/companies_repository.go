package repositories

import (
	"github.com/tjarkmeyer/companies/companies/internal/v1/models"
	"github.com/tjarkmeyer/companies/companies/pkg/error_adapters/sql_adapter"
	"gorm.io/gorm"
)

// CompaniesRepository - defines companies repository
type CompaniesRepository struct {
	db           *gorm.DB
	errorAdapter sql_adapter.IErrorAdapter
}

// NewCompaniesRepository - creates new companies repository
func NewCompaniesRepository(db *gorm.DB, errorAdapter sql_adapter.IErrorAdapter) *CompaniesRepository {
	return &CompaniesRepository{
		db:           db,
		errorAdapter: errorAdapter,
	}
}

// Create - creates a new company in DB
func (r *CompaniesRepository) Create(company *models.Company) (err error) {
	if err = r.db.Transaction(func(tx *gorm.DB) (err error) {
		if err = r.upsertIndustry(tx, &company.Industry); err != nil {
			return
		}
		company.IndustryID = company.Industry.ID

		if err = r.createCompany(tx, company); err != nil {
			return
		}

		return
	}); err != nil {
		return r.errorAdapter.AdaptSqlErr(err)
	}

	return nil
}

// Update - updates a company
func (r *CompaniesRepository) Update(company *models.Company) (err error) {
	if err = r.db.Transaction(func(tx *gorm.DB) (err error) {
		if err = r.upsertIndustry(tx, &company.Industry); err != nil {
			return
		}

		if err = r.updateCompany(tx, company); err != nil {
			return
		}
		return
	}); err != nil {
		return r.errorAdapter.AdaptSqlErr(err)
	}
	return nil
}

// GetByID - returns company by ID from DB
func (r *CompaniesRepository) GetByID(id string) (*models.Company, error) {
	var company *models.Company
	err := r.db.Preload("Industry").First(&company, "id = ?", id).Error
	if company.ID == "" {
		return nil, ErrNotFound
	}
	return company, r.errorAdapter.AdaptSqlErr(err)
}

// DeleteByID - deletes a company by ID in DB
func (r *CompaniesRepository) DeleteByID(id string) error {
	res := r.db.Delete(&models.Company{ID: id})
	if res.Error != nil {
		return r.errorAdapter.AdaptSqlErr(res.Error)
	}
	if res.RowsAffected == 0 {
		return ErrNotFound
	}
	return nil
}

func (r *CompaniesRepository) upsertIndustry(tx *gorm.DB, industry *models.Industry) error {
	if industry.ID.String() != "" {
		res := tx.Model(industry).Where("id = ?", industry.ID).Updates(industry)
		if res.RowsAffected != 0 {
			return res.Error
		}
	}

	return tx.Create(industry).Error
}

func (r *CompaniesRepository) createCompany(tx *gorm.DB, company *models.Company) (err error) {
	return tx.Create(company).Error
}

func (r *CompaniesRepository) updateCompany(tx *gorm.DB, company *models.Company) (err error) {
	res := tx.Model(&company).Where("id = ?", company.ID).Updates(&company)
	if err = res.Error; err != nil {
		return err
	}

	if res.RowsAffected == 0 {
		return ErrNotFound
	}

	return nil
}
