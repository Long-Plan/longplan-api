package model

type StudentCurriculumCourse struct {
	ID                  int    `gorm:"primaryKey;autoIncrement" json:"id"`
	StudentCurriculumID int    `gorm:"column:student_curriculum_id;type:int;not null" json:"student_curriculum_id"`
	Semester            int    `gorm:"column:semester;type:int;not null" json:"semester"`
	Year                int    `gorm:"column:year;type:int;not null" json:"year"`
	CourseNo            string `gorm:"column:course_no;type:varchar(6)" json:"course_no"`
	CategoryID          int    `gorm:"column:category_id;type:int" json:"category_id"`
	UpdatedAt           string `gorm:"autoUpdateTime" json:"updated_at"`
}
