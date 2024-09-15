package handler

import (
	"regexp"

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

func (h *enrolledCourseHandler) GetEnrolledCoursesByStudentID(c *fiber.Ctx) error {
	studentID := c.Params("studentId")
	studentIDRegex := regexp.MustCompile(`^\d{9}$`)

	if studentID == "" {
		return lodash.ResponseBadRequest(c)
	}
	if !studentIDRegex.MatchString(studentID) {
		return lodash.ResponseBadRequest(c)
	}

	mappings, err := h.serv.GetEnrolledCoursesByStudentID(studentID)
	if err != nil {
		return lodash.ResponseError(c, err)
	}

	return lodash.ResponseOK(c, mappings)
}
