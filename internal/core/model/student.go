package model

import "time"

type Student struct {
	Code           int       `gorm:"primaryKey" json:"code"`
	MajorID        int       `gorm:"column:major_id" json:"major_id,omitempty"`
	IsTermAccepted bool      `gorm:"column:is_term_accepted;not null;default:false" json:"is_term_accepted"`
	CreatedAt      time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
