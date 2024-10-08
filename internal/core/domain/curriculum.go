package domain

import (
	"github.com/Long-Plan/longplan-api/internal/core/dto"
)

type CurriculumService interface {
	GetAll() ([]dto.Curriculum, error)
	GetAllByMajorID(majorID int) ([]dto.Curriculum, error)
	GetByID(curriculumID int) (*dto.Curriculum, error)
	GetCoursesByCurriculumID(curriculumID int) ([]dto.CategoryCourse, error)
	Create(curriculum dto.Curriculum) error
	Update(curriculum dto.Curriculum) error
	Delete(curriculumID int) error
}
