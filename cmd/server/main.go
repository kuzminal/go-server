package main

import (
	"flag"
	"log"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/kuzminal/http-server-prod/internal/config"
	"github.com/kuzminal/http-server-prod/internal/server"
	"github.com/kuzminal/http-server-prod/pkg/api"
)

func main() {
	var confPath string
	flag.StringVar(&confPath, "conf", "", "Path to config file")
	flag.Parse()

	conf := config.LoadConfig(confPath)

	server := server.NewServer()

	r := chi.NewRouter()

	handler := api.HandlerFromMux(server, r)

	s := &http.Server{
		Handler: handler,
		Addr:    "0.0.0.0:" + conf.Port,
	}

	slog.Info("Starting server...", "port", conf.Port)
	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err.Error())
	}
}
