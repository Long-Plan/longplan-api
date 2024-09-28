package middlewares

import (
	"strings"

	"github.com/Long-Plan/longplan-api/config"
	"github.com/Long-Plan/longplan-api/pkg/errors"
	"github.com/Long-Plan/longplan-api/pkg/lodash"
	"github.com/Long-Plan/longplan-api/pkg/oauth"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/samber/lo"
)

func AuthMiddleware() func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		config := config.Config.Application

		invalidToken := errors.NewUnauthorizedError(errors.AuthErr("invalid token").Error())

		token := c.Get("Authorization")
		if lo.IsEmpty(token) {
			return lodash.ResponseError(c, errors.NewUnauthorizedError("empty token"))
		}

		token = strings.Split(token, " ")[1]

		parsedAccessToken, err := jwt.ParseWithClaims(token, &oauth.UserClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.Secret), nil
		})
		if err != nil {
			return lodash.ResponseError(c, invalidToken)
		}
		user := &parsedAccessToken.Claims.(*oauth.UserClaims).User

		c.Locals("cmuitaccount", user.Cmuitaccount)
		c.Locals("student_code", user.StudentID)
		return c.Next()
	}
}
