package routes

import "github.com/gofiber/fiber/v2"

type LoginForm struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SettingsForm struct {
	Amount   int  `json:"amount"`
	SearchOn bool `json:"searchOn"`
	AddNew   bool `json:"addNew"`
}

func SetRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Welcome to Halim Search API",
		})
	})
}
