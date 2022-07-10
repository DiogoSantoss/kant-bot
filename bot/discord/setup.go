package discord

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"

	"github.com/DiogoSantoss/kant-bot/bot/config"
	"github.com/DiogoSantoss/kant-bot/bot/handlers"
)

func Setup() *discordgo.Session {
	
	dS := start(config.GetToken())
	setIntents(dS)
	addHandlers(dS)
	openConnection(dS)

	return dS
}

func Stop(discordSession *discordgo.Session) {

	// Cleanly close down the Discord session.
	discordSession.Close()
}

func start(token string) *discordgo.Session {

	// Create a new Discord session using the provided bot token.
	discordSession, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("Error creating Discord session: ", err)
	}

	return discordSession
}

func openConnection(discordSession *discordgo.Session) {
	err := discordSession.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}
}

func setIntents(discordSession *discordgo.Session) {

	// Only care about receiving message events.
	discordSession.Identify.Intents = discordgo.IntentsGuildMessages
}

func addHandlers(discordSession *discordgo.Session) {

	// Add a handler for the messageCreate events.
	discordSession.AddHandler(commandHandler)
}

func commandHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	// Ignore messages that don't start with the prefix.
	if !strings.HasPrefix(m.Content, config.GetPrefix()) {
		return
	}

	// Get the arguments
	args := strings.Split(m.Content, " ")[1:]
	// Ensure valid command
	if len(args) == 0 {
		s.ChannelMessageSend(m.ChannelID, "I'm sorry, I don't know what you want me to do.")
		return
	}

	// Call the corresponding handler
	switch args[0] {
	case "help":
		handlers.HelpCommand(s, m)
	case "stations":
		handlers.StationsCommand(s, m)
	default:
		s.ChannelMessageSend(m.ChannelID, "This command does not exist, use help to list all commands.")
	}
}

