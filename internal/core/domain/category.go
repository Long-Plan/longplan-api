package domain

import (
	"github.com/Long-Plan/longplan-api/internal/core/dto"
	"github.com/Long-Plan/longplan-api/internal/core/model"
)

type CategoryService interface {
	GetTypes() ([]model.SysCategoryType, error)
	CreateType(types model.SysCategoryType) error
	UpdateType(types model.SysCategoryType) error
	DeleteType(typeID int) error
	GetByCurriculumID(curriculumID int) ([]dto.Category, error)
	Create(category dto.Category) error
	Update(category dto.Category) error
	Delete(categoryID int) error
}
