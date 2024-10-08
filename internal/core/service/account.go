package service

import (
	"github.com/Long-Plan/longplan-api/internal/core/domain"
	"github.com/Long-Plan/longplan-api/internal/core/model"
	"github.com/Long-Plan/longplan-api/internal/core/port"
)

type accountService struct {
	accountRepo      port.AccountRepo
	accountTypeRepo  port.AccountTypeRepo
	organizationRepo port.OrganizationRepo
}

func NewAccountService(
	accountRepo port.AccountRepo,
	accountTypeRepo port.AccountTypeRepo,
	organizationRepo port.OrganizationRepo,
) domain.AccountService {
	return &accountService{
		accountRepo:      accountRepo,
		accountTypeRepo:  accountTypeRepo,
		organizationRepo: organizationRepo,
	}
}

func (s *accountService) GetByCMUITAccount(CMUITAccount string) (*model.Account, error) {
	account, err := s.accountRepo.GetByCMUITAccount(CMUITAccount)
	if err != nil {
		return nil, err
	}
	return account, nil
}

func (s *accountService) Save(account model.Account) error {
	return s.accountRepo.Save(&account)
}

func (s *accountService) Delete(CMUITAccount string) error {
	return s.accountRepo.Delete(CMUITAccount)
}
