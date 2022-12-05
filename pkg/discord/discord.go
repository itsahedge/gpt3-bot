package discord

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

type Bot struct {
	ds            *discordgo.Session
	channelId     string
	alertEveryone bool
}

func InitializedDiscord() (*Bot, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dg, err := discordgo.New("Bot " + os.Getenv("TOKEN"))
	if err != nil {
		fmt.Println("error creating Discord session", err)
		return nil, err
	}
	fmt.Println("Connected to Discord server")

	bot := &Bot{ds: dg, channelId: os.Getenv("CHANNEL_ID"), alertEveryone: true}
	dg.AddHandler(messageCreate)
	// dg.AddHandler(cb.OnNewMessage)
	// dg.AddHandler(cb.OnDisconnect)
	// cb.alertEveryone = conf.Discord.AlertEveryone
	return bot, err
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "!generate" {
		// TODO: replace w/ clients functions to get some data..
		resp, err := http.Get("https://example.com")
		if err != nil {
			fmt.Println(err)
		}
		defer resp.Body.Close()
		if resp.StatusCode == 200 {
			_, err = s.ChannelMessageSend(os.Getenv("CHANNEL_ID"), "hello\nnext line\n3rd line")
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Printf("status code = %v", resp.Status)
		}
	}

}
