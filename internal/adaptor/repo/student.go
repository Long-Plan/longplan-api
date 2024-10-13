package repo

import (
	"errors"

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
	if err := r.db.Where("code = ?", studentCode).First(&student).Error; err != nil {
		return nil, err
	}
	return &student, nil
}

func (r *studentRepo) Save(student *model.Student) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := r.db.Where("code = ?", student.Code).First(&model.Student{}).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return tx.Create(student).Error
			}
			return err
		}
		return tx.Save(student).Error
	})
}

func (r *studentRepo) Delete(studentCode int) error {
	return r.db.Delete(&model.Student{}, studentCode).Error
}
