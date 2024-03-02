package main

import (
	"net/http"
)

type Location struct {
	latitude  float32
	longitude float32
}

func handleLocation(w http.ResponseWriter, r *http.Request) {

}

func main() {
	http.HandleFunc("/location", handleLocation)
}
