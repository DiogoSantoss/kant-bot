package handlers

import (
	"strings"

	"github.com/bwmarrin/discordgo"

	"github.com/DiogoSantoss/kant-bot/bot/config"
)

func Setup(discordSession *discordgo.Session) {

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
		CommandHelp(s, m)
	case "stations":
		CommandStations(s, m)
	case "lines":
		CommandLines(s, m)
	case "time":
		CommandWaitingTimes(s, m)
	default:
		s.ChannelMessageSend(m.ChannelID, "This command does not exist, use help to list all commands.")
	}
}
