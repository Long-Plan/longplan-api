package api

import (
	"github.com/Long-Plan/longplan-api/infrastructure"
	"github.com/Long-Plan/longplan-api/internal/adaptor/handler"
	"github.com/Long-Plan/longplan-api/internal/adaptor/repo"
	"github.com/Long-Plan/longplan-api/internal/core/service"
	"github.com/gofiber/fiber/v2"
)

const CATEGORY_PREFIX = "/categories"

func bindCategoryRouter(router fiber.Router) {
	category := router.Group(CATEGORY_PREFIX)
	categoryRepo := repo.NewSysCategoryRepo(infrastructure.DB)
	categoryTypeRepo := repo.NewSysCategoryTypeRepo(infrastructure.DB)
	categoryRequirementRepo := repo.NewSysCategoryRequirementRepo(infrastructure.DB)
	categoryRelationshipRepo := repo.NewSysCategoryRelationshipRepo(infrastructure.DB)
	categoryCourseRepo := repo.NewSysCategoryCourseRepo(infrastructure.DB)
	categoryCourseRequisiteRepo := repo.NewSysCategoryCourseRequisiteRepo(infrastructure.DB)
	courseDetailRepo := repo.NewSysCourseDetailRepo(infrastructure.DB)
	categoryService := service.NewCategoryService(
		categoryRepo,
		categoryTypeRepo,
		categoryRequirementRepo,
		categoryRelationshipRepo,
		categoryCourseRepo,
		categoryCourseRequisiteRepo,
		courseDetailRepo,
	)
	handler := handler.NewCategoryHandler(categoryService)

	category.Get("/types", handler.GetTypes)
	category.Get(":curriculumId", handler.GetByCurriculumID)
}
