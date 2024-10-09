package main

import (
	"backend/routes"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/joho/godotenv"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	env := godotenv.Load() // Load environment variables
	if env != nil {
		panic("Cannot find environment variables") // If .env file is not found, panic
	}
	port := os.Getenv("PORT") // Get port from environment variables
	if port == "" {
		port = ":4000" // Default port is 4000
	} else {
		port = ":" + port // Add colon to port number
	}
	app := fiber.New(fiber.Config{
		IdleTimeout: 5 * time.Second, // Set idle timeout to 5 seconds
	})

	app.Use(compress.New()) // Use compression middleware

	routes.SetRoutes(app)

	// Start server and listen for shutdown command
	go func() {
		if err := app.Listen(port); err != nil {
			log.Panic(err)
		}
	}()

	c := make(chan os.Signal, 1)                    // Create a channel to listen for signals. 1 is the buffer size
	signal.Notify(c, os.Interrupt, syscall.SIGTERM) // Listen for interrupt signal or SIGTERM

	<-c // Block the main thread until interrupted
	err := app.Shutdown()
	if err != nil {
		return
	} // Shutdown server
	fmt.Println("Shutting down server")
}
