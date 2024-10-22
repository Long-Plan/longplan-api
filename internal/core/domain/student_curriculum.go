package domain

import (
	"github.com/Long-Plan/longplan-api/internal/core/dto"
	"github.com/Long-Plan/longplan-api/internal/core/model"
)

type StudentCurriculumService interface {
	GetByStudentCode(studentCode int) ([]dto.StudentCurriculum, error)
	GetByStudentCurriculumID(studentCurriculumID int) (*dto.StudentCurriculum, error)
	Create(studentCurriculum dto.StudentCurriculumCreate) (*int, error)
	Update(studentCurriculum model.StudentCurriculum) error
	Delete(studentCurriculumID int) error
	UpdateCourses(studentCurriculumID int, courses []model.StudentCurriculumCourse) error
	UpdateQuestionAnswers(studentCurriculumID int, questions []dto.StudentCurriculumQuestionAnswer) error
}
