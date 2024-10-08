package port

import "github.com/Long-Plan/longplan-api/internal/core/model"

type AccountTypeRepo interface {
	GetAll() ([]model.AccountType, error)
	Create(accountType *model.AccountType) error
	Update(accountType *model.AccountType) error
	Delete(accountTypeID int) error
}
