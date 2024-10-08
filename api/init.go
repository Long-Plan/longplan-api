package api

import (
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
	origin := config.ClientOrigin

	if lo.IsEmpty(origin) {
		origin = "*"
	}

	app.Use(cors.New(cors.Config{
		AllowOrigins: origin,
	}))

	app.Use(logger.New())
	app.Use(recover.New())

	router := app.Group(API_PREFIX)
	bindFirstVersionRouter(router)
}
