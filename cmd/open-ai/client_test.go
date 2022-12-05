package open_ai

import (
	"context"
	"fmt"
	gogpt "github.com/sashabaranov/go-gpt3"
	"log"
	"os"
	"testing"
)

func TestNewClient_BuildGoProgram(t *testing.T) {
	// Create a GPT3 client
	c := gogpt.NewClient(os.Getenv("API_KEY"))
	prompt := "Build a Hello World program in Golang"
	req := gogpt.CompletionRequest{
		Model:            "text-davinci-003",
		Prompt:           prompt,
		Temperature:      0.7,
		MaxTokens:        2000,
		TopP:             1,
		FrequencyPenalty: 0,
		PresencePenalty:  0,
	}
	resp, err := c.CreateCompletion(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}
	choices := resp.Choices
	fmt.Println(choices[0].Text)
}
