package domain

import "github.com/Long-Plan/longplan-api/internal/core/model"

type AccountService interface {
	GetByCMUITAccount(cmuitAccount string) (*model.Account, error)
	Save(account model.Account) error
	Delete(cmuitAccount string) error
}
