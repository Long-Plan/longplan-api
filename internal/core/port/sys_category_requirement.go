package port

import "github.com/Long-Plan/longplan-api/internal/core/model"

type SysCategoryRequirementRepo interface {
	GetByCategoryID(categoryID int) ([]model.SysCategoryRequirement, error)
	Create(categoryRequirement *model.SysCategoryRequirement) error
	Update(categoryRequirement *model.SysCategoryRequirement) error
	Delete(categoryRequirementID int) error
}
