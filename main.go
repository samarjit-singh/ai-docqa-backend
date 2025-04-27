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
	err := godotenv.Load()
	if err != nil {
		log.Fatal("🚀 Could not load environment variables: ", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	client := db.NewClient()
	if err := client.Prisma.Connect(); err != nil {
		log.Fatal(err)
	}
	defer client.Prisma.Disconnect()

	app := fiber.New()
	app.Use(cors.New())
	routes.SetupRoutes(app, client)

	log.Println("🌍 Server running on port", port)
	log.Fatal(app.Listen(":" + port))
}
