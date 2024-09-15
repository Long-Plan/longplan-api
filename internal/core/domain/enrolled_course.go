package domain

import "github.com/Long-Plan/longplan-api/internal/core/dto"

type EnrolledCourseService interface {
	GetEnrolledCoursesByStudentID(studentID string) ([]dto.MappingEnrolledCourse, error)
}
