package services

import (
	"context"

	db "ai-docqa-backend/generated/prisma-client"
	"ai-docqa-backend/internal/utils"
)


func ProcessQuery(client *db.PrismaClient, document, question string) (string, error) {
	answer, err := utils.AskGemini(document, question)
	if err != nil {
		return "", err
	}

	_, err = client.Query.CreateOne(
		db.Query.Document.Set(document),
		db.Query.Question.Set(question),
		db.Query.Answer.Set(answer),
	).Exec(context.Background())

	return answer, err
}

func GetQueryHistory(client *db.PrismaClient) ([]db.QueryModel, error) {
	return client.Query.FindMany().OrderBy(
		db.Query.CreatedAt.Order(db.SortOrderDesc),
	).Exec(context.Background())
}