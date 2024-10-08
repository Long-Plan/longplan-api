package repo

import (
	"github.com/Long-Plan/longplan-api/internal/core/model"
	"github.com/Long-Plan/longplan-api/internal/core/port"
	"gorm.io/gorm"
)

type sysCategoryCourseRepo struct {
	db *gorm.DB
}

func NewSysCategoryCourseRepo(db *gorm.DB) port.SysCategoryCourseRepo {
	return &sysCategoryCourseRepo{db}
}

func (r *sysCategoryCourseRepo) GetByCategoryID(categoryID int) ([]model.SysCategoryCourse, error) {
	var categoryCourses []model.SysCategoryCourse
	err := r.db.Where("category_id = ?", categoryID).Find(&categoryCourses).Error
	return categoryCourses, err
}

func (r *sysCategoryCourseRepo) GetByCurriculumID(curriculumID int) ([]model.SysCategoryCourse, error) {
	var categoryCourses []model.SysCategoryCourse
	err := r.db.Joins("inner join sys_categories on sys_categories.id = sys_category_courses.category_id").Where("sys_categories.curriculum_id = ?", curriculumID).Find(&categoryCourses).Error
	return categoryCourses, err
}

func (r *sysCategoryCourseRepo) Create(categoryCourse *model.SysCategoryCourse) error {
	return r.db.Create(categoryCourse).Error
}

func (r *sysCategoryCourseRepo) Update(categoryCourse *model.SysCategoryCourse) error {
	return r.db.Updates(categoryCourse).Error
}

func (r *sysCategoryCourseRepo) Delete(categoryCourseID int) error {
	return r.db.Delete(&model.SysCategoryCourse{}, categoryCourseID).Error
}
