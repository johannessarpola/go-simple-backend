package server

import (
	"fmt"
	"log/slog"
	"net/http"
)

type Server struct {
	Host       string
	Port       int
	httpServer *http.ServeMux
}

func NewServer(host string, port int) *Server {
	s := &Server{
		Host: host,
		Port: port,
	}
	return s
}

func (s *Server) Run() {
	s.httpServer = http.NewServeMux()

	port := s.Port
	host := s.Host

	// Define handlers for /health and / endpoints
	s.httpServer.HandleFunc("/health", s.healthHandler)
	s.httpServer.HandleFunc("/", s.rootHandler)

	slog.Info(
		"server is running on",
		"port", port,
		"host", host,
	)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), s.httpServer)
	if err != nil {
		slog.Error("Error starting the server:", err)
	}
}

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	slog.Debug(
		"healthcheck called",
		"from", r.Host,
		"headers", r.Header,
	)
	// write ok to response
	fmt.Fprint(w, "OK")
}

func (s *Server) rootHandler(w http.ResponseWriter, r *http.Request) {
	slog.Debug(
		"root called",
		"from", r.Host,
		"headers", r.Header,
	)

	response := fmt.Sprintf("Hello, this is the root endpoint on %s!", s.url())
	// write simple message to response
	fmt.Fprint(w, response)
}

func (s *Server) url() string {
	return fmt.Sprintf("%s:%d", s.Host, s.Port)
}
