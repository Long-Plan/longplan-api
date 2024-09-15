package api

import (
	"github.com/Long-Plan/longplan-api/infrastructure"
	"github.com/Long-Plan/longplan-api/internal/adaptor/handler"
	"github.com/Long-Plan/longplan-api/internal/adaptor/repo"
	"github.com/Long-Plan/longplan-api/internal/core/service"
	"github.com/gofiber/fiber/v2"
)

const COURSE_DETAIL_PREFIX = "/course-details"

func bindCourseDetailRouter(router fiber.Router) {
	courseDetail := router.Group(COURSE_DETAIL_PREFIX)
	repo := repo.NewSysCourseDetailRepo(infrastructure.DB)
	service := service.NewCourseDetailService(repo)
	handler := handler.NewCourseDetailHandler(service)

	courseDetail.Get("", handler.GetAll)
}
