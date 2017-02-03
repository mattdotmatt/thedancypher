package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/mattdotmatt/thedancypher/engine"
)

func cypherHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	text := vars["text"]

	cyphered := engine.Encrypt(text)

	w.Write([]byte(cyphered))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/cypher/{text}", cypherHandler)
	http.Handle("/", r)

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
