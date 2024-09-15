package repo

import (
	"github.com/Long-Plan/longplan-api/internal/core/model"
	"github.com/Long-Plan/longplan-api/internal/core/port"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type sysCategoryRepo struct {
	db *gorm.DB
}

func NewSysCategoryRepo(db *gorm.DB) port.SysCategoryRepo {
	return &sysCategoryRepo{db}
}

func (r *sysCategoryRepo) GetByCurriculumID(curriculumID int) ([]model.SysCategory, error) {
	var categories []model.SysCategory
	if err := r.db.Preload(clause.Associations).Preload("Courses.Requisites").Where("curriculum_id = ?", curriculumID).Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func (r *sysCategoryRepo) Create(category *model.SysCategory) error {
	return r.db.Create(category).Error
}

func (r *sysCategoryRepo) Update(category *model.SysCategory) error {
	return r.db.Updates(category).Error
}

func (r *sysCategoryRepo) Delete(categoryID int) error {
	return r.db.Delete(&model.SysCategory{}, categoryID).Error
}
