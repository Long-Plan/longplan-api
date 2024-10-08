package handler

import (
	"github.com/Long-Plan/longplan-api/internal/core/domain"
	"github.com/Long-Plan/longplan-api/pkg/errors"
	"github.com/Long-Plan/longplan-api/pkg/lodash"
	"github.com/gofiber/fiber/v2"
)

type categoryHandler struct {
	categoryService domain.CategoryService
}

func NewCategoryHandler(categoryService domain.CategoryService) *categoryHandler {
	return &categoryHandler{categoryService}
}

func (h *categoryHandler) GetByCurriculumID(c *fiber.Ctx) error {
	curriculumID, err := c.ParamsInt("curriculumId", 0)
	if err != nil {
		return lodash.ResponseError(c, errors.NewBadRequestError(err.Error()))
	}

	category, err := h.categoryService.GetByCurriculumID(curriculumID)
	if err != nil {
		return lodash.ResponseError(c, errors.NewInternalError(err.Error()))
	}

	return lodash.ResponseOK(c, category)
}

func (h *categoryHandler) GetTypes(c *fiber.Ctx) error {
	types, err := h.categoryService.GetTypes()
	if err != nil {
		return lodash.ResponseError(c, errors.NewInternalError(err.Error()))
	}

	return lodash.ResponseOK(c, types)
}
