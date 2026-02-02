package fiber

import (
	"cafe-pos/config"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func NewServer(cfg *config.Config) *fiber.App {
	app := fiber.New(
		fiber.Config{
			AppName: "Pdora POS",
		},
	)

	app.Use(cors.New())

	app.Use(logger.New())

	return app
}
