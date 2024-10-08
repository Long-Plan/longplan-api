package repo

import (
	"github.com/Long-Plan/longplan-api/internal/core/model"
	"github.com/Long-Plan/longplan-api/internal/core/port"
	"gorm.io/gorm"
)

type sysCategoryTypeRepo struct {
	db *gorm.DB
}

func NewSysCategoryTypeRepo(db *gorm.DB) port.SysCategoryTypeRepo {
	return &sysCategoryTypeRepo{db}
}

func (r *sysCategoryTypeRepo) GetAll() ([]model.SysCategoryType, error) {
	var categoryTypes []model.SysCategoryType
	if err := r.db.Find(&categoryTypes).Error; err != nil {
		return nil, err
	}
	return categoryTypes, nil
}

func (r *sysCategoryTypeRepo) Create(categoryType *model.SysCategoryType) error {
	return r.db.Create(categoryType).Error
}

func (r *sysCategoryTypeRepo) Update(categoryType *model.SysCategoryType) error {
	return r.db.Updates(categoryType).Error
}

func (r *sysCategoryTypeRepo) Delete(categoryTypeID int) error {
	return r.db.Delete(&model.SysCategoryType{}, categoryTypeID).Error
}
