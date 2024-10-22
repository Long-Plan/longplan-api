package api

import (
	"github.com/Long-Plan/longplan-api/infrastructure"
	"github.com/Long-Plan/longplan-api/internal/adaptor/handler"
	"github.com/Long-Plan/longplan-api/internal/adaptor/repo"
	"github.com/Long-Plan/longplan-api/internal/core/service"
	middlewares "github.com/Long-Plan/longplan-api/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

const STUDENT_CURRICULUM_PREFIX = "/student-curricula"

func bindStudentCurriculumRouter(router fiber.Router) {
	studentCurriculum := router.Group(STUDENT_CURRICULUM_PREFIX)
	studentCurriculumRepo := repo.NewStudentCurriculumRepo(infrastructure.DB)
	studentCurriculumCourseRepo := repo.NewStudentCurriculumCourseRepo(infrastructure.DB)
	studentCurriculumQuestionAnswerRepo := repo.NewStudentCurriculumQuestionAnswerRepo(infrastructure.DB)
	sysCategoryCourseRepo := repo.NewSysCategoryCourseRepo(infrastructure.DB)
	studentCurriculumService := service.NewStudentCurriculumService(
		studentCurriculumRepo,
		studentCurriculumCourseRepo,
		studentCurriculumQuestionAnswerRepo,
		sysCategoryCourseRepo,
	)
	hdl := handler.NewStudentCurriculumHandler(studentCurriculumService)

	studentCurriculum.Get("/student", middlewares.AuthMiddleware(), hdl.GetByStudentCode)
	studentCurriculum.Get("/:studentCurriculumID", hdl.GetByStudentCurriculumID)
	studentCurriculum.Post("", hdl.Create)
	studentCurriculum.Put("", hdl.Update)
	studentCurriculum.Delete("/:studentCurriculumID", hdl.Delete)
	studentCurriculum.Put("/:studentCurriculumID/courses", hdl.UpdateCourses)
	studentCurriculum.Put("/:studentCurriculumID/questions", hdl.UpdateQuestionAnswers)
}
