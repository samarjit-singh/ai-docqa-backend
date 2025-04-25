package handlers

import (
	db "ai-docqa-backend/generated/prisma-client"
	"ai-docqa-backend/internal/models"
	"ai-docqa-backend/internal/services"

	"github.com/gofiber/fiber/v2"
)

func HandleQuery(client *db.PrismaClient) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var body models.QueryRequest
		if err := c.BodyParser(&body); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "invalid request"})
		}

		answer, err := services.ProcessQuery(client, body.Document, body.Question)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "AI failed"})
		}

		return c.JSON(fiber.Map{
			"document": body.Document,
			"question": body.Question,
			"answer":   answer,
		})
	}
}

func HandleHistory(client *db.PrismaClient) fiber.Handler {
	return func(c *fiber.Ctx) error {
		history, err := services.GetQueryHistory(client)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "could not fetch history"})
		}
		return c.JSON(history)
	}
}

