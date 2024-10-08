package repo

import (
	"github.com/Long-Plan/longplan-api/internal/core/model"
	"github.com/Long-Plan/longplan-api/internal/core/port"
	"gorm.io/gorm"
)

type sysMajorRepo struct {
	db *gorm.DB
}

func NewSysMajorRepo(db *gorm.DB) port.SysMajorRepo {
	return &sysMajorRepo{db}
}

func (r *sysMajorRepo) GetAll() ([]model.SysMajor, error) {
	var majors []model.SysMajor
	if err := r.db.Find(&majors).Error; err != nil {
		return nil, err
	}
	return majors, nil
}

func (r *sysMajorRepo) Create(major *model.SysMajor) error {
	return r.db.Create(major).Error
}

func (r *sysMajorRepo) Update(major *model.SysMajor) error {
	return r.db.Updates(major).Error
}

func (r *sysMajorRepo) Delete(majorId int) error {
	return r.db.Delete(&model.SysMajor{}, majorId).Error
}
