package fiber

import (
	"cafe-pos/internal/interface/http/fiber/user"

	"github.com/gofiber/fiber/v2"
)

// Setup Main routes
func NewRouter(
	app *fiber.App,
	userHandler *user.Handler,
) {
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	app.Get("/docs", func(c *fiber.Ctx) error {
		return c.SendString("Not implemented")
	})

	main := app.Group("/api/v1")

	setupUserRoutes(main, userHandler)
}

// User routes
func setupUserRoutes(group fiber.Router, userHandler *user.Handler) {
	userGroup := group.Group("/users")

	// Routes
	userGroup.Post("/", userHandler.CreateUser)
	userGroup.Get("/:id", userHandler.GetUserByID)
}
