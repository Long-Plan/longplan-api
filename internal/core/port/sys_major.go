package port

import "github.com/Long-Plan/longplan-api/internal/core/model"

type SysMajorRepo interface {
	GetAll() ([]model.SysMajor, error)
	Create(major *model.SysMajor) error
	Update(major *model.SysMajor) error
	Delete(majorId int) error
}
