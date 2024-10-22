package handler

import (
	"log"
	"strconv"

	"github.com/Long-Plan/longplan-api/internal/core/domain"
	"github.com/Long-Plan/longplan-api/internal/core/dto"
	"github.com/Long-Plan/longplan-api/pkg/errors"
	"github.com/Long-Plan/longplan-api/pkg/lodash"
	"github.com/gofiber/fiber/v2"
)

type studentHandler struct {
	serv domain.StudentService
}

func NewStudentHandler(serv domain.StudentService) *studentHandler {
	return &studentHandler{
		serv: serv,
	}
}

func (h studentHandler) UpdateTerm(c *fiber.Ctx) error {
	studentCodeStr, ok := c.Locals("student_code").(string)
	if !ok {
		log.Println("student_code is not a string")
		return lodash.ResponseBadRequest(c)
	}
	studentCode, err := strconv.Atoi(studentCodeStr)
	if err != nil {
		log.Println(err)
		return lodash.ResponseBadRequest(c)
	}

	student, err := h.serv.GetByStudentCode(studentCode)
	if err != nil {
		return lodash.ResponseError(c, errors.NewInternalError(err.Error()))
	}

	student.IsTermAccepted = true

	if err := h.serv.Save(*student); err != nil {
		return lodash.ResponseError(c, errors.NewInternalError(err.Error()))
	}

	return lodash.ResponseOK(c, student)
}

func (h studentHandler) UpdateMajor(c *fiber.Ctx) error {
	studentCodeStr, ok := c.Locals("student_code").(string)
	if !ok {
		log.Println("student_code is not a string")
		return lodash.ResponseBadRequest(c)
	}
	studentCode, err := strconv.Atoi(studentCodeStr)
	if err != nil {
		log.Println(err)
		return lodash.ResponseBadRequest(c)
	}

	var studentMajorUpdateDto dto.StudentMajorUpdateDto
	if err := c.BodyParser(&studentMajorUpdateDto); err != nil {
		return lodash.ResponseError(c, errors.NewBadRequestError(err.Error()))
	}

	student, err := h.serv.GetByStudentCode(studentCode)
	if err != nil {
		return lodash.ResponseError(c, errors.NewInternalError(err.Error()))
	}

	student.MajorID = &studentMajorUpdateDto.MajorID

	if err := h.serv.Save(*student); err != nil {
		return lodash.ResponseError(c, errors.NewInternalError(err.Error()))
	}

	return lodash.ResponseOK(c, student)
}

func (h studentHandler) UpdateCurriculum(c *fiber.Ctx) error {
	studentCodeStr, ok := c.Locals("student_code").(string)
	if !ok {
		log.Println("student_code is not a string")
		return lodash.ResponseBadRequest(c)
	}
	studentCode, err := strconv.Atoi(studentCodeStr)
	if err != nil {
		log.Println(err)
		return lodash.ResponseBadRequest(c)
	}

	var StudentCurriculumUpdateDto dto.StudentCurriculumUpdateDto
	if err := c.BodyParser(&StudentCurriculumUpdateDto); err != nil {
		return lodash.ResponseError(c, errors.NewBadRequestError(err.Error()))
	}

	student, err := h.serv.GetByStudentCode(studentCode)
	if err != nil {
		return lodash.ResponseError(c, errors.NewInternalError(err.Error()))
	}

	student.StudentCurriculumID = &StudentCurriculumUpdateDto.StudentCurriculumID

	if err := h.serv.Save(*student); err != nil {
		return lodash.ResponseError(c, errors.NewInternalError(err.Error()))
	}

	return lodash.ResponseOK(c, student)
}
