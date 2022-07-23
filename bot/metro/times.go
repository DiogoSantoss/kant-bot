package metro

import (
	"strconv"

	"github.com/bwmarrin/discordgo"

	"github.com/DiogoSantoss/kant-bot/bot/discord"
)

type Time struct {
	Id       string    `json:"id"`
	Pier     string    `json:"pier"`
	Time     string    `json:"time"`
	Arrivals []Arrival `json:"arrivals"`
	Dest     string    `json:"dest"`
	Exit     string    `json:"exit"`
	UT       string    `json:"ut"`
}

type Arrival struct {
	Train     string `json:"train"`
	Time      string `json:"time"`
	Remaining string `json:"remaining"`
}

func SendMessageTimes(s *discordgo.Session, m *discordgo.MessageCreate, times []Time) {

	if len(times) == 0 {
		s.ChannelMessageSend(m.ChannelID, "Estação não encontrada")
		return
	}

	embedPages := []*discordgo.MessageEmbed{}

	// Array of times for each destination
	timesByDest := map[string][]Time{}

	for _, time := range times {
		timesByDest[time.Dest] = append(timesByDest[time.Dest], time)
	}

	// Create embed for each destination
	i := 1
	for dest, times := range timesByDest {
		embed := &discordgo.MessageEmbed{
			Title:  "Estação " + times[0].Id + " com destino " + dest,
			Fields: []*discordgo.MessageEmbedField{},
			Color:  discord.Orange,
			Footer: &discordgo.MessageEmbedFooter{
				Text: strconv.Itoa(i) + "/" + strconv.Itoa(len(timesByDest)),
			},
		}
		// Each destination tecnically has only one time
		for _, singleTime := range times {
			// Each time has multiple arrivals
			for idx, arrival := range singleTime.Arrivals {
				idx++
				idxStr := strconv.Itoa(idx)
				embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
					Name:  idxStr + " - " + "Hora de chegada",
					Value: arrival.Time + " (" + arrival.Remaining + ")",
				})
			}
		}
		embedPages = append(embedPages, embed)
		i++
	}

	message, _ := s.ChannelMessageSendEmbed(m.ChannelID, embedPages[0])
	discord.CreatePageEmbed(s, embedPages, message)
}
