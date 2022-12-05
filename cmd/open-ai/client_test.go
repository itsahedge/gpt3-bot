package open_ai

import (
	"context"
	"testing"
)

func TestNewClient_BuildGoProgram(t *testing.T) {
	// Create a GPT3 client
	gpt := NewClient()
	prompt := "Build a program in Golang that gets the Price of Ethereum using the Coingecko API and post the response as a message to a Discord channel using github.com/bwmarrin/discordgo."

	resp, err := gpt.Codegen(context.Background(), prompt)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp)

	gpt.WriteToFile(resp)
}
