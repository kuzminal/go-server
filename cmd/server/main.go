package main

import (
	"flag"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	_ "net/http/pprof"

	"github.com/kuzminal/http-server-prod/internal/config"
	"github.com/kuzminal/http-server-prod/internal/server"
	"github.com/kuzminal/http-server-prod/pkg/api"
)

func main() {
	var confPath string
	flag.StringVar(&confPath, "conf", "", "Path to config file")
	flag.Parse()
	conf := config.LoadConfig(confPath)

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.Level(conf.LogLevel)}))
	srv := server.NewServer(logger)

	r := chi.NewRouter()

	handler := api.HandlerFromMux(srv, r)

	s := &http.Server{
		Handler: handler,
		Addr:    "0.0.0.0:" + conf.Port,
	}

	slog.Info("Starting server...", "port", conf.Port)
	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err.Error())
	}
}
