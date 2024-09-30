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

func (h studentHandler) Update(c *fiber.Ctx) error {
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

	var studentUpdateDto dto.StudentUpdateDto
	err = c.BodyParser(&studentUpdateDto)
	if err != nil {
		log.Println(err)
		return lodash.ResponseError(c, errors.NewBadRequestError(err.Error()))
	}

	student, err := h.serv.GetByStudentCode(studentCode)
	if err != nil {
		return lodash.ResponseError(c, errors.NewBadRequestError("student not found"))
	}

	student.IsTermAccepted = studentUpdateDto.IsTermAccepted
	student.MajorID = studentUpdateDto.MajorID

	if err := h.serv.Save(*student); err != nil {
		return lodash.ResponseError(c, errors.NewInternalError(err.Error()))
	}

	return lodash.ResponseOK(c, student)
}
