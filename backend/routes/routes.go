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

	app.Post("/settings", func(c *fiber.Ctx) error {
		input := new(SettingsForm)
		if err := c.BodyParser(input); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid input",
			})
		}
		// Process the settings...
		return c.JSON(fiber.Map{
			"message":  "Settings updated successfully",
			"settings": input,
		})
	})

	app.Get("/login", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Please provide login credentials",
		})
	})

	app.Post("/login", func(c *fiber.Ctx) error {
		input := new(LoginForm)
		if err := c.BodyParser(input); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid input",
			})
		}
		// Process the login...
		return c.JSON(fiber.Map{
			"message": "Login successful",
		})
	})
}
