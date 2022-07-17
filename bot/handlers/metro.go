package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/bwmarrin/discordgo"

	"github.com/DiogoSantoss/kant-bot/bot/config"
	"github.com/DiogoSantoss/kant-bot/bot/metro"
)

type ResponseStations struct {
	Stations []metro.Station `json:"stations"`
}

type ResponseLines struct {
	Lines []metro.Line `json:"lines"`
}

type ResponseWaitingTimes struct {
	Times []metro.Time `json:"times"`
}

func CommandStations(s *discordgo.Session, m *discordgo.MessageCreate) {

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

	metro.SendMessageStations(s, m, response.Stations)
}

func CommandLines(s *discordgo.Session, m *discordgo.MessageCreate) {

	endpoint := "http://localhost:8080/lines"

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
	var response ResponseLines
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		config.GetLogger().Println(err)
	}

	metro.SendMessageLines(s, m, response.Lines)
}

func CommandWaitingTimes(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Split the message into an array of words
	args := strings.Split(m.Content, " ")

	if len(args) < 3 {
		s.ChannelMessageSend(m.ChannelID, "Please provide a station name")
		return
	}

	// TODO: verify args
	endpoint := "http://localhost:8080/waitingtime?station=" + args[2]

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
	var response ResponseWaitingTimes
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		config.GetLogger().Println(err)
	}

	metro.SendMessageTimes(s, m, response.Times)
}
