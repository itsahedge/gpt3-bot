package gpt3_bot

import (
	"fmt"
	"github.com/joho/godotenv"
	gogpt "github.com/sashabaranov/go-gpt3"
	"log"
	"os"
)

type GptClient struct {
	client *gogpt.Client
}

func NewClient() *GptClient {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	API_KEY := os.Getenv("API_KEY")
	gpt := gogpt.NewClient(API_KEY)
	gptClient := &GptClient{
		client: gpt,
	}
	return gptClient
}

func (g *GptClient) WriteToFile(content string) error {
	f, err := os.Create("generated.go")
	if err != nil {
		return err
	}
	defer f.Close()
	_, err2 := f.WriteString(content)
	if err2 != nil {
		return err
	}
	fmt.Println("done")
	return nil
}
