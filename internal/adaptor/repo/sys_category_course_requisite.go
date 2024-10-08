package repo

import (
	"github.com/Long-Plan/longplan-api/internal/core/model"
	"github.com/Long-Plan/longplan-api/internal/core/port"
	"gorm.io/gorm"
)

type sysCategoryCourseRequisiteRepo struct {
	db *gorm.DB
}

func NewSysCategoryCourseRequisiteRepo(db *gorm.DB) port.SysCategoryCourseRequisiteRepo {
	return &sysCategoryCourseRequisiteRepo{
		db: db,
	}
}

func (r *sysCategoryCourseRequisiteRepo) GetByCategoryCourseID(categoryCourseID int) ([]model.SysCategoryCourseRequisite, error) {
	var categoryCourseRequisites []model.SysCategoryCourseRequisite
	err := r.db.Where("category_course_id = ?", categoryCourseID).Find(&categoryCourseRequisites).Error
	return categoryCourseRequisites, err
}

func (r *sysCategoryCourseRequisiteRepo) Create(categoryCourseRequisite *model.SysCategoryCourseRequisite) error {
	return r.db.Create(categoryCourseRequisite).Error
}

func (r *sysCategoryCourseRequisiteRepo) Update(categoryCourseRequisite *model.SysCategoryCourseRequisite) error {
	return r.db.Updates(categoryCourseRequisite).Error
}

func (r *sysCategoryCourseRequisiteRepo) Delete(categoryCourseRequisiteId int) error {
	return r.db.Delete(&model.SysCategoryCourseRequisite{}, categoryCourseRequisiteId).Error
}
