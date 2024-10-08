package api

import (
	"github.com/Long-Plan/longplan-api/infrastructure"
	"github.com/Long-Plan/longplan-api/internal/adaptor/handler"
	"github.com/Long-Plan/longplan-api/internal/adaptor/repo"
	"github.com/Long-Plan/longplan-api/internal/core/service"
	"github.com/gofiber/fiber/v2"
)

const CURRICULUM_PREFIX = "/curricula"

func bindCurriculumRouter(router fiber.Router) {
	curriculum := router.Group(CURRICULUM_PREFIX)
	curriculumRepo := repo.NewSysCurriculumRepo(infrastructure.DB)
	curriculumQuestionRepo := repo.NewSysCurriculumQuestionRepo(infrastructure.DB)
	curriculumQuestionChoiceRepo := repo.NewSysCurriculumQuestionChoiceRepo(infrastructure.DB)
	categoryCourseRepo := repo.NewSysCategoryCourseRepo(infrastructure.DB)
	courseDetailRepo := repo.NewSysCourseDetailRepo(infrastructure.DB)
	curriculumService := service.NewCurriculumService(curriculumRepo, curriculumQuestionRepo, curriculumQuestionChoiceRepo, categoryCourseRepo, courseDetailRepo)
	curriculumHandler := handler.NewCurriculumHandler(curriculumService)
	curriculum.Get("", curriculumHandler.All)
	curriculum.Get("/major/:majorId", curriculumHandler.AllByMajorID)
	curriculum.Get("/:curriculumId", curriculumHandler.GetByID)
	curriculum.Get("/courses/:curriculumId", curriculumHandler.GetCoursesByCurriculumID)
}
