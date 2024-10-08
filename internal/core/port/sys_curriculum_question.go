package port

import "github.com/Long-Plan/longplan-api/internal/core/model"

type SysCurriculumQuestionRepo interface {
	GetByCurriculumID(curriculumID int) ([]model.SysCurriculumQuestion, error)
	Create(curriculumQuestion *model.SysCurriculumQuestion) error
	Update(curriculumQuestion *model.SysCurriculumQuestion) error
	Delete(curriculumQuestionID int) error
}
