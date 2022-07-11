package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/bwmarrin/discordgo"

	"github.com/DiogoSantoss/kant-bot/bot/config"
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

type ResponseLines struct {
	Response Lines  `json:"resposta"`
	Code     string `json:"codigo"`
}

type Lines struct {
	YellowStatus string `json:"amarela"`
	YellowMsg    string `json:"tipo_msg_am"`
	RedStatus    string `json:"vermelha"`
	RedMsg       string `json:"tipo_msg_vm"`
	BlueStatus   string `json:"azul"`
	BlueMsg      string `json:"tipo_msg_az"`
	GreenStatus  string `json:"verde"`
	GreenMsg     string `json:"tipo_msg_vd"`
}

type ResponseWaitingTimes struct {
	Response []WaitingTime `json:"resposta"`
	Code     string        `json:"codigo"`
}

type WaitingTime struct {
	Id     string `json:"stop_id"`
	Pier   string `json:"cais"`
	Hour   string `json:"hora"`
	Train1 string `json:"comboio"`
	Time1  string `json:"tempoChegada1"`
	Train2 string `json:"comboio2"`
	Time2  string `json:"tempoChegada2"`
	Train3 string `json:"comboio3"`
	Time3  string `json:"tempoChegada3"`
	Dest   string `json:"destino"`
	Exit   string `json:"sairServico"`
	UT     string `json:"UT"`
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

	message := ""
	for _, station := range response.Response {
		message = message + station.Name + "\n"
	}

	s.ChannelMessageSend(m.ChannelID, message)
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

	message := ""
	message = message + "Amarela: " + response.Response.YellowStatus + "\n"
	message = message + "Vermelha: " + response.Response.RedStatus + "\n"
	message = message + "Azul: " + response.Response.BlueStatus + "\n"
	message = message + "Verde: " + response.Response.GreenStatus + "\n"

	s.ChannelMessageSend(m.ChannelID, message)
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

	message := ""
	for _, waitingTime := range response.Response {
		message = message + waitingTime.Id + "\n"
		message = message + waitingTime.Pier + "\n"
		message = message + waitingTime.Hour + "\n"
		message = message + waitingTime.Train1 + "\n"
		message = message + waitingTime.Time1 + "\n"
		message = message + waitingTime.Train2 + "\n"
		message = message + waitingTime.Time2 + "\n"
		message = message + waitingTime.Train3 + "\n"
		message = message + waitingTime.Time3 + "\n"
		message = message + waitingTime.Dest + "\n"
		message = message + waitingTime.Exit + "\n"
		message = message + waitingTime.UT + "\n"
	}

	s.ChannelMessageSend(m.ChannelID, message)
}
