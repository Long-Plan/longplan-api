package repo

import (
	"github.com/Long-Plan/longplan-api/internal/core/model"
	"github.com/Long-Plan/longplan-api/internal/core/port"
	"gorm.io/gorm"
)

type accountRepo struct {
	db *gorm.DB
}

func NewAccountRepo(db *gorm.DB) port.AccountRepo {
	return &accountRepo{db}
}

func (r *accountRepo) GetByCMUITAccount(CMUITAccount string) (*model.Account, error) {
	var account model.Account
	if err := r.db.Where("cmuitaccount = ?", CMUITAccount).First(&account).Error; err != nil {
		return nil, err
	}
	return &account, nil
}

func (r *accountRepo) Save(account *model.Account) error {
	return r.db.Save(account).Error
}

func (r *accountRepo) Delete(CMUITAccount string) error {
	return r.db.Delete(&model.Account{}, CMUITAccount).Error
}
