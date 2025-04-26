package handlers

import (
	db "ai-docqa-backend/generated/prisma-client"
	"ai-docqa-backend/internal/services"
	"ai-docqa-backend/internal/utils"
	"fmt"
	"io"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
)

func HandleQuery(client *db.PrismaClient) fiber.Handler {
	return func(c *fiber.Ctx) error {
		question := c.FormValue("question")
		if question == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "missing 'question' field in form data"})
		}

		fileHeader, err := c.FormFile("document")
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "missing 'document' file in form data"})
		}

		file, err := fileHeader.Open()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to open uploaded file"})
		}
		defer file.Close()

		fileBytes, err := io.ReadAll(file)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to read file content"})
		}

		var documentContent string
		fileExt := filepath.Ext(fileHeader.Filename)

		switch fileExt {
		case ".pdf":
			documentContent, err = utils.ExtractPDFContent(fileBytes)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"error": fmt.Sprintf("failed to extract PDF content: %v", err),
				})
			}
		case ".txt":
			documentContent = string(fileBytes)
		default:
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": fmt.Sprintf("unsupported file type: %s", fileExt),
			})
		}

		if documentContent == "" {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "failed to extract content from file or file was empty",
			})
		}

		answer, err := services.ProcessQuery(client, documentContent, question)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "AI processing failed"})
		}

		return c.JSON(fiber.Map{
			"filename": fileHeader.Filename,
			"question": question,
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