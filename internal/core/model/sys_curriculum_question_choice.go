package model

import "time"

type SysCurriculumQuestionChoice struct {
	ID         int       `gorm:"primaryKey;autoIncrement" json:"id"`
	QuestionID int       `gorm:"column:question_id;type:int;not null" json:"question_id"`
	Position   int       `gorm:"column:position;type:int;default:1" json:"position"`
	NameTH     string    `gorm:"column:name_th;type:varchar(255)" json:"name_th"`
	NameEN     string    `gorm:"column:name_en;type:varchar(255)" json:"name_en"`
	CreatedAt  time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
