package port

import "github.com/Long-Plan/longplan-api/internal/core/model"

type StudentCurriculumQuestionAnswerRepo interface {
	GetByStudentCurriculumID(studentCurriculumID int) ([]model.StudentCurriculumQuestionAnswer, error)
	Create(studentCurriculumQuestionAnswer *model.StudentCurriculumQuestionAnswer) error
	Updates(studentCurriculumQuestionAnswer []model.StudentCurriculumQuestionAnswer) error
	Delete(studentCurriculumQuestionAnswerID int) error
}
