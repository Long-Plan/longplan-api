package handler

import (
	"log"
	"strconv"

	"github.com/Long-Plan/longplan-api/internal/core/domain"
	"github.com/Long-Plan/longplan-api/internal/core/dto"
	"github.com/Long-Plan/longplan-api/internal/core/model"
	"github.com/Long-Plan/longplan-api/pkg/errors"
	"github.com/Long-Plan/longplan-api/pkg/lodash"
	"github.com/gofiber/fiber/v2"
)

type studentCurriculumHandler struct {
	studentCurriculumService domain.StudentCurriculumService
}

func NewStudentCurriculumHandler(studentCurriculumService domain.StudentCurriculumService) *studentCurriculumHandler {
	return &studentCurriculumHandler{studentCurriculumService}
}

func (h *studentCurriculumHandler) GetByStudentCode(c *fiber.Ctx) error {
	majorId, err := c.ParamsInt("majorId", 0)
	if err != nil {
		return lodash.ResponseError(c, errors.NewBadRequestError(err.Error()))
	}

	studentCodeStr, ok := c.Locals("student_code").(string)
	if !ok {
		log.Println("student_code is not a string")
		return lodash.ResponseBadRequest(c)
	}
	studentCode, err := strconv.Atoi(studentCodeStr)
	if err != nil {
		return lodash.ResponseBadRequest(c)
	}

	studentCurricula, err := h.studentCurriculumService.GetByStudentCode(studentCode, majorId)
	if err != nil {
		return lodash.ResponseError(c, errors.NewInternalError(err.Error()))
	}

	return lodash.ResponseOK(c, studentCurricula)
}

func (h *studentCurriculumHandler) GetByStudentCurriculumID(c *fiber.Ctx) error {
	studentCurriculumID, err := c.ParamsInt("studentCurriculumID", 0)
	if err != nil {
		return lodash.ResponseError(c, errors.NewBadRequestError(err.Error()))
	}

	studentCurriculum, err := h.studentCurriculumService.GetByStudentCurriculumID(studentCurriculumID)
	if err != nil {
		return lodash.ResponseError(c, errors.NewInternalError(err.Error()))
	}

	return lodash.ResponseOK(c, studentCurriculum)
}

func (h *studentCurriculumHandler) Create(c *fiber.Ctx) error {
	var studentCurriculum dto.StudentCurriculumCreate
	if err := c.BodyParser(&studentCurriculum); err != nil {
		return lodash.ResponseError(c, errors.NewBadRequestError(err.Error()))
	}

	studentCurriculumID, err := h.studentCurriculumService.Create(studentCurriculum)
	if err != nil {
		return lodash.ResponseError(c, errors.NewInternalError(err.Error()))
	}

	return lodash.ResponseOK(c, *studentCurriculumID)
}

func (h *studentCurriculumHandler) Update(c *fiber.Ctx) error {
	var studentCurriculum model.StudentCurriculum
	if err := c.BodyParser(&studentCurriculum); err != nil {
		return lodash.ResponseError(c, errors.NewBadRequestError(err.Error()))
	}

	if err := h.studentCurriculumService.Update(studentCurriculum); err != nil {
		return lodash.ResponseError(c, errors.NewInternalError(err.Error()))
	}

	return lodash.ResponseOK(c, studentCurriculum)
}

func (h *studentCurriculumHandler) Delete(c *fiber.Ctx) error {
	studentCurriculumID, err := c.ParamsInt("studentCurriculumID", 0)
	if err != nil {
		return lodash.ResponseError(c, errors.NewBadRequestError(err.Error()))
	}

	if err := h.studentCurriculumService.Delete(studentCurriculumID); err != nil {
		return lodash.ResponseError(c, errors.NewInternalError(err.Error()))
	}

	return lodash.ResponseOK(c, nil)
}

func (h *studentCurriculumHandler) UpdateCourses(c *fiber.Ctx) error {
	studentCurriculumID, err := c.ParamsInt("studentCurriculumID", 0)
	if err != nil {
		return lodash.ResponseError(c, errors.NewBadRequestError(err.Error()))
	}

	var courses []model.StudentCurriculumCourse
	if err := c.BodyParser(&courses); err != nil {
		return lodash.ResponseError(c, errors.NewBadRequestError(err.Error()))
	}

	if err := h.studentCurriculumService.UpdateCourses(studentCurriculumID, courses); err != nil {
		return lodash.ResponseError(c, errors.NewInternalError(err.Error()))
	}

	return lodash.ResponseOK(c, courses)
}

func (h *studentCurriculumHandler) UpdateQuestionAnswers(c *fiber.Ctx) error {
	studentCurriculumID, err := c.ParamsInt("studentCurriculumID", 0)
	if err != nil {
		return lodash.ResponseError(c, errors.NewBadRequestError(err.Error()))
	}

	var questions []dto.StudentCurriculumQuestionAnswer
	if err := c.BodyParser(&questions); err != nil {
		return lodash.ResponseError(c, errors.NewBadRequestError(err.Error()))
	}

	if err := h.studentCurriculumService.UpdateQuestionAnswers(studentCurriculumID, questions); err != nil {
		return lodash.ResponseError(c, errors.NewInternalError(err.Error()))
	}

	return lodash.ResponseOK(c, questions)
}
