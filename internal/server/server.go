package server

import (
	"encoding/json"
	"net/http"

	"github.com/kuzminal/http-server-prod/pkg/api"
)

type Server struct{}

func NewServer() Server {
	return Server{}
}

func (Server) Get(w http.ResponseWriter, r *http.Request, params api.GetParams) {
	name := ""
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
