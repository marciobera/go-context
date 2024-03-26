package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	router := http.NewServeMux()

	router.HandleFunc("GET /", home)

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	log.Println("Starting server on port :8080")
	server.ListenAndServe()
}

func home(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println("request started")
	defer log.Println("request closed")

	select {
	case <-time.After(time.Second * 5):
		log.Println("request processed successfully")
		w.Write([]byte("request processed successfully"))
	case <-ctx.Done():
		http.Error(w, "request canceled", http.StatusRequestTimeout)
	}
}
