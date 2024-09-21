package repo

import (
	"github.com/Long-Plan/longplan-api/internal/core/model"
	"github.com/Long-Plan/longplan-api/internal/core/port"
	"gorm.io/gorm"
)

type accountTypeRepo struct {
	db *gorm.DB
}

func NewAccountTypeRepo(db *gorm.DB) port.AccountTypeRepo {
	return &accountTypeRepo{db}
}

func (r *accountTypeRepo) GetAll() ([]model.AccountType, error) {
	var accountTypes []model.AccountType
	if err := r.db.Find(&accountTypes).Error; err != nil {
		return nil, err
	}
	return accountTypes, nil
}

func (r *accountTypeRepo) Create(accountType *model.AccountType) error {
	return r.db.Create(accountType).Error
}

func (r *accountTypeRepo) Update(accountType *model.AccountType) error {
	return r.db.Updates(accountType).Error
}

func (r *accountTypeRepo) Delete(accountTypeID int) error {
	return r.db.Delete(&model.AccountType{}, accountTypeID).Error
}
