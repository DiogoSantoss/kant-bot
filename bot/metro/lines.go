package metro

import "github.com/bwmarrin/discordgo"

type Line struct {
	Name   string `json:"name"`
	Status string `json:"status"`
	Msg    string `json:"msg"`
}

func SendMessageLines(s *discordgo.Session, m *discordgo.MessageCreate, lines []Line) {

	// Embed for each line
	embeds := []*discordgo.MessageEmbed{}
	for _, line := range lines {

		embed := &discordgo.MessageEmbed{
			Title:  "Linha " + line.Name,
			Fields: []*discordgo.MessageEmbedField{},
			Color:  colorForLine(line.Name),
		}

		embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
			Name:  "Status",
			Value: line.Status,
		})

		embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
			Name:  "Status message",
			Value: line.Msg,
		})

		embeds = append(embeds, embed)
	}

	s.ChannelMessageSendEmbeds(m.ChannelID, embeds)
}
