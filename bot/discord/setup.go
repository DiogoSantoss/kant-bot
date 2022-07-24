package discord

import (
	"fmt"

	"github.com/bwmarrin/discordgo"

	"github.com/DiogoSantoss/kant-bot/bot/config"
)

// Create discord session with intents
func Setup() (*discordgo.Session, error) {

	// Create a new Discord session using the provided bot token.
	discordSession, err := discordgo.New("Bot " + config.GetToken())
	if err != nil {
		fmt.Println("Error creating Discord session: ", err)
		return nil, err
	}

	// Only care about receiving message events.
	discordSession.Identify.Intents = discordgo.IntentsGuildMessages | discordgo.IntentGuildMessageReactions

	return discordSession, nil
}

// Start discord session
func Start(discordSession *discordgo.Session) error {

	err := discordSession.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return err
	}

	return nil
}

// Close discord session
func Stop(discordSession *discordgo.Session) {

	// Cleanly close down the Discord session.
	err := discordSession.Close()
	if err != nil {
		fmt.Println("error closing connection,", err)
	}
}
