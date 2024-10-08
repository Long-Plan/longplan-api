package repo

import (
	"github.com/Long-Plan/longplan-api/internal/core/model"
	"github.com/Long-Plan/longplan-api/internal/core/port"
	"gorm.io/gorm"
)

type sysCategoryRelationshipRepo struct {
	db *gorm.DB
}

func NewSysCategoryRelationshipRepo(db *gorm.DB) port.SysCategoryRelationshipRepo {
	return &sysCategoryRelationshipRepo{db}
}

func (r *sysCategoryRelationshipRepo) GetByParentCategoryID(parentCategoryID int) ([]model.SysCategoryRelationship, error) {
	var categoryRelationships []model.SysCategoryRelationship
	if err := r.db.Where("parent_category_id = ?", parentCategoryID).Find(&categoryRelationships).Error; err != nil {
		return nil, err
	}
	return categoryRelationships, nil
}

func (r *sysCategoryRelationshipRepo) Create(categoryRelationship *model.SysCategoryRelationship) error {
	return r.db.Create(categoryRelationship).Error
}

func (r *sysCategoryRelationshipRepo) Update(categoryRelationship *model.SysCategoryRelationship) error {
	return r.db.Updates(categoryRelationship).Error
}

func (r *sysCategoryRelationshipRepo) Delete(categoryRelationshipID int) error {
	return r.db.Delete(&model.SysCategoryRelationship{}, categoryRelationshipID).Error
}
