package service

import (
	"github.com/Long-Plan/longplan-api/internal/core/domain"
	"github.com/Long-Plan/longplan-api/internal/core/model"
	"github.com/Long-Plan/longplan-api/internal/core/port"
)

type majorService struct {
	majorRepo port.SysMajorRepo
}

func NewMajorService(majorRepo port.SysMajorRepo) domain.MajorService {
	return &majorService{
		majorRepo: majorRepo,
	}
}

func (s *majorService) GetAll() ([]model.SysMajor, error) {
	return s.majorRepo.GetAll()
}

func (s *majorService) Create(major model.SysMajor) error {
	return s.majorRepo.Create(&major)
}

func (s *majorService) Update(major model.SysMajor) error {
	return s.majorRepo.Update(&major)
}

func (s *majorService) Delete(majorID int) error {
	return s.majorRepo.Delete(majorID)
}
