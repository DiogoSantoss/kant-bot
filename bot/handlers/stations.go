package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/DiogoSantoss/kant-bot/bot/config"
	"github.com/bwmarrin/discordgo"
)

type ResponseStations struct {
	Response []Station `json:"resposta"`
	Code     string    `json:"codigo"`
}

type Station struct {
	Id      string `json:"stop_id"`
	Name    string `json:"stop_name"`
	Lat     string `json:"stop_lat"`
	Lon     string `json:"stop_lon"`
	Url     string `json:"stop_url"`
	Linha   string `json:"linha"`
	Zone_id string `json:"zone_id"`
}

func StationsCommand(s *discordgo.Session, m *discordgo.MessageCreate) {

	endpoint := "http://localhost:8080/stations"

	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		config.GetLogger().Println(err)
	}
	res, err := config.GetClient().Do(req)
	if err != nil {
		config.GetLogger().Println(err)
	}

	// Close the connection to reuse it
	defer res.Body.Close()

	// Parse response to struct
	var response ResponseStations
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		config.GetLogger().Println(err)
	}

	fmt.Println(response)

	message := ""
	for _, station := range response.Response {
		message = message + station.Name + "\n"
	}

	s.ChannelMessageSend(m.ChannelID, message)
}
