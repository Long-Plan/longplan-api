package port

import "github.com/Long-Plan/longplan-api/internal/core/model"

type StudentCurriculumRepo interface {
	GetByStudentCode(studentCode int) ([]model.StudentCurriculum, error)
	GetByStudentCurriculumID(studentCurriculumID int) (*model.StudentCurriculum, error)
	Create(studentCurriculum *model.StudentCurriculum) error
	Update(studentCurriculum *model.StudentCurriculum) error
	Delete(studentCurriculumID int) error
}
