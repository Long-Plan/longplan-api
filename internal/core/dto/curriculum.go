package dto

import "time"

type Curriculum struct {
	ID        int       `json:"id"`
	MajorID   int       `json:"major_id"`
	Code      string    `json:"code"`
	NameTH    string    `json:"name_th"`
	NameEN    string    `json:"name_en"`
	ShortName string    `json:"short_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Questions []CurriculumQuestion `json:"curriculum_questions"`
}

type CurriculumQuestion struct {
	ID     int    `json:"id"`
	NameTH string `json:"name_th"`
	NameEN string `json:"name_en"`

	Choices []CurriculumQuestionChoice `json:"choices"`
}

type CurriculumQuestionChoice struct {
	ID       int    `json:"id"`
	Position int    `json:"position"`
	NameTH   string `json:"name_th"`
	NameEN   string `json:"name_en"`
}
