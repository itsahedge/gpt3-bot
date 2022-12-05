package gpt3_bot

import (
	"context"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"os"
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

func TestNewClient_BuildWithDiscord(t *testing.T) {
	gpt := NewClient()
	prompt := "Build a program in Golang that gets the Price of Ethereum using the Coingecko API and post the response as a message to a Discord channel using github.com/bwmarrin/discordgo."
	resp, err := gpt.Codegen(context.Background(), prompt)
	if err != nil {
		t.Fatalf("Codegen() err: %v", err)
	}
	t.Log(resp)

	// Create a Discord session
	dg, err := discordgo.New("Bot " + os.Getenv("TOKEN"))
	if err != nil {
		t.Fatalf("Error creating Discord session: %v", err)
	}

	// Post a message to the Discord channel as a code snippet
	discord_msg := fmt.Sprintf("```go\n\n%s```", resp)
	_, err = dg.ChannelMessageSend(os.Getenv("CHANNEL_ID"), discord_msg)
	if err != nil {
		t.Fatalf("Error sending message to Discord channel: %v", err)
	}

}
