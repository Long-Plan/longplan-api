package port

import "github.com/Long-Plan/longplan-api/internal/core/model"

type SysCategoryCourseRequisiteRepo interface {
	GetByCategoryCourseID(categoryCourseID int) ([]model.SysCategoryCourseRequisite, error)
	Create(categoryCourseRequisite *model.SysCategoryCourseRequisite) error
	Update(categoryCourseRequisite *model.SysCategoryCourseRequisite) error
	Delete(categoryCourseRequisiteId int) error
}
