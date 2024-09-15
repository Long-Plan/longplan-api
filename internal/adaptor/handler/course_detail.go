package handler

import (
	"github.com/Long-Plan/longplan-api/internal/core/domain"
	"github.com/Long-Plan/longplan-api/pkg/lodash"
	"github.com/gofiber/fiber/v2"
)

type courseDetailHandler struct {
	serv domain.CourseDetailService
}

func NewCourseDetailHandler(serv domain.CourseDetailService) *courseDetailHandler {
	return &courseDetailHandler{
		serv: serv,
	}
}

func (h *courseDetailHandler) GetAll(c *fiber.Ctx) error {
	courseDetails, err := h.serv.GetAll()
	if err != nil {
		return lodash.ResponseError(c, err)
	}

	return lodash.ResponseOK(c, courseDetails)
}
