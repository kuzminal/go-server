package main

import (
	"fmt"
	"net/http"

	"github.com/kuzminal/http-server-prod/internal/config"
)

var log = config.Logger

func main() {
	conf := config.Params
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!")
	})
	log.Info().Msgf("Starting server on %s port...", conf.Port)
	if err := http.ListenAndServe(":"+conf.Port, nil); err != nil {
		log.Fatal().Err(err)
	}
}
