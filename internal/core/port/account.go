package port

import "github.com/Long-Plan/longplan-api/internal/core/model"

type AccountRepo interface {
	GetByCMUITAccount(CMUITAccount string) (*model.Account, error)
	Save(account *model.Account) error
	Delete(CMUITAccount string) error
}
