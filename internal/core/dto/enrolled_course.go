package dto

type EnrolledCourse struct {
	CourseNo string
	Credit   string
	Grade    string
}

type MappingEnrolledCourse struct {
	Year     string
	Semester string
	Courses  []EnrolledCourse
}
