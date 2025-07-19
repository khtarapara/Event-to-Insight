package ai

import (
	"context"
	"fmt"
	"os"

	openai "github.com/sashabaranov/go-openai"
)

type OpenAIClient struct {
	client *openai.Client
}

func NewOpenAIClient() *OpenAIClient {
	config := openai.DefaultConfig(os.Getenv("OPENROUTER_API_KEY"))
	config.BaseURL = "https://openrouter.ai/api/v1"

	return &OpenAIClient{
		client: openai.NewClientWithConfig(config),
	}
}

func (c *OpenAIClient) ClassifyIncident(title, description string) (string, string, error) {
	prompt := MakePromptForEventClassification(title, description)

	resp, err := c.client.CreateChatCompletion(context.Background(), openai.ChatCompletionRequest{
		Model: "google/gemma-3n-e2b-it:free",
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: prompt,
			},
		},
	})

	if err != nil {
		fmt.Println("Error creating chat completion:", err)
		return "", "", err
	}

	output := resp.Choices[0].Message.Content

	severity, category := ParseAIClassification(output)

	return severity, category, nil
}
