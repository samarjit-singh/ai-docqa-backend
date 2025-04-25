package main

import (
	"context"
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

	// Test GET
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("🚀 Hello from Go Fiber!")
	})

	// POST endpoint to create a message
	app.Post("/message", func(c *fiber.Ctx) error {
		type RequestBody struct {
			Content string `json:"content"`
		}

		var body RequestBody
		if err := c.BodyParser(&body); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "cannot parse JSON",
			})
		}

		created, err := client.Message.CreateOne(
			db.Message.Content.Set(body.Content),
		).Exec(context.Background())

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "failed to save message",
			})
		}

		return c.JSON(created)
	})

	log.Println("🌍 Server running on port", port)
	log.Fatal(app.Listen(":" + port))
}
