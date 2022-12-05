package gpt3_bot

import (
	"context"
	"errors"
	gogpt "github.com/sashabaranov/go-gpt3"
)

func (g *GptClient) Codegen(ctx context.Context, prompt string) (string, error) {
	if prompt == "" {
		return "", errors.New("must provided a string prompt")
	}
	req := gogpt.CompletionRequest{
		Model:            "text-davinci-003",
		Prompt:           prompt,
		Temperature:      0.7,
		MaxTokens:        2000,
		TopP:             1,
		FrequencyPenalty: 0,
		PresencePenalty:  0,
	}

	resp, err := g.client.CreateCompletion(ctx, req)
	if err != nil {
		return "", err
	}
	choices := resp.Choices
	//fmt.Println(choices[0].Text)
	raw := choices[0].Text
	return raw, nil
}
