package port

import "github.com/Long-Plan/longplan-api/internal/core/model"

type SysCategoryRepo interface {
	GetByCurriculumID(curriculumID int) ([]model.SysCategory, error)
	Create(category *model.SysCategory) error
	Update(category *model.SysCategory) error
	Delete(categoryID int) error
}
