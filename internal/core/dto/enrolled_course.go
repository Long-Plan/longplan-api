package dto

type EnrolledCourse struct {
	CourseNo string `json:"course_no"`
	Credit   string `json:"credit"`
	Grade    string `json:"grade"`
}

type MappingEnrolledCourse struct {
	Year     string           `json:"year"`
	Semester string           `json:"semester"`
	Courses  []EnrolledCourse `json:"courses"`
}
