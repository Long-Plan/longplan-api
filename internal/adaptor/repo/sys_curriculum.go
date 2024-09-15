package repo

import (
	"github.com/Long-Plan/longplan-api/internal/core/model"
	"github.com/Long-Plan/longplan-api/internal/core/port"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type sysCurriculumRepo struct {
	db *gorm.DB
}

func NewSysCurriculumRepo(db *gorm.DB) port.SysCurriculumRepo {
	return &sysCurriculumRepo{db}
}

func (r *sysCurriculumRepo) GetAll() ([]model.SysCurriculum, error) {
	var curriculums []model.SysCurriculum
	if err := r.db.Preload(clause.Associations).Preload("Questions.Choices").Find(&curriculums).Error; err != nil {
		return nil, err
	}
	return curriculums, nil
}

func (r *sysCurriculumRepo) GetAllByMajorID(majorID int) ([]model.SysCurriculum, error) {
	var curriculums []model.SysCurriculum
	if err := r.db.Preload(clause.Associations).Preload("Questions.Choices").Where("major_id = ?", majorID).Find(&curriculums).Error; err != nil {
		return nil, err
	}
	return curriculums, nil
}

func (r *sysCurriculumRepo) GetByID(curriculumId int) (*model.SysCurriculum, error) {
	var curriculum model.SysCurriculum
	if err := r.db.Preload(clause.Associations).Preload("Questions.Choices").First(&curriculum, curriculumId).Error; err != nil {
		return nil, err
	}
	return &curriculum, nil
}

func (r *sysCurriculumRepo) Create(curriculum *model.SysCurriculum) error {
	return r.db.Create(curriculum).Error
}

func (r *sysCurriculumRepo) Update(curriculum *model.SysCurriculum) error {
	return r.db.Updates(curriculum).Error
}

func (r *sysCurriculumRepo) Delete(curriculumId int) error {
	return r.db.Delete(&model.SysCurriculum{}, curriculumId).Error
}
