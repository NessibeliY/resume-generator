package service

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	openai "github.com/sashabaranov/go-openai"

	"github.com/NessibeliY/resume-generator/internal/models"
)

type AIService struct {
	client *openai.Client
}

func NewAIService(apiKey string) *AIService {
	return &AIService{
		client: openai.NewClient(apiKey),
	}
}

func (s *AIService) GenerateResume(req models.ResumeRequest) (string, error) {
	prompt := fmt.Sprintf(`Сгенерируй профессиональное резюме на основе следующих данных:
Имя: %s
Роль: %s
Навыки: %s
Опыт: %s`, req.Name, req.Role, req.Skills, req.Experience)

	resp, err := s.client.CreateChatCompletion(context.Background(), openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo,
		Messages: []openai.ChatCompletionMessage{
			{Role: "user", Content: prompt},
		},
	})
	if err != nil {
		return "", errors.Wrap(err, "create chat completion")
	}

	return resp.Choices[0].Message.Content, nil
}
