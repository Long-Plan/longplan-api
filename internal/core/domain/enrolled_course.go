package domain

import "github.com/Long-Plan/longplan-api/internal/core/dto"

type EnrolledCourseService interface {
	GetEnrolledCoursesByStudentCode(studentCode int) ([]dto.MappingEnrolledCourse, error)
}
