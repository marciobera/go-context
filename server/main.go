package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", home)
	http.ListenAndServe(":8080", nil)
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
