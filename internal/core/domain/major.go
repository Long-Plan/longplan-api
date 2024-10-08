package domain

import "github.com/Long-Plan/longplan-api/internal/core/model"

type MajorService interface {
	GetAll() ([]model.SysMajor, error)
	Create(major model.SysMajor) error
	Update(major model.SysMajor) error
	Delete(majorID int) error
}
