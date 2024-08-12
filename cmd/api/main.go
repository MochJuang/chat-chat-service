package main

import (
	"chat-service/internal/config"
	"chat-service/internal/delivery/http/route"
	"chat-service/internal/repository/postgresql"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Could not load config: %v", err)
	}

	// Initialize database
	db, err := postgresql.NewConnector(cfg)
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}
	cfg.DB = db

	app := fiber.New()

	app.Use(logger.New())

	// Setup routes
	route.SetupRoutes(app, cfg)

	// Setup error handler middleware

	// Start server
	log.Fatal(app.Listen(cfg.ServerAddress))
}
