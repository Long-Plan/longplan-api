package model

import "time"

type SysCategoryCourse struct {
	ID         int       `gorm:"primaryKey;autoIncrement" json:"id"`
	CategoryID int       `gorm:"column:category_id;type:int;not null" json:"category_id"`
	CourseNo   string    `gorm:"column:course_no;type:varchar(6);not null" json:"course_no"`
	Semester   *int      `gorm:"column:semester;type:int" json:"semester,omitempty"`
	Year       *int      `gorm:"column:year;type:int" json:"year,omitempty"`
	Credit     int       `gorm:"column:credit;type:int;default:0;not null" json:"credit"`
	CreatedAt  time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	Detail     *SysCourseDetail             `gorm:"foreignKey:CourseNo;references:CourseNo" json:"detail,omitempty"`
	Requisites []SysCategoryCourseRequisite `gorm:"foreignKey:CategoryCoursesID;references:ID" json:"requisites"`
}
