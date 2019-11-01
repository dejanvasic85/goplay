package main

import (
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

var users = map[string]string {
	"foo": "Mister Fooooo",
	"bar": "Missus Barrrr",
}

func HandleGetUser(w http.ResponseWriter, r *http.Request) {
	userID := r.Form.Get("id")
	log.Infof("Fetching user %s", userID)

	if value, exists := users.Exists

	w.Write([]byte("Ok"))
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/users/{id}", HandleGetUser)

	log.Info("Starting server on port 9001")
	log.Fatal(http.ListenAndServe("localhost:9001", r), "Listening on port 9001")
}
