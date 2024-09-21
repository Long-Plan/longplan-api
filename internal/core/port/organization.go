package port

import "github.com/Long-Plan/longplan-api/internal/core/model"

type OrganizationRepo interface {
	GetAll() ([]model.Organization, error)
	Create(organization *model.Organization) error
	Update(organization *model.Organization) error
	Delete(organizationID int) error
}
