package routes

import (
	db "ai-docqa-backend/generated/prisma-client"
	"ai-docqa-backend/internal/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, client *db.PrismaClient) {
	app.Post("/query", handlers.HandleQuery(client))
}
