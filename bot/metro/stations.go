package metro

import (
	"strconv"

	"github.com/bwmarrin/discordgo"

	"github.com/DiogoSantoss/kant-bot/bot/discord"
)

type Station struct {
	Id      string   `json:"id"`
	Name    string   `json:"name"`
	Lat     string   `json:"lat"`
	Lon     string   `json:"lon"`
	Urls    []string `json:"urls"`
	Lines   []string `json:"lines"`
	Zone_id string   `json:"zone_id"`
}

func SendMessageStations(s *discordgo.Session, m *discordgo.MessageCreate, stations []Station) {

	// Embed has a page for each line in metro
	embedPages := []*discordgo.MessageEmbed{}

	// Array of stations for each line
	stationsByLine := map[string][]Station{}

	for _, station := range stations {
		for _, line := range station.Lines {
			stationsByLine[line] = append(stationsByLine[line], station)
		}
	}

	// Create embed for each line
	i := 1
	for line, stations := range stationsByLine {
		embed := &discordgo.MessageEmbed{
			Title:  "Linha " + line,
			Fields: []*discordgo.MessageEmbedField{},
			Color:  colorForLine(line),
			Footer: &discordgo.MessageEmbedFooter{
				Text: strconv.Itoa(i) + "/" + strconv.Itoa(len(stationsByLine)),
			},
		}
		for _, station := range stations {

			embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
				Name:  station.Name,
				Value: station.Id + " @ " + "(" + station.Lat + ", " + station.Lon + ")",
			})
		}
		embedPages = append(embedPages, embed)
		i++
	}

	message, _ := s.ChannelMessageSendEmbed(m.ChannelID, embedPages[0])
	discord.CreatePageEmbed(s, embedPages, message)
}
