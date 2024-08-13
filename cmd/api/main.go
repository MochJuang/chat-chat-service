package main

import (
	"chat-service/internal/config"
	grpcserver "chat-service/internal/delivery/grpc/server"
	"chat-service/internal/delivery/http/route"
	"chat-service/internal/repository/postgresql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"log"
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

	go grpcserver.SetupGrpc(cfg)

	// Start server
	log.Fatal(app.Listen(cfg.ServerAddress))
}
