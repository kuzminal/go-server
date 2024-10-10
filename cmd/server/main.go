package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kuzminal/http-server-prod/internal/config"
)

func main() {
	c := config.Params
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!")
	})
	log.Printf("Starting server on %s port...", c.Port)
	if err := http.ListenAndServe(":"+c.Port, nil); err != nil {
		log.Fatal(err)
	}
}
