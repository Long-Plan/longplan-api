package dto

import "github.com/Long-Plan/longplan-api/internal/core/model"

type Category struct {
	ID        int    `json:"id"`
	NameTH    string `json:"name_th"`
	NameEN    string `json:"name_en"`
	AtLeast   int    `json:"at_least"`
	Credit    int    `json:"credit"`
	TypeID    int    `json:"type_id"`
	Note      string `json:"note"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`

	Requirements  []CategoryRequirement  `json:"requirements"`
	Relationships []CategoryRelationship `json:"relationships"`
	Courses       []CategoryCourse       `json:"courses"`
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
	QuestionID      int  `json:"question_id"`
	ChoiceID        int  `json:"choice_id"`
	CrossCategoryID int  `json:"cross_category_id"`
}

type CategoryCourse struct {
	ID       int     `json:"id"`
	CourseNo *string `json:"course_no,omitempty"`
	Semester *int    `json:"semester,omitempty"`
	Years    *int    `json:"years,omitempty"`

	Requisites []CategoryCourseRequisite `json:"requisites"`
}

type CategoryCourseRequisite struct {
	ID              int                 `json:"id"`
	RelatedCourseNo string              `json:"related_course_no"`
	RequisiteType   model.RequisiteType `json:"requisite_type"`
}
