package api

import (
	"github.com/Long-Plan/longplan-api/infrastructure"
	"github.com/Long-Plan/longplan-api/internal/adaptor/handler"
	"github.com/Long-Plan/longplan-api/internal/adaptor/repo"
	"github.com/Long-Plan/longplan-api/internal/core/service"
	"github.com/gofiber/fiber/v2"
)

const MAJOR_PREFIX = "/majors"

func bindMajorRouter(router fiber.Router) {
	major := router.Group(MAJOR_PREFIX)
	repo := repo.NewSysMajorRepo(infrastructure.DB)
	service := service.NewMajorService(repo)
	handler := handler.NewMajorHandler(service)

	major.Get("", handler.GetAll)
}
