package handlers

import "github.com/bwmarrin/discordgo"


func HelpCommand(s *discordgo.Session, m *discordgo.MessageCreate) {

	title := "Kant Bot - List of commands"
	description := "Testing embeds"
	fields := []*discordgo.MessageEmbedField{
		{
			Name: "kant help",
			Value: "Shows this message",
		},
		{
			Name: "kant ping",
			Value: "Replies with pong",
		},
	}


	s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Title:       title,
		Description: description,
		Fields:      fields,
		Color:       3447003,
	})
}