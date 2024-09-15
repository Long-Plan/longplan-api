package port

import "github.com/Long-Plan/longplan-api/internal/core/model"

type SysCurriculumQuestionChoiceRepo interface {
	GetByQuestionID(questionID int) ([]model.SysCurriculumQuestionChoice, error)
	Create(curriculumQuestionChoice *model.SysCurriculumQuestionChoice) error
	Update(curriculumQuestionChoice *model.SysCurriculumQuestionChoice) error
	Delete(curriculumQuestionChoiceID int) error
}
