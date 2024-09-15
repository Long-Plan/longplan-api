package port

import "github.com/Long-Plan/longplan-api/internal/core/model"

type SysCategoryTypeRepo interface {
	GetAll() ([]model.SysCategoryType, error)
	Create(categoryType *model.SysCategoryType) error
	Update(categoryType *model.SysCategoryType) error
	Delete(categoryTypeID int) error
}
