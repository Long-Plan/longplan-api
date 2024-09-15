package handler

import (
	"github.com/Long-Plan/longplan-api/internal/core/domain"
	"github.com/Long-Plan/longplan-api/pkg/lodash"
	"github.com/gofiber/fiber/v2"
)

type majorHandler struct {
	majorService domain.MajorService
}

func NewMajorHandler(majorService domain.MajorService) *majorHandler {
	return &majorHandler{
		majorService: majorService,
	}
}

func (h *majorHandler) GetAll(c *fiber.Ctx) error {
	majors, err := h.majorService.GetAll()
	if err != nil {
		return lodash.ResponseError(c, err)
	}

	return lodash.ResponseOK(c, majors)
}
