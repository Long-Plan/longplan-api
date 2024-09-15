package port

import "github.com/Long-Plan/longplan-api/internal/core/model"

type SysCategoryRelationshipRepo interface {
	GetByParentCategoryID(parentCategoryID int) ([]model.SysCategoryRelationship, error)
	Create(categoryRelationship *model.SysCategoryRelationship) error
	Update(categoryRelationship *model.SysCategoryRelationship) error
	Delete(categoryRelationshipID int) error
}
