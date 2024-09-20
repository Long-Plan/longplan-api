package service

import (
	"github.com/Long-Plan/longplan-api/internal/core/domain"
	"github.com/Long-Plan/longplan-api/internal/core/model"
	"github.com/Long-Plan/longplan-api/internal/core/port"
)

type courseDetailService struct {
	courseDetailRepo port.SysCourseDetailRepo
}

func NewCourseDetailService(courseDetailRepo port.SysCourseDetailRepo) domain.CourseDetailService {
	return &courseDetailService{
		courseDetailRepo: courseDetailRepo,
	}
}

func (s *courseDetailService) GetAll() ([]model.SysCourseDetail, error) {
	return s.courseDetailRepo.GetAll()
}

func (s *courseDetailService) GetByCourseNo(courseNo string) (*model.SysCourseDetail, error) {
	return s.courseDetailRepo.GetByCourseNo(courseNo)
}

func (s *courseDetailService) Create(courseDetail model.SysCourseDetail) error {
	return s.courseDetailRepo.Create(&courseDetail)
}

func (s *courseDetailService) Update(courseDetail model.SysCourseDetail) error {
	return s.courseDetailRepo.Update(&courseDetail)
}

func (s *courseDetailService) Delete(courseNo string) error {
	return s.courseDetailRepo.Delete(courseNo)
}
