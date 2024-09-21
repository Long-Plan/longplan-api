package model

type Student struct {
	Code           int    `gorm:"primaryKey" json:"code"`
	MajorID        int    `gorm:"column:major_id" json:"major_id,omitempty"`
	IsTermAccepted bool   `gorm:"column:is_term_accepted;not null;default:false" json:"is_term_accepted"`
	CreatedAt      string `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt      string `gorm:"autoUpdateTime" json:"updated_at"`
}
