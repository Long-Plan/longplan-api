package dto

type StudentMajorUpdateDto struct {
	MajorID int `json:"major_id"`
}

type StudentCurriculumUpdateDto struct {
	StudentCurriculumID int `json:"student_curriculum_id"`
}
