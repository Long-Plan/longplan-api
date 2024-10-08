package service

import (
	"github.com/Long-Plan/longplan-api/internal/core/domain"
	"github.com/Long-Plan/longplan-api/internal/core/model"
	"github.com/Long-Plan/longplan-api/internal/core/port"
)

type studentService struct {
	studentRepo port.StudentRepo
}

func NewStudentService(studentRepo port.StudentRepo) domain.StudentService {
	return &studentService{
		studentRepo: studentRepo,
	}
}

func (s *studentService) GetByStudentCode(studentCode int) (*model.Student, error) {
	student, err := s.studentRepo.GetByStudentCode(studentCode)
	if err != nil {
		return nil, err
	}
	return student, nil
}

func (s *studentService) Save(student model.Student) error {
	return s.studentRepo.Save(&student)
}

func (s *studentService) Delete(studentCode int) error {
	return s.studentRepo.Delete(studentCode)
}
