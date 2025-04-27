package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"

	db "ai-docqa-backend/generated/prisma-client"
	"ai-docqa-backend/internal/routes"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️  No .env file found, continuing...")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	app := fiber.New()
	app.Use(cors.New())
	
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	client := db.NewClient()
	if err := client.Prisma.Connect(); err != nil {
		log.Printf("⚠️ Database connection error: %v", err)
	} else {
		defer client.Prisma.Disconnect()
		routes.SetupRoutes(app, client)
	}

	log.Println("🌍 Server starting on port", port)
	if err := app.Listen("0.0.0.0:" + port); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}