package port

import "github.com/Long-Plan/longplan-api/internal/core/model"

type StudentRepo interface {
	GetByStudentCode(studentCode int) (*model.Student, error)
	Save(student *model.Student) error
	Delete(studentCode int) error
}
