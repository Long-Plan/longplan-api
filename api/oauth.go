package api

import (
	"github.com/Long-Plan/longplan-api/infrastructure"
	"github.com/Long-Plan/longplan-api/internal/adaptor/handler"
	"github.com/Long-Plan/longplan-api/internal/adaptor/repo"
	"github.com/Long-Plan/longplan-api/internal/core/service"
	middlewares "github.com/Long-Plan/longplan-api/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

const OAUTH_PREFIX = "/oauth"

func bindOauthRouter(router fiber.Router) {
	oauth := router.Group(OAUTH_PREFIX)

	accountRepo := repo.NewAccountRepo(infrastructure.DB)
	accountTypeRepo := repo.NewAccountTypeRepo(infrastructure.DB)
	organizationRepo := repo.NewOrganizationRepo(infrastructure.DB)
	accountService := service.NewAccountService(accountRepo, accountTypeRepo, organizationRepo)
	hdl := handler.NewOauthHandler(accountService)
	oauth.Get("/me", middlewares.AuthMiddleware(), hdl.GetUser)

	oauth.Post("", hdl.SignIn)
	oauth.Post("/signout", middlewares.AuthMiddleware(), hdl.Logout)
}
