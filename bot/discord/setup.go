package discord

import (
	"fmt"

	"github.com/bwmarrin/discordgo"

	"github.com/DiogoSantoss/kant-bot/bot/config"
)

// Create discord session with intents
func Setup() *discordgo.Session {

	// Create a new Discord session using the provided bot token.
	discordSession, err := discordgo.New("Bot " + config.GetToken())
	if err != nil {
		fmt.Println("Error creating Discord session: ", err)
	}

	// Only care about receiving message events.
	discordSession.Identify.Intents = discordgo.IntentsGuildMessages | discordgo.IntentGuildMessageReactions

	return discordSession
}

// Start discord session
func Start(discordSession *discordgo.Session) {

	err := discordSession.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}
}

// Close discord session
func Stop(discordSession *discordgo.Session) {

	// Cleanly close down the Discord session.
	err := discordSession.Close()
	if err != nil {
		fmt.Println("error closing connection,", err)
	}
}
