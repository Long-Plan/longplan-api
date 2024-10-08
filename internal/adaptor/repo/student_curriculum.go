package repo

import (
	"github.com/Long-Plan/longplan-api/internal/core/model"
	"github.com/Long-Plan/longplan-api/internal/core/port"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type studentCurriculumRepo struct {
	db *gorm.DB
}

func NewStudentCurriculumRepo(db *gorm.DB) port.StudentCurriculumRepo {
	return &studentCurriculumRepo{db}
}

func (r *studentCurriculumRepo) GetByStudentCode(studentCode int) ([]model.StudentCurriculum, error) {
	var studentCurriculums []model.StudentCurriculum
	if err := r.db.Preload(clause.Associations).Where("student_code = ?", studentCode).Find(&studentCurriculums).Error; err != nil {
		return nil, err
	}
	return studentCurriculums, nil
}

func (r *studentCurriculumRepo) GetByStudentCurriculumID(studentCurriculumID int) (*model.StudentCurriculum, error) {
	var studentCurriculum model.StudentCurriculum
	if err := r.db.Preload(clause.Associations).Where("id = ?", studentCurriculumID).First(&studentCurriculum).Error; err != nil {
		return nil, err
	}
	return &studentCurriculum, nil
}

func (r *studentCurriculumRepo) Create(studentCurriculum *model.StudentCurriculum) error {
	return r.db.Create(studentCurriculum).Error
}

func (r *studentCurriculumRepo) Update(studentCurriculum *model.StudentCurriculum) error {
	return r.db.Updates(studentCurriculum).Error
}

func (r *studentCurriculumRepo) Delete(studentCurriculumID int) error {
	return r.db.Delete(&model.StudentCurriculum{}, studentCurriculumID).Error
}
