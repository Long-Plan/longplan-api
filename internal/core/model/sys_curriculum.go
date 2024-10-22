package model

import "time"

type SysCurriculum struct {
	ID        int       `gorm:"primaryKey;autoIncrement" json:"id"`
	MajorID   int       `gorm:"column:major_id;type:int;not null" json:"major_id"`
	Code      string    `gorm:"column:code;type:varchar(255);unique" json:"code"`
	NameTH    string    `gorm:"column:name_th;type:varchar(255);unique" json:"name_th"`
	NameEN    string    `gorm:"column:name_en;type:varchar(255);unique" json:"name_en"`
	ShortName string    `gorm:"column:short_name;type:varchar(10)" json:"short_name"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	Categories []SysCategory           `gorm:"foreignKey:CurriculumID;references:ID" json:"categories"`
	Questions  []SysCurriculumQuestion `gorm:"foreignKey:CurriculumID;references:ID" json:"questions"`
}

func (SysCurriculum) TableName() string {
	return "sys_curricula"
}
