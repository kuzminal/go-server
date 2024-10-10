package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/kuzminal/http-server-prod/internal/config"
	"github.com/kuzminal/http-server-prod/internal/server"
	"github.com/kuzminal/http-server-prod/pkg/api"
)

var log = config.Logger

func main() {
	conf := config.Params
	server := server.NewServer()

	r := chi.NewRouter()

	handler := api.HandlerFromMux(server, r)

	s := &http.Server{
		Handler: handler,
		Addr:    "0.0.0.0:" + conf.Port,
	}

	log.Info().Msgf("Starting server on %s port...", conf.Port)
	if err := s.ListenAndServe(); err != nil {
		log.Fatal().Err(err)
	}
}
