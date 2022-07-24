package handlers

import (
	"strings"
	"net/http"
	"encoding/json"

	"github.com/bwmarrin/discordgo"

	"github.com/DiogoSantoss/kant-bot/bot/metro"
	"github.com/DiogoSantoss/kant-bot/bot/config"
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

// Address of the metro service
const address = "http://metro:8080"

// Handler for the stations command
func CommandStations(s *discordgo.Session, m *discordgo.MessageCreate) {

	endpoint := address + "/stations"

	res, err := sendRequest(endpoint)
	if err != nil {
		config.GetLogger().Println(err)
		return
	}

	// Close the connection to reuse it
	defer res.Body.Close()

	// Parse response to struct
	var response ResponseStations
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		config.GetLogger().Println(err)
		return
	}

	metro.SendMessageStations(s, m, response.Stations)
}

// Handler for the lines command
func CommandLines(s *discordgo.Session, m *discordgo.MessageCreate) {

	endpoint := address + "/lines"

	res, err := sendRequest(endpoint)
	if err != nil {
		config.GetLogger().Println(err)
		return
	}
	// Close the connection to reuse it
	defer res.Body.Close()

	// Parse response to struct
	var response ResponseLines
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		config.GetLogger().Println(err)
		return
	}

	metro.SendMessageLines(s, m, response.Lines)
}

// Handler for the time command
func CommandWaitingTimes(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Split the message into an array of words
	args := strings.Split(m.Content, " ")

	if len(args) < 3 {
		s.ChannelMessageSend(m.ChannelID, "Please provide a station name")
		return
	}

	// TODO: verify args
	endpoint := address + "/waitingtime?station=" + args[2]

	res, err := sendRequest(endpoint)
	if err != nil {
		config.GetLogger().Println(err)
		return
	}

	// Close the connection to reuse it
	defer res.Body.Close()

	// Parse response to struct
	var response ResponseWaitingTimes
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		config.GetLogger().Println(err)
		return
	}

	metro.SendMessageTimes(s, m, response.Times)
}

// Generic function to send a request
func sendRequest(endpoint string) (*http.Response, error) {

	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}
	res, err := config.GetClient().Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
