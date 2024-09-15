package api

import (
	"github.com/Long-Plan/longplan-api/internal/adaptor/handler"
	"github.com/Long-Plan/longplan-api/internal/core/service"
	"github.com/gofiber/fiber/v2"
)

const ENROLLED_COURSE_PREFIX = "/enrolled-courses"

func bindEnrolledCourseRouter(router fiber.Router) {
	enrolledCourse := router.Group(ENROLLED_COURSE_PREFIX)
	serv := service.NewEnrolledCourseService()
	handler := handler.NewEnrolledCourseHandler(serv)

	enrolledCourse.Get(":studentId", handler.GetEnrolledCoursesByStudentID)
}
