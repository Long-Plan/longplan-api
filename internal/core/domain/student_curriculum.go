package domain

import "github.com/Long-Plan/longplan-api/internal/core/model"

type StudentCurriculumService interface {
	GetByStudentCode(studentCode int) ([]model.StudentCurriculum, error)
	GetByStudentCurriculumID(studentCurriculumID int) (*model.StudentCurriculum, error)
	Create(studentCurriculum model.StudentCurriculum) error
	Update(studentCurriculum model.StudentCurriculum) error
	Delete(studentCurriculumID int) error
	UpdateCourses(studentCurriculumID int, courses []model.StudentCurriculumCourse) error
	UpdateQuestionAnswers(studentCurriculumID int, questions []model.StudentCurriculumQuestionAnswer) error
}
