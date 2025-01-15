package main

import (
	"fmt"
	"time"

	"github.com/feynmaz/fiberg/configs"
	"github.com/feynmaz/fiberg/database"
	"github.com/feynmaz/fiberg/handlers"
	"github.com/feynmaz/fiberg/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"go.uber.org/zap"
)

func main() {
	config := configs.GetDefault()
	utils.InitLogger(config.Debug)
	db := database.Connect()

	app := fiber.New()

	app.Use(recover.New())
	app.Use(requestid.New())
	app.Use(logger.New(logger.Config{
		Format:     "{\"level\":\"info\",\"time\":\"${time}\",\"request_id\":\"${locals:requestid}\",\"status\":${status},\"method\":\"${method}\",\"path\":\"${path}\"}\n",
		TimeFormat: time.RFC3339,
	}))

	handler := handlers.NewHandler(db)

	app.Get("/healthcheck", handler.HealthCheck)

	v1 := app.Group("/api/v1")
	v1.Get("/users", handler.UserList)
	v1.Post("/users", handler.UserCreate)

	app.Static("/", "./static/public")

	app.Use(handlers.NotFound)

	zap.L().With(
		zap.Error(
			app.Listen(fmt.Sprintf(":%d", config.Port)),
		),
	).Fatal("")
}
