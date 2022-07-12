package handlers

import (
	"strings"

	"github.com/bwmarrin/discordgo"

	"github.com/DiogoSantoss/kant-bot/bot/config"
	"github.com/DiogoSantoss/kant-bot/bot/discord"
)

func Setup(discordSession *discordgo.Session) {

	// Add a handler for the messageCreate events.
	discordSession.AddHandler(commandHandler)
	discordSession.AddHandler(reactionHandler)
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

func reactionHandler(s *discordgo.Session, r *discordgo.MessageReactionAdd) {

	// Ignore all messages created by the bot itself
	if r.UserID == s.State.User.ID {
		return
	}

	// Fetch some extra information about the message associated to the reaction
	m, err := s.ChannelMessage(r.ChannelID, r.MessageID)
	// Ignore reactions on messages that have an error or that have not been sent by the bot
	if err != nil || m == nil || m.Author.ID != s.State.User.ID {
		return
	}

	// Ignore messages that are not embeds with a command in the footer
	if len(m.Embeds) != 1 {
		return
	}

	user, err := s.User(r.UserID)
	// Ignore when sender is invalid or is a bot
	if err != nil || user == nil || user.Bot {
		return
	}

	// handle reactions for paged embeds

	// Stop responding to reactions after 15 minutes
	discord.DeleteTimeout()

	// Used to move pages in embeds
	// TODO refactor
	// This uses a global hash of embeds which is not ideal
	pagedEmbed, found := discord.PagedEmbeds[r.MessageID]
	if !found {
		return
	}

	pagedEmbed.SwitchPage(r)
}
