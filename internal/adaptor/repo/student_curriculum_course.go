package repo

import (
	"github.com/Long-Plan/longplan-api/internal/core/model"
	"github.com/Long-Plan/longplan-api/internal/core/port"
	"gorm.io/gorm"
)

type StudentCurriculumCourseRepo struct {
	db *gorm.DB
}

func NewStudentCurriculumCourseRepo(db *gorm.DB) port.StudentCurriculumCourseRepo {
	return &StudentCurriculumCourseRepo{db}
}

func (r *StudentCurriculumCourseRepo) GetByStudentCurriculumID(studentCurriculumID int) ([]model.StudentCurriculumCourse, error) {
	var studentCurriculumCourses []model.StudentCurriculumCourse
	if err := r.db.Where("student_curriculum_id = ?", studentCurriculumID).Find(&studentCurriculumCourses).Error; err != nil {
		return nil, err
	}
	return studentCurriculumCourses, nil
}

func (r *StudentCurriculumCourseRepo) Create(studentCurriculumCourse *model.StudentCurriculumCourse) error {
	return r.db.Create(studentCurriculumCourse).Error
}

func (r *StudentCurriculumCourseRepo) Updates(studentCurriculumCourse []model.StudentCurriculumCourse) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		for _, course := range studentCurriculumCourse {
			if err := tx.Updates(course).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

func (r *StudentCurriculumCourseRepo) Delete(studentCurriculumCourseID int) error {
	return r.db.Delete(&model.StudentCurriculumCourse{}, studentCurriculumCourseID).Error
}
