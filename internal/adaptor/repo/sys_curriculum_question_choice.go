package repo

import (
	"github.com/Long-Plan/longplan-api/internal/core/model"
	"github.com/Long-Plan/longplan-api/internal/core/port"
	"gorm.io/gorm"
)

type sysCurriculumQuestionChoiceRepo struct {
	db *gorm.DB
}

func NewSysCurriculumQuestionChoiceRepo(db *gorm.DB) port.SysCurriculumQuestionChoiceRepo {
	return &sysCurriculumQuestionChoiceRepo{db}
}

func (r *sysCurriculumQuestionChoiceRepo) GetByQuestionID(questionID int) ([]model.SysCurriculumQuestionChoice, error) {
	var curriculumQuestionChoices []model.SysCurriculumQuestionChoice
	if err := r.db.Where("question_id = ?", questionID).Find(&curriculumQuestionChoices).Error; err != nil {
		return nil, err
	}
	return curriculumQuestionChoices, nil
}

func (r *sysCurriculumQuestionChoiceRepo) Create(curriculumQuestionChoice *model.SysCurriculumQuestionChoice) error {
	return r.db.Create(curriculumQuestionChoice).Error
}

func (r *sysCurriculumQuestionChoiceRepo) Update(curriculumQuestionChoice *model.SysCurriculumQuestionChoice) error {
	return r.db.Updates(curriculumQuestionChoice).Error
}

func (r *sysCurriculumQuestionChoiceRepo) Delete(curriculumQuestionChoiceID int) error {
	return r.db.Delete(&model.SysCurriculumQuestionChoice{}, curriculumQuestionChoiceID).Error
}
