package api

import (
	"fmt"

	"github.com/Long-Plan/longplan-api/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/samber/lo"
)

const API_PREFIX = "/api"

func InitAPI(app *fiber.App) {
	config := config.Config.Application
	domain := config.Domain
	if lo.IsEmpty(domain) || domain == "localhost" {
		domain = "localhost:3000"
	}
	app.Use(cors.New(cors.Config{AllowOrigins: fmt.Sprintf("https://%v, http://%v", domain, domain), AllowCredentials: true}))
	app.Use(logger.New())
	app.Use(recover.New())

	router := app.Group(API_PREFIX)
	bindFirstVersionRouter(router)
}
