package port

import "github.com/Long-Plan/longplan-api/internal/core/model"

type StudentCurriculumCourseRepo interface {
	GetByStudentCurriculumID(studentCurriculumID int) ([]model.StudentCurriculumCourse, error)
	Create(studentCurriculumCourse *model.StudentCurriculumCourse) error
	Updates(studentCurriculumCourse []model.StudentCurriculumCourse) error
	Delete(studentCurriculumCourseID int) error
}
