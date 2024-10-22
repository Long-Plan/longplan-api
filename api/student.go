package api

import (
	"github.com/Long-Plan/longplan-api/infrastructure"
	"github.com/Long-Plan/longplan-api/internal/adaptor/handler"
	"github.com/Long-Plan/longplan-api/internal/adaptor/repo"
	"github.com/Long-Plan/longplan-api/internal/core/service"
	middlewares "github.com/Long-Plan/longplan-api/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

const STUDENT_PREFIX = "/students"

func bindStudentRouter(router fiber.Router) {
	student := router.Group(STUDENT_PREFIX)
	studentRepo := repo.NewStudentRepo(infrastructure.DB)
	studentService := service.NewStudentService(studentRepo)
	hdl := handler.NewStudentHandler(studentService)
	student.Put("/major", middlewares.AuthMiddleware(), hdl.UpdateMajor)
	student.Post("/term", middlewares.AuthMiddleware(), hdl.UpdateTerm)
	student.Put("/curriculum", middlewares.AuthMiddleware(), hdl.UpdateCurriculum)
}
