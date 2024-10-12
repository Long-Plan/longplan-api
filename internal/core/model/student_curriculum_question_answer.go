package model

import "time"

type StudentCurriculumQuestionAnswer struct {
	ID                  int       `gorm:"primaryKey;autoIncrement" json:"id"`
	StudentCurriculumID int       `gorm:"column:student_curriculum_id;type:int;not null" json:"student_curriculum_id"`
	QuestionID          int       `gorm:"column:question_id;type:int;not null" json:"question_id"`
	ChoiceID            int       `gorm:"column:choice_id;type:int;not null" json:"choice_id"`
	CreatedAt           time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt           time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
