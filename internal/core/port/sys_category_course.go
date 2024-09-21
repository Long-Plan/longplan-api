package port

import "github.com/Long-Plan/longplan-api/internal/core/model"

type SysCategoryCourseRepo interface {
	GetByCategoryID(categoryID int) ([]model.SysCategoryCourse, error)
	GetByCurriculumID(curriculumID int) ([]model.SysCategoryCourse, error)
	Create(categoryCourse *model.SysCategoryCourse) error
	Update(categoryCourse *model.SysCategoryCourse) error
	Delete(categoryCourseID int) error
}
