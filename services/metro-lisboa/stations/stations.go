package stations

import (
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type Handlers struct {
	client *http.Client
	logger *log.Logger
}

type Station struct {
	id string		`json:"stop_id"`
	name string		`json:"stop_name"`
	lat float32		`json:"stop_lat"`
	long float32	`json:"stop_long"`
	url []string	`json:"stop_url"`
	linha []string	`json:"linha"`
	zone_id string	`json:"zone_id"`
}

func (h *Handlers) GetStations(w http.ResponseWriter, r *http.Request) {
	h.logger.Println("request received")

	endpoint := "https://api.metrolisboa.pt:8243/estadoServicoML/1.0.1/infoEstacao/todos"

	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		h.logger.Printf("failed to create request: %v", err)
	}
	req.Header = http.Header {
		"Authorization": {"Bearer " + os.Getenv("METRO_TOKEN")},
	}
	res, err := h.client.Do(req)
	if err != nil {
		h.logger.Printf("failed request: %v", err)
	}

	// Close the connection to reuse it
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		h.logger.Printf("failed to read responde body")
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(body))
}

// Logger Middleware
func (h *Handlers) Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		defer h.logger.Printf("request processed in %s\n", time.Since(startTime))
		next(w, r)
	}
}

func (h *Handlers) SetupRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/stations", h.Logger(h.GetStations))
}

// Dependency injection
func NewHandlers(client *http.Client, logger *log.Logger) *Handlers {
	return &Handlers{
		client: client,
		logger: logger,
	}
}