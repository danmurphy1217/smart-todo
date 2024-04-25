package llm

import "github.com/sashabaranov/go-openai"

func NewClient(accessToken string) *openai.Client {
	return openai.NewClient(accessToken)
}
