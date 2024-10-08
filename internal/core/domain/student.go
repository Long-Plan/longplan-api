package domain

import "github.com/Long-Plan/longplan-api/internal/core/model"

type StudentService interface {
	GetByStudentCode(studentCode int) (*model.Student, error)
	Save(student model.Student) error
	Delete(studentCode int) error
}
