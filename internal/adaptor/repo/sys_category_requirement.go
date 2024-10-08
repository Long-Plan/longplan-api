package repo

import (
	"github.com/Long-Plan/longplan-api/internal/core/model"
	"github.com/Long-Plan/longplan-api/internal/core/port"
	"gorm.io/gorm"
)

type sysCategoryRequirementRepo struct {
	db *gorm.DB
}

func NewSysCategoryRequirementRepo(db *gorm.DB) port.SysCategoryRequirementRepo {
	return &sysCategoryRequirementRepo{db}
}

func (r *sysCategoryRequirementRepo) GetByCategoryID(categoryID int) ([]model.SysCategoryRequirement, error) {
	var categoryRequirements []model.SysCategoryRequirement
	err := r.db.Where("category_id = ?", categoryID).Find(&categoryRequirements).Error
	return categoryRequirements, err
}

func (r *sysCategoryRequirementRepo) Create(categoryRequirement *model.SysCategoryRequirement) error {
	return r.db.Create(categoryRequirement).Error
}

func (r *sysCategoryRequirementRepo) Update(categoryRequirement *model.SysCategoryRequirement) error {
	return r.db.Updates(categoryRequirement).Error
}

func (r *sysCategoryRequirementRepo) Delete(categoryRequirementID int) error {
	return r.db.Delete(&model.SysCategoryRequirement{}, categoryRequirementID).Error
}
