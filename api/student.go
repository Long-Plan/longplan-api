package api

import (
	"github.com/Long-Plan/longplan-api/infrastructure"
	"github.com/Long-Plan/longplan-api/internal/adaptor/handler"
	"github.com/Long-Plan/longplan-api/internal/adaptor/repo"
	"github.com/Long-Plan/longplan-api/internal/core/service"
	"github.com/gofiber/fiber/v2"
)

const STUDENT_PREFIX = "/students"

func bindStudentRouter(router fiber.Router) {
	student := router.Group(STUDENT_PREFIX)
	studentRepo := repo.NewStudentRepo(infrastructure.DB)
	studentService := service.NewStudentService(studentRepo)
	hdl := handler.NewStudentHandler(studentService)
	student.Put("", hdl.Update)
}
