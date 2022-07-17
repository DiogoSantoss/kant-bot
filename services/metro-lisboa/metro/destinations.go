package metro

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
)

type ResponseDestinations struct {
	Response []Destination `json:"resposta"`
	Code     string        `json:"codigo"`
}

type Destination struct {
	Id   string `json:"id_destino"`
	Name string `json:"nome_destino"`
}

// Global variable to store destinations
var destinations map[string]string

func GetDestination(id string) string {
	return destinations[id]
}

func LoadDestinations(client *http.Client) error {

	// Initialize map
	destinations = make(map[string]string)

	endpoint := "https://api.metrolisboa.pt:8243/estadoServicoML/1.0.1/infoDestinos/todos"

	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return err
	}
	req.Header = http.Header{
		"Authorization": {"Bearer " + os.Getenv("METRO_TOKEN")},
	}
	res, err := client.Do(req)
	if err != nil {
		return err
	}

	// Close the connection to reuse it
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	var destinationsResponse ResponseDestinations
	err = json.Unmarshal(body, &destinationsResponse)
	if err != nil {
		return err
	}

	// Read body response to destinations map
	for _, destinationResponse := range destinationsResponse.Response {
		destinations[destinationResponse.Id] = destinationResponse.Name
	}

	return nil
}
