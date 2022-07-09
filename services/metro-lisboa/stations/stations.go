package stations

import (
	"log"
	"net/http"
	"time"
)

type Handlers struct {
	logger *log.Logger
}

func (h *Handlers) GetStations(w http.ResponseWriter, r *http.Request) {
	h.logger.Println("request received")
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello World"))
}

// Logger Middleware
func (h *Handlers) Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		defer h.logger.Printf("request processed in %s\n", time.Now().Sub(startTime))
		next(w, r)
	}
}

func (h *Handlers) SetupRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/stations", h.Logger(h.GetStations))
}

// Logger dependency injection
func NewHandlers(logger *log.Logger) *Handlers {
	return &Handlers{
		logger: logger,
	}
}