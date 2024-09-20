package dto

import (
	"time"

	"github.com/Long-Plan/longplan-api/internal/core/model"
)

type Category struct {
	ID        int       `json:"id"`
	NameTH    string    `json:"name_th"`
	NameEN    string    `json:"name_en"`
	AtLeast   bool      `json:"at_least"`
	Credit    int       `json:"credit"`
	TypeID    int       `json:"type_id"`
	Note      string    `json:"note"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Requirements    []CategoryRequirement  `json:"requirements"`
	Relationships   []CategoryRelationship `json:"relationships"`
	ChildCategories []Category             `json:"child_categories"`
	Courses         []CategoryCourse       `json:"courses"`
}

type CategoryRequirement struct {
	ID     int     `json:"id"`
	Regex  *string `json:"regex,omitempty"`
	Credit *int    `json:"credit,omitempty"`
}

type CategoryRelationship struct {
	ID              int  `json:"id"`
	ChildCategoryID int  `json:"child_category_id"`
	RequireAll      bool `json:"require_all"`
	Position        int  `json:"position"`
	QuestionID      *int `json:"question_id"`
	ChoiceID        *int `json:"choice_id"`
	CrossCategoryID *int `json:"cross_category_id"`
}

type CategoryCourse struct {
	ID       int     `json:"id"`
	CourseNo *string `json:"course_no,omitempty"`
	Semester *int    `json:"semester,omitempty"`
	Year     *int    `json:"years,omitempty"`
	Credit   int     `json:"credit"`

	Requisites []CategoryCourseRequisite `json:"requisites"`
	Detail     model.SysCourseDetail     `json:"detail"`
}

type CategoryCourseRequisite struct {
	ID              int                 `json:"id"`
	RelatedCourseNo string              `json:"related_course_no"`
	RequisiteType   model.RequisiteType `json:"requisite_type"`
}
