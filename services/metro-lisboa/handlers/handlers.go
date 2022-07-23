package handlers

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

// Dependency injection
func NewHandlers(client *http.Client, logger *log.Logger) *Handlers {
	return &Handlers{
		client: client,
		logger: logger,
	}
}

// Logger Middleware
func (h *Handlers) Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		defer h.logger.Printf("request processed in %s\n", time.Since(startTime))
		next(w, r)
	}
}

// Routes
func (h *Handlers) SetupRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/lines", h.Logger(h.GetLines))
	mux.HandleFunc("/stations", h.Logger(h.GetStations))
	mux.HandleFunc("/waitingtime", h.Logger(h.GetWaitingTime))
}

func (h *Handlers) GetStations(w http.ResponseWriter, r *http.Request) {
	h.logger.Println("request received to /stations endpoint")

	endpoint := "https://api.metrolisboa.pt:8243/estadoServicoML/1.0.1/infoEstacao/todos"

	res, err := h.sendRequest(endpoint)
	if err != nil {
		return 
	}

	// Close the connection to reuse it
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		h.logger.Printf("failed to read responde body")
	}

	// Correct some fields from the response
	body, err = ParseStations(body)
	if err != nil {
		h.logger.Printf("failed to parse response body")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(body)
		return
	}

	// Send response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func (h *Handlers) GetLines(w http.ResponseWriter, r *http.Request) {
	h.logger.Println("request received /lines endpoint")

	endpoint := "https://api.metrolisboa.pt:8243/estadoServicoML/1.0.1/estadoLinha/todos"

	res, err := h.sendRequest(endpoint)
	if err != nil {
		return 
	}

	// Close the connection to reuse it
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		h.logger.Printf("failed to read responde body")
	}

	// Correct some fields from the response
	body, err = ParseLines(body)
	if err != nil {
		h.logger.Printf("failed to parse response body")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(body)
		return
	}

	// Send response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

// NOTE: This endpoint accepts requests with the following query parameters:
// station: The station id
// Example: /waitingtime?station=CG
func (h *Handlers) GetWaitingTime(w http.ResponseWriter, r *http.Request) {
	h.logger.Println("request received /waitingtime endpoint")

	// URL Query parameters
	station := r.URL.Query().Get("station")

	endpoint := "https://api.metrolisboa.pt:8243/estadoServicoML/1.0.1/tempoEspera/Estacao/" + station

	res, err := h.sendRequest(endpoint)
	if err != nil {
		return 
	}

	// Close the connection to reuse it
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		h.logger.Printf("failed to read responde body")
	}

	// Correct some fields from the response
	body, err = ParseTimes(body)
	if err != nil {
		h.logger.Printf("failed to parse response body")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(body)
		return
	}

	// Send response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

// Generic request sender
func (h *Handlers) sendRequest(endpoint string) (*http.Response, error) {

	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		h.logger.Printf("failed to create request: %v", err)
		return nil, err
	}
	req.Header = http.Header{
		"Authorization": {"Bearer " + os.Getenv("METRO_TOKEN")},
	}
	res, err := h.client.Do(req)
	if err != nil {
		h.logger.Printf("failed request: %v", err)
		return nil, err
	}

	return res, nil
}
