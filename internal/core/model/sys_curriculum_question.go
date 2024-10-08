package model

import "time"

type SysCurriculumQuestion struct {
	ID           int       `gorm:"primaryKey;autoIncrement" json:"id"`
	CurriculumID int       `gorm:"column:curriculum_id;type:int" json:"curriculum_id"`
	NameTH       string    `gorm:"column:name_th;type:varchar(255)" json:"name_th"`
	NameEN       string    `gorm:"column:name_en;type:varchar(255)" json:"name_en"`
	CreatedAt    time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	Choices []SysCurriculumQuestionChoice `gorm:"foreignKey:QuestionID;references:ID" json:"choices"`
}
