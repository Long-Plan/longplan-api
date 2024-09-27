package handler

import (
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/Long-Plan/longplan-api/config"
	"github.com/Long-Plan/longplan-api/internal/core/domain"
	"github.com/Long-Plan/longplan-api/internal/core/dto"
	"github.com/Long-Plan/longplan-api/internal/core/model"
	"github.com/Long-Plan/longplan-api/pkg/errors"
	"github.com/Long-Plan/longplan-api/pkg/lodash"
	"github.com/Long-Plan/longplan-api/pkg/oauth"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/samber/lo"
)

type oauthHandler struct {
	accountService domain.AccountService
	studentService domain.StudentService
}

func NewOauthHandler(accountService domain.AccountService) *oauthHandler {
	return &oauthHandler{
		accountService: accountService,
	}
}

func (h oauthHandler) SignIn(c *fiber.Ctx) error {
	config := config.Config.Application
	origin := c.Get("Origin")
	var isLocalOrigin bool
	if strings.Contains(origin, "localhost") {
		isLocalOrigin = true
	}
	code := c.Query("code", "")
	if lo.IsEmpty(code) {
		return lodash.ResponseBadRequest(c)
	}
	user, err := oauth.CmuOauthValidation(code, isLocalOrigin)
	if err != nil {
		return lodash.ResponseError(c, errors.NewStatusBadGatewayError(err.Error()))
	}

	accountModel := model.Account{
		CMUITAccount: user.Cmuitaccount,
		Prename:      "",
		Firstname:    user.FirstnameEN,
		Lastname:     user.LastnameEN,
		AccountType:  user.ItaccounttypeID,
		Organization: user.OrganizationNameEN,
	}

	err = h.accountService.Save(accountModel)
	if err != nil {
		return lodash.ResponseError(c, errors.NewInternalError(err.Error()))
	}

	if user.ItaccounttypeID == "StdAcc" {
		code, err := strconv.Atoi(user.StudentID)
		if err != nil {
			return lodash.ResponseError(c, errors.NewInternalError(err.Error()))
		}
		studentModel := model.Student{
			Code: code,
		}
		err = h.studentService.Save(studentModel)
		if err != nil {
			return lodash.ResponseError(c, errors.NewInternalError(err.Error()))
		}
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, &oauth.UserClaims{
		User: *user,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)),
		},
	})

	token, err := claims.SignedString([]byte(config.Secret))
	if err != nil {
		log.Print(err)
		return lodash.ResponseError(c, errors.NewInternalError(err.Error()))
	}

	return lodash.ResponseOK(c, token)
}

func (h oauthHandler) GetUser(c *fiber.Ctx) error {
	cmuitaccount := c.Locals("cmuitaccount").(string)
	user, err := h.accountService.GetByCMUITAccount(cmuitaccount)
	if err != nil {
		return lodash.ResponseError(c, errors.NewInternalError(err.Error()))
	}

	var student *model.Student
	if user.AccountType == "StdAcc" {
		studentCode := c.Locals("student_code").(string)
		code, err := strconv.Atoi(studentCode)
		if err != nil {
			return lodash.ResponseError(c, errors.NewInternalError(err.Error()))
		}
		student, err = h.studentService.GetByStudentCode(code)
		if err != nil {
			return lodash.ResponseError(c, errors.NewInternalError(err.Error()))
		}
	}

	return lodash.ResponseOK(c, dto.Account{
		UserData:    *user,
		StudentData: student,
	})
}

func (h oauthHandler) Logout(c *fiber.Ctx) error {
	return lodash.ResponseNoContent(c, nil)
}
