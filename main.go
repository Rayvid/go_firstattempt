package main

import (
	"log"
	"net/http"

	"github.com/Rayvid/go_firstattempt/internal/session"
)

func handler(w http.ResponseWriter, r *http.Request) {
	session.HandleSession(w, r)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
