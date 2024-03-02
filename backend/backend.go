package main

import (
	"encoding/json"
	"net/http"
	"time"
)

type Location struct {
	Latitude  float32   `json:"latitude"`
	Longitude float32   `json:"longitude"`
	Timestamp time.Time `json:"timestamp"`
}

var locations map[string]Location

func handleNearest(w http.ResponseWriter, r *http.Request) {

}

func handleLocation(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	var uid string
	if uid = q.Get("user_id"); uid == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	var l Location
	switch r.Method {
	case "GET":
		l, ok := locations[uid]
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(l)
	case "POST":
		if err := json.NewDecoder(r.Body).Decode(&l); err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		locations[uid] = l
	}
}

func main() {
	http.HandleFunc("/location/nearest", handleNearest)
	http.HandleFunc("/location", handleLocation)
}
