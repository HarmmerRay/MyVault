package services

import (
	"myvault-backend/pkg/ai"
)

type AIService struct {
	client *ai.OpenAIClient
}

func NewAIService(apiKey string) *AIService {
	return &AIService{
		client: ai.NewOpenAIClient(apiKey),
	}
}

func (s *AIService) GenerateSummary(prompt string) (string, error) {
	return s.client.GenerateSummary(prompt)
}