package repo

import (
	"github.com/Long-Plan/longplan-api/internal/core/model"
	"github.com/Long-Plan/longplan-api/internal/core/port"
	"gorm.io/gorm"
)

type sysCourseDetailRepo struct {
	db *gorm.DB
}

func NewSysCourseDetailRepo(db *gorm.DB) port.SysCourseDetailRepo {
	return &sysCourseDetailRepo{db}
}

func (r *sysCourseDetailRepo) GetAll() ([]model.SysCourseDetail, error) {
	var courseDetails []model.SysCourseDetail
	if err := r.db.Find(&courseDetails).Error; err != nil {
		return nil, err
	}
	return courseDetails, nil
}

func (r *sysCourseDetailRepo) GetByCourseNo(courseNo string) (*model.SysCourseDetail, error) {
	var courseDetail model.SysCourseDetail
	if err := r.db.Where("course_no = ?", courseNo).First(&courseDetail).Error; err != nil {
		return nil, err
	}
	return &courseDetail, nil
}

func (r *sysCourseDetailRepo) Create(courseDetail *model.SysCourseDetail) error {
	return r.db.Create(courseDetail).Error
}

func (r *sysCourseDetailRepo) Update(courseDetail *model.SysCourseDetail) error {
	return r.db.Updates(courseDetail).Error
}

func (r *sysCourseDetailRepo) Delete(courseNo string) error {
	return r.db.Delete(&model.SysCourseDetail{}, courseNo).Error
}
