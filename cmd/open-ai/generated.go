package open_ai

import (
	"encoding/json"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"io/ioutil"
	"net/http"
)

// Structure for the Coingecko Ethereum Price API 
type EthereumPrice struct {
	MarketPriceUsd float64 `json:"market_data"`
}

func main() {
	// Fetch the Ethereum Price from Coingecko
	response, err := http.Get("https://api.coingecko.com/api/v3/simple/price?ids=ethereum&vs_currencies=usd")
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)

		// Unmarshal the JSON response into the EthereumPrice structure
		var ethereumPrice EthereumPrice
		json.Unmarshal(data, &ethereumPrice)

		// Print the Ethereum price
		fmt.Printf("The current Ethereum price is %f USD\n", ethereumPrice.MarketPriceUsd)

		// Create a new Discord Session
		discord, err := discordgo.New("Bot TOKEN")
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		// Post the Ethereum price to a Discord channel
		discord.ChannelMessageSend("CHANNEL_ID", fmt.Sprintf("The current Ethereum price is %f USD\n", ethereumPrice.MarketPriceUsd))
	}
}
