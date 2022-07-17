package handlers

import (
	"github.com/bwmarrin/discordgo"
	
	"github.com/DiogoSantoss/kant-bot/bot/discord"
)

func CommandHelp(s *discordgo.Session, m *discordgo.MessageCreate) {

	title := "Kant Bot - List of commands"
	description := "Testing embeds"
	fields := []*discordgo.MessageEmbedField{
		{
			Name:  "kant help",
			Value: "Shows this message",
		},
		{
			Name:  "kant stations",
			Value: "Replies with all stations from Lisbon's Metro",
		},
		{
			Name:  "kant lines",
			Value: "Replies with all lines from Lisbon's Metro",
		},
	}

	s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Title:       title,
		Description: description,
		Fields:      fields,
		Color:       discord.Blue,
	})
}
