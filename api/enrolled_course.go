package api

import (
	"github.com/Long-Plan/longplan-api/internal/adaptor/handler"
	"github.com/Long-Plan/longplan-api/internal/core/service"
	middlewares "github.com/Long-Plan/longplan-api/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

const ENROLLED_COURSE_PREFIX = "/enrolled-courses"

func bindEnrolledCourseRouter(router fiber.Router) {
	enrolledCourse := router.Group(ENROLLED_COURSE_PREFIX)
	serv := service.NewEnrolledCourseService()
	handler := handler.NewEnrolledCourseHandler(serv)

	enrolledCourse.Get("", middlewares.AuthMiddleware(), handler.GetEnrolledCoursesByStudentCode)
	enrolledCourse.Get("/:student_code", handler.GetEnrolledCoursesByStudentCodeParam)
}
