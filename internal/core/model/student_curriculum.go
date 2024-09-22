package model

type StudentCurriculum struct {
	ID           int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name         string `gorm:"column:name;type:varchar(255);not null" json:"name"`
	StudentCode  int    `gorm:"column:student_code;type:int;not null" json:"student_code"`
	CurriculumID int    `gorm:"column:curriculum_id;type:int;not null" json:"curriculum_id"`
	IsSystem     bool   `gorm:"column:is_system;type:bool;default:false;not null" json:"is_system"`
	IsDefault    bool   `gorm:"column:is_default;type:bool;default:false;not null" json:"is_default"`
	CreatedAt    string `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    string `gorm:"autoUpdateTime" json:"updated_at"`
}

func (StudentCurriculum) TableName() string {
	return "student_curricula"
}