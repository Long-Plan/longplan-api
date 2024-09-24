package handler

import (
	"regexp"
	"strconv"

	"github.com/Long-Plan/longplan-api/internal/core/domain"
	"github.com/Long-Plan/longplan-api/pkg/lodash"
	"github.com/gofiber/fiber/v2"
)

type enrolledCourseHandler struct {
	serv domain.EnrolledCourseService
}

func NewEnrolledCourseHandler(serv domain.EnrolledCourseService) *enrolledCourseHandler {
	return &enrolledCourseHandler{serv}
}

func (h *enrolledCourseHandler) GetEnrolledCoursesByStudentCode(c *fiber.Ctx) error {
	studentCode := c.Locals("student_code").(string)
	studentIDRegex := regexp.MustCompile(`^\d{9}$`)

	if studentCode == "" {
		return lodash.ResponseBadRequest(c)
	}
	if !studentIDRegex.MatchString(studentCode) {
		return lodash.ResponseBadRequest(c)
	}

	studentCodeInt, err := strconv.Atoi(studentCode)
	if err != nil {
		return lodash.ResponseBadRequest(c)
	}

	mappings, err := h.serv.GetEnrolledCoursesByStudentCode(studentCodeInt)
	if err != nil {
		return lodash.ResponseError(c, err)
	}

	return lodash.ResponseOK(c, mappings)
}

func (h *enrolledCourseHandler) GetEnrolledCoursesByStudentCodeParam(c *fiber.Ctx) error {
	studentCode := c.Params("student_code")
	studentIDRegex := regexp.MustCompile(`^\d{9}$`)

	if studentCode == "" {
		return lodash.ResponseBadRequest(c)
	}
	if !studentIDRegex.MatchString(studentCode) {
		return lodash.ResponseBadRequest(c)
	}

	studentCodeInt, err := strconv.Atoi(studentCode)
	if err != nil {
		return lodash.ResponseBadRequest(c)
	}

	mappings, err := h.serv.GetEnrolledCoursesByStudentCode(studentCodeInt)
	if err != nil {
		return lodash.ResponseError(c, err)
	}

	return lodash.ResponseOK(c, mappings)
}
