package model

import "time"

type SysCourseDetail struct {
	ID           int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	CourseNo     string    `gorm:"column:course_no;type:varchar(6);unique;not null" json:"course_no"`
	TitleLongTH  string    `gorm:"column:title_long_th;type:varchar(255);unique" json:"title_long_th"`
	TitleLongEN  string    `gorm:"column:title_long_en;type:varchar(255);unique" json:"title_long_en"`
	TitleShortEN *string   `gorm:"column:title_short_en;type:varchar(255)" json:"title_short_en,omitempty"`
	CourseDescTH *string   `gorm:"column:course_desc_th;type:varchar(255)" json:"course_desc_th,omitempty"`
	CourseDescEN *string   `gorm:"column:course_desc_en;type:varchar(255)" json:"course_desc_en,omitempty"`
	Credit       int       `gorm:"column:credit;type:int;default:0;not null" json:"credit"`
	Prerequisite *string   `gorm:"column:prerequisite;type:varchar(255)" json:"prerequisite,omitempty"`
	CreatedAt    time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
