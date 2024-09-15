package model

import "time"

type RequisiteType string

const (
	Any RequisiteType = "Any"
	All RequisiteType = "All"
	Co  RequisiteType = "Co"
)

type SysCategoryCourseRequisite struct {
	ID                int64         `gorm:"primaryKey;autoIncrement" json:"id"`
	CategoryCoursesID int           `gorm:"column:category_courses_id;type:int;not null" json:"category_courses_id"`
	RelatedCourseNo   string        `gorm:"column:related_course_no;type:varchar(6);not null" json:"related_course_no"`
	RequisiteType     RequisiteType `gorm:"column:requisite_type;type:requisite_type;not null" json:"requisite_type"`
	CreatedAt         time.Time     `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt         time.Time     `gorm:"autoUpdateTime" json:"updated_at"`
}
