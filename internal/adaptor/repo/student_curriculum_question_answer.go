package repo

import (
	"github.com/Long-Plan/longplan-api/internal/core/model"
	"github.com/Long-Plan/longplan-api/internal/core/port"
	"gorm.io/gorm"
)

type studentCurriculumQuestionAnswerRepo struct {
	db *gorm.DB
}

func NewStudentCurriculumQuestionAnswerRepo(db *gorm.DB) port.StudentCurriculumQuestionAnswerRepo {
	return &studentCurriculumQuestionAnswerRepo{db}
}

func (r *studentCurriculumQuestionAnswerRepo) GetByStudentCurriculumID(studentCurriculumID int) ([]model.StudentCurriculumQuestionAnswer, error) {
	var answers []model.StudentCurriculumQuestionAnswer
	if err := r.db.Where("student_curriculum_id = ?", studentCurriculumID).Find(&answers).Error; err != nil {
		return nil, err
	}
	return answers, nil
}

func (r *studentCurriculumQuestionAnswerRepo) Create(studentCurriculumQuestionAnswer *model.StudentCurriculumQuestionAnswer) error {
	return r.db.Create(studentCurriculumQuestionAnswer).Error
}

func (r *studentCurriculumQuestionAnswerRepo) Updates(studentCurriculumQuestionAnswer []model.StudentCurriculumQuestionAnswer) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		for _, answer := range studentCurriculumQuestionAnswer {
			if err := tx.Updates(answer).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

func (r *studentCurriculumQuestionAnswerRepo) Delete(studentCurriculumQuestionAnswerID int) error {
	return r.db.Delete(&model.StudentCurriculumQuestionAnswer{}, studentCurriculumQuestionAnswerID).Error
}
