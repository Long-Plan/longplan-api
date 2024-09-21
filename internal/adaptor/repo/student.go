package repo

import (
	"github.com/Long-Plan/longplan-api/internal/core/model"
	"github.com/Long-Plan/longplan-api/internal/core/port"
	"gorm.io/gorm"
)

type studentRepo struct {
	db *gorm.DB
}

func NewStudentRepo(db *gorm.DB) port.StudentRepo {
	return &studentRepo{db}
}

func (r *studentRepo) GetByStudentCode(studentCode int) (*model.Student, error) {
	var student model.Student
	if err := r.db.Where("student_code = ?", studentCode).First(&student).Error; err != nil {
		return nil, err
	}
	return &student, nil
}

func (r *studentRepo) Save(student *model.Student) error {
	return r.db.Save(student).Error
}

func (r *studentRepo) Delete(studentCode int) error {
	return r.db.Delete(&model.Student{}, studentCode).Error
}
