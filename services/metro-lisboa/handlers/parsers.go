package handlers

// TODO
// This is not very pretty, it should be refactored
// The ideia is to transform the response from the metro API
// to another (better) struct

import (
	"encoding/json"
	"strings"
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
	Urls    string `json:"stop_url"`
	Lines   string `json:"linha"`
	Zone_id string `json:"zone_id"`
}

type ParsedResponseStations struct {
	Stations []ParsedStation `json:"stations"`
}

type ParsedStation struct {
	Id      string   `json:"id"`
	Name    string   `json:"name"`
	Lat     string   `json:"lat"`
	Lon     string   `json:"lon"`
	Urls    []string `json:"urls"`
	Lines   []string `json:"lines"`
	Zone_id string   `json:"zone_id"`
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

type ParsedResponseLines struct {
	Lines []ParsedLine `json:"lines"`
}

type ParsedLine struct {
	Name   string `json:"name"`
	Status string `json:"status"`
	Msg    string `json:"msg"`
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

type ParsedResponseTimes struct {
	Times []ParsedTime `json:"times"`
}

type ParsedTime struct {
	Id       string    `json:"id"`
	Pier     string    `json:"pier"`
	Hour     string    `json:"hour"`
	Arrivals []Arrival `json:"arrivals"`
	Dest     string    `json:"dest"`
	Exit     string    `json:"exit"`
	UT       string    `json:"ut"`
}

type Arrival struct {
	Train string `json:"train"`
	Time  string `json:"time"`
}

func ParseStations(body []byte) ([]byte, error) {

	var response ResponseStations
	json.Unmarshal(body, &response)

	var parsedResponse ParsedResponseStations

	for _, station := range response.Response {
		station.Urls = strings.Replace(station.Urls, "[", "", -1)
		station.Urls = strings.Replace(station.Urls, "]", "", -1)
		station.Urls = strings.Replace(station.Urls, " ", "", -1)
		urls := strings.Split(station.Urls, ",")

		station.Lines = strings.Replace(station.Lines, "[", "", -1)
		station.Lines = strings.Replace(station.Lines, "]", "", -1)
		station.Lines = strings.Replace(station.Lines, " ", "", -1)
		lines := strings.Split(station.Lines, ",")

		parsedStation := &ParsedStation{
			Id:      station.Id,
			Name:    station.Name,
			Lat:     station.Lat,
			Lon:     station.Lon,
			Urls:    urls,
			Lines:   lines,
			Zone_id: station.Zone_id,
		}

		parsedResponse.Stations = append(parsedResponse.Stations, *parsedStation)
	}

	toSend, err := json.Marshal(parsedResponse)
	if err != nil {
		return toSend, err
	}

	return toSend, nil
}

func ParseLines(body []byte) ([]byte, error) {

	var response ResponseLines
	json.Unmarshal(body, &response)

	var parsedResponse ParsedResponseLines

	yellow := &ParsedLine{
		Name:   "Yellow",
		Status: strings.Replace(response.Response.YellowStatus, " ", "", -1),
		Msg:    response.Response.YellowMsg,
	}

	red := &ParsedLine{
		Name:   "Red",
		Status: strings.Replace(response.Response.RedStatus, " ", "", -1),
		Msg:    response.Response.RedMsg,
	}

	blue := &ParsedLine{
		Name:   "Blue",
		Status: strings.Replace(response.Response.BlueStatus, " ", "", -1),
		Msg:    response.Response.BlueMsg,
	}

	green := &ParsedLine{
		Name:   "Green",
		Status: strings.Replace(response.Response.GreenStatus, " ", "", -1),
		Msg:    response.Response.GreenMsg,
	}

	parsedResponse.Lines = append(parsedResponse.Lines, *yellow)
	parsedResponse.Lines = append(parsedResponse.Lines, *red)
	parsedResponse.Lines = append(parsedResponse.Lines, *blue)
	parsedResponse.Lines = append(parsedResponse.Lines, *green)

	toSend, err := json.Marshal(parsedResponse)
	if err != nil {
		return toSend, err
	}

	return toSend, nil
}

func ParseTimes(body []byte) ([]byte, error) {

	var response ResponseWaitingTimes
	json.Unmarshal(body, &response)

	var parsedResponse ParsedResponseTimes

	for _, waitingTime := range response.Response {

		arrivals := []Arrival{
			{
				Train: waitingTime.Train1,
				Time:  waitingTime.Time1,
			},
			{
				Train: waitingTime.Train2,
				Time:  waitingTime.Time2,
			},
			{
				Train: waitingTime.Train3,
				Time:  waitingTime.Time3,
			},
		}

		parsedTime := &ParsedTime{
			Id:       waitingTime.Id,
			Pier:     waitingTime.Pier,
			Hour:     waitingTime.Hour,
			Arrivals: arrivals,
			Dest:     waitingTime.Dest,
			Exit:     waitingTime.Exit,
			UT:       waitingTime.UT,
		}

		parsedResponse.Times = append(parsedResponse.Times, *parsedTime)
	}

	toSend, err := json.Marshal(parsedResponse)
	if err != nil {
		return toSend, err
	}

	return toSend, nil
}
