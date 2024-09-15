package api

import (
	"fmt"
	"os"

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
	mode := os.Getenv("mode")
	origins := fmt.Sprintf("https://%v, http://%v", domain, domain)

	switch {
	case lo.IsEmpty(domain) || mode == "local":
		domain = "localhost:3000"
		origins = fmt.Sprintf("https://%v, http://%v", domain, domain)
	case mode == "dev":
		origins = fmt.Sprintf("https://localhost:3000, http://localhost:3000, https://%v:8080, http://%v:8080", domain, domain)
	}

	app.Use(cors.New(cors.Config{
		AllowOrigins:     origins,
		AllowCredentials: true,
	}))

	app.Use(logger.New())
	app.Use(recover.New())

	router := app.Group(API_PREFIX)
	bindFirstVersionRouter(router)
}
