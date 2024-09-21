package repo

import (
	"github.com/Long-Plan/longplan-api/internal/core/model"
	"github.com/Long-Plan/longplan-api/internal/core/port"
	"gorm.io/gorm"
)

type organizationRepo struct {
	db *gorm.DB
}

func NewOrganizationRepo(db *gorm.DB) port.OrganizationRepo {
	return &organizationRepo{db}
}

func (r *organizationRepo) GetAll() ([]model.Organization, error) {
	var organizations []model.Organization
	if err := r.db.Find(&organizations).Error; err != nil {
		return nil, err
	}
	return organizations, nil
}

func (r *organizationRepo) Create(organization *model.Organization) error {
	return r.db.Create(organization).Error
}

func (r *organizationRepo) Update(organization *model.Organization) error {
	return r.db.Updates(organization).Error
}

func (r *organizationRepo) Delete(organizationID int) error {
	return r.db.Delete(&model.Organization{}, organizationID).Error
}
