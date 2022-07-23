package handlers

import (
	"github.com/bwmarrin/discordgo"
	
	"github.com/DiogoSantoss/kant-bot/bot/discord"
)

func CommandHelp(s *discordgo.Session, m *discordgo.MessageCreate) {

	title := "Kant Bot - List of commands"
	description := "List of commands available in the bot"
	fields := []*discordgo.MessageEmbedField{
		{
			Name:  "kant help",
			Value: "Shows this message",
		},
		// Metro Service
		{
			Name:  "kant stations",
			Value: "Replies with all stations from Lisbon's Metro organized by station color",
		},
		{
			Name:  "kant lines",
			Value: "Replies with all lines from Lisbon's Metro and their current status",
		},
		{
			Name:  "kant time <line_id>",
			Value: "Replies with the current waiting time for the given line",
		},
	}

	s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Title:       title,
		Description: description,
		Fields:      fields,
		Color:       discord.Blue,
	})
}
