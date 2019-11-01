package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

var users = map[string]string{
	"foo": "Mister Fooooo",
	"bar": "Missus Barrrr",
}

func CorrelationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := r.Header.Get("Correlation-ID")
		ctx := context.WithValue(r.Context(), "CorrelationID", id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func HandleGetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID := params["id"]

	log.Infof("Fetching user %s", userID)

	if value, exists := users[userID]; exists {
		log.Infof("Found user %s", value)
		data := struct {
			Hello string
		}{
			Hello: value,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(data)
		return
	}

	w.WriteHeader(http.StatusNotFound)
}

func main() {
	r := mux.NewRouter()

	r.Use(CorrelationMiddleware)
	r.HandleFunc("/users/{id}", HandleGetUser)

	log.Info("Starting server on port 9001")
	log.Fatal(http.ListenAndServe("localhost:9001", r), "Listening on port 9001")
}
