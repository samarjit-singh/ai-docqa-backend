package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"

	db "ai-docqa-backend/generated/prisma-client"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("🚀 Could not load environment variables: ", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	client := db.NewClient()
	if err := client.Prisma.Connect(); err != nil {
		log.Fatal(err)
	}
	defer client.Prisma.Disconnect()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("🚀 Hello from Go Fiber!")
	})

	log.Println("🌍 Server running on port", port)
	log.Fatal(app.Listen(":" + port))
}
