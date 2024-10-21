package server

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/kuzminal/http-server-prod/pkg/api"
)

type Server struct {
	Logger *slog.Logger
}

func NewServer(logger *slog.Logger) Server {
	return Server{Logger: logger}
}

func (s Server) Get(w http.ResponseWriter, r *http.Request, params api.GetParams) {
	name := ""
	s.Logger.Info("Request host was..", "headers", r.Host)
	if params.Name == nil {
		name = "World"
	} else {
		name = *params.Name
	}
	resp := api.Hello{
		Name: name,
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(resp)
}
