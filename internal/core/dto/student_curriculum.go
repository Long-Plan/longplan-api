package dto

type StudentCurriculum struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	StudentCode  int    `json:"student_code"`
	CurriculumID int    `json:"curriculum_id"`
	IsSystem     bool   `json:"is_system"`

	Courses []StudentCurriculumCourse         `json:"courses"`
	Answers []StudentCurriculumQuestionAnswer `json:"answers"`
}

type StudentCurriculumCreate struct {
	Name         string `json:"name"`
	StudentCode  int    `json:"student_code"`
	CurriculumID int    `json:"curriculum_id"`
	IsSystem     bool   `json:"is_system"`

	Answers []StudentCurriculumQuestionAnswer `json:"answers"`
}

type StudentCurriculumCourse struct {
	ID         int    `json:"id"`
	Semester   int    `json:"semester"`
	Year       int    `json:"year"`
	CourseNo   string `json:"course_no"`
	CategoryID int    `json:"category_id"`
}

type StudentCurriculumQuestionAnswer struct {
	ID         int `json:"id"`
	QuestionID int `json:"question_id"`
	ChoiceID   int `json:"choice_id"`
}
