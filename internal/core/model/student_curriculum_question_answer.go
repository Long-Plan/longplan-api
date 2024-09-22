package model

type StudentCurriculumQuestionAnswer struct {
	ID           int    `gorm:"primaryKey;autoIncrement" json:"id"`
	StudentCode  int    `gorm:"column:student_code;type:int;not null" json:"student_code"`
	CurriculumID int    `gorm:"column:curriculum_id;type:int;not null" json:"curriculum_id"`
	QuestionID   int    `gorm:"column:question_id;type:int;not null" json:"question_id"`
	ChoiceID     int    `gorm:"column:choice_id;type:int;not null" json:"choice_id"`
	CreatedAt    string `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    string `gorm:"autoUpdateTime" json:"updated_at"`
}
