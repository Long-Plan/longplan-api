package handler

import (
	"github.com/Long-Plan/longplan-api/internal/core/domain"
	"github.com/Long-Plan/longplan-api/pkg/errors"
	"github.com/Long-Plan/longplan-api/pkg/lodash"
	"github.com/gofiber/fiber/v2"
)

type curriculumHandler struct {
	serv domain.CurriculumService
}

func NewCurriculumHandler(serv domain.CurriculumService) *curriculumHandler {
	return &curriculumHandler{
		serv: serv,
	}
}

func (h *curriculumHandler) All(c *fiber.Ctx) error {
	curriculums, err := h.serv.GetAll()
	if err != nil {
		return lodash.ResponseError(c, errors.NewInternalError(err.Error()))
	}

	return lodash.ResponseOK(c, curriculums)
}

func (h *curriculumHandler) AllByMajorID(c *fiber.Ctx) error {
	majorID, err := c.ParamsInt("majorId", 0)
	if err != nil {
		return lodash.ResponseError(c, errors.NewBadRequestError(err.Error()))
	}

	curriculums, err := h.serv.GetAllByMajorID(majorID)
	if err != nil {
		return lodash.ResponseError(c, errors.NewInternalError(err.Error()))
	}

	return lodash.ResponseOK(c, curriculums)
}

func (h *curriculumHandler) GetByID(c *fiber.Ctx) error {
	curriculumID, err := c.ParamsInt("curriculumId", 0)
	if err != nil {
		return lodash.ResponseError(c, errors.NewBadRequestError(err.Error()))
	}

	curriculum, err := h.serv.GetByID(curriculumID)
	if err != nil {
		return lodash.ResponseError(c, errors.NewInternalError(err.Error()))
	}

	return lodash.ResponseOK(c, curriculum)
}
