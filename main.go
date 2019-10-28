package main

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func RequestMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestID := r.Header.Get("X-Request-ID")
		if len(requestID) == 0 {
			requestID = uuid.New().String()
		}

		ctx := context.WithValue(r.Context(), "requestId", requestID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func HandleGetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Ok"))
}

func main() {
	r := mux.NewRouter()

	r.Use(RequestMiddleware)
	r.HandleFunc("/users/{id}", HandleGetUser)

	log.Info("Starting server on port 9001")
	log.Fatal(http.ListenAndServe("localhost:9001", r), "Listening on port 9001")
}
