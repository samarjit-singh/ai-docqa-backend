package models

type QueryRequest struct {
	Document string `json:"document"`
	Question string `json:"question"`
}
