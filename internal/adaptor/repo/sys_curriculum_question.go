package repo

import (
	"github.com/Long-Plan/longplan-api/internal/core/model"
	"github.com/Long-Plan/longplan-api/internal/core/port"
	"gorm.io/gorm"
)

type sysCurriculumQuestionRepo struct {
	db *gorm.DB
}

func NewSysCurriculumQuestionRepo(db *gorm.DB) port.SysCurriculumQuestionRepo {
	return &sysCurriculumQuestionRepo{db}
}

func (r *sysCurriculumQuestionRepo) GetByCurriculumID(curriculumID int) ([]model.SysCurriculumQuestion, error) {
	var curriculumQuestions []model.SysCurriculumQuestion
	if err := r.db.Where("curriculum_id = ?", curriculumID).Find(&curriculumQuestions).Error; err != nil {
		return nil, err
	}
	return curriculumQuestions, nil
}

func (r *sysCurriculumQuestionRepo) Create(curriculumQuestion *model.SysCurriculumQuestion) error {
	return r.db.Create(curriculumQuestion).Error
}

func (r *sysCurriculumQuestionRepo) Update(curriculumQuestion *model.SysCurriculumQuestion) error {
	return r.db.Updates(curriculumQuestion).Error
}

func (r *sysCurriculumQuestionRepo) Delete(curriculumQuestionID int) error {
	return r.db.Delete(&model.SysCurriculumQuestion{}, curriculumQuestionID).Error
}
