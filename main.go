package main

import (
	"fmt"
	"log"

	"github.com/Long-Plan/longplan-api/api"
	"github.com/Long-Plan/longplan-api/config"
	"github.com/Long-Plan/longplan-api/infrastructure"
	"github.com/Long-Plan/longplan-api/pkg/lodash"
	"github.com/gofiber/fiber/v2"
)

var cfg config.ApplicationConfig

func init() {
	lodash.SetTimeZone("Asia/Bangkok")
	config.InitConfig()
	cfg = config.Config.Application
	infrastructure.InitDB()
}

func main() {
	app := fiber.New()
	api.InitAPI(app)

	addr := getAddress()
	log.Printf("%v started at %v", cfg.Name, cfg.Port)
	if err := app.Listen(addr); err != nil {
		log.Fatal(err)
	}
}

func getAddress() string {
	addr := ":8000"
	if cfg.Port != "" {
		addr = fmt.Sprintf(":%v", cfg.Port)
	}
	return addr
}
