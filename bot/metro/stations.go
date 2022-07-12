package metro

import (
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

func CreateMessageStations(stations []Station) *discordgo.MessageEmbed {

	// Embed has a page for each line in metro
	embedPages := []*discordgo.MessageEmbed{}

	// Array of stations for each line
	stationsByLine := map[string][]Station{}

	for _, station := range stations {
		for _, line := range station.Lines {
			stationsByLine[line] = append(stationsByLine[line], station)
		}
	}

	for line, stations := range stationsByLine {
		embed := &discordgo.MessageEmbed{
			Title:  "Linha " + line,
			Fields: []*discordgo.MessageEmbedField{},
			Color:  colorForLine(line),
		}
		for _, station := range stations {

			embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
				Name:  station.Name,
				Value: station.Id + " @ " + "(" + station.Lat + ", " + station.Lon + ")",
			})
		}
		embedPages = append(embedPages, embed)
	}

	return embedPages[0]
}

func colorForLine(line string) int {
	switch line {
	case "Amarela":
		return discord.Yellow
	case "Azul":
		return discord.Blue
	case "Verde":
		return discord.Green
	case "Vermelha":
		return discord.Red
	default:
		return discord.White
	}
}
