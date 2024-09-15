package model

import "time"

type SysCategoryRelationship struct {
	ID               int       `gorm:"primaryKey;autoIncrement" json:"id"`
	ParentCategoryID int       `gorm:"column:parent_category_id;type:int;not null" json:"parent_category_id"`
	ChildCategoryID  int       `gorm:"column:child_category_id;type:int;not null" json:"child_category_id"`
	CurriculumID     int       `gorm:"column:curriculum_id;type:int;not null" json:"curriculum_id"`
	RequireAll       bool      `gorm:"column:require_all;type:boolean" json:"require_all"`
	Position         int       `gorm:"column:position;type:int;default:1" json:"position"`
	QuestionID       *int      `gorm:"column:question_id;type:int" json:"question_id,omitempty"`
	ChoiceID         *int      `gorm:"column:choice_id;type:int" json:"choice_id,omitempty"`
	CrossCategoryID  *int      `gorm:"column:cross_category_id;type:int" json:"cross_category_id,omitempty"`
	CreatedAt        time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt        time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
