package main

import (
	"log"
	"net/http"

	"github.com/Rayvid/go_firstattempt/internal/session"
	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.Handle("GET", "/infocenter/:topic", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		session.HandleSession(w, r, p) /* TODO smth with error */
	})
	router.Handle("POST", "/infocenter/:topic", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		session.HandleSession(w, r, p) /* TODO smth with error */
	})
	log.Fatal(http.ListenAndServe(":8080", router))
}
