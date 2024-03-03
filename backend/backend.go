package main

import (
	"encoding/json"
	"errors"
	"log"
	"math"
	"net/http"
	"sort"
	"strconv"
	"time"
)

var errRange = errors.New("Location out of range")

// const mongoURI = "mongodb://user:pass@mongo.lone-faerie.xyz:27107"

type Location struct {
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
	Timestamp time.Time `json:"timestamp"`
	UserID    string    `json:"user_id,omitempty"`
	distance  float64   `json:"-"`
}

func (l Location) Distance(o Location, maxRange float64) (float64, error) {
	dlat := math.Abs(l.Latitude - o.Latitude)
	dlon := math.Abs(l.Longitude - o.Longitude)
	if dlat > maxRange || dlon > maxRange {
		return 0, errRange
	}
	return math.Sqrt(math.Pow(dlat, 2) + math.Pow(dlon, 2)), nil
}

var locations map[string]Location

// var locationsColl *mongo.Collection

func handleNearest(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	var uid string
	if uid = q.Get("user_id"); uid == "" {
		http.Error(w, "empty user_id", http.StatusBadRequest)
		return
	}
	l, ok := locations[uid]
	if !ok {
		http.Error(w, "location not found", http.StatusNotFound)
		return
	}
	var (
		rn float64
		n  int
	)
	if x, err := strconv.Atoi(q.Get("range")); err == nil {
		rn = float64(x)
	}
	if x, err := strconv.Atoi(q.Get("count")); err == nil {
		n = x
	}
	if n == 0 {
		n = 1
	}

	nearest := make([]Location, 0, n)
	for k, o := range locations {
		if k == uid {
			continue
		}
		if d, err := l.Distance(o, rn); err == nil {
			o.distance = d
			i := sort.Search(len(nearest), func(i int) bool {
				return nearest[i].distance >= o.distance
			})
			if i < len(nearest) {
				copy(nearest[i+1:cap(nearest)], nearest[i:])
				nearest[i] = o
			} else if len(nearest) < cap(nearest) {
				nearest = append(nearest, o)
			}
		}
	}
	json.NewEncoder(w).Encode(nearest)
}

func handleLocation(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	var uid string
	if uid = q.Get("user_id"); uid == "" {
		http.Error(w, "empty user_id", http.StatusBadRequest)
		return
	}
	var l Location
	switch r.Method {
	case http.MethodGet:
		l, ok := locations[uid]
		if !ok {
			http.Error(w, "location not found", http.StatusNotFound)
			return
		}
		l.UserID = ""
		json.NewEncoder(w).Encode(l)
	case http.MethodPost:
		if err := json.NewDecoder(r.Body).Decode(&l); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		l.UserID = uid
		locations[uid] = l
	}
}

func main() {
	locations = make(map[string]Location)

	// client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoURI))
	// if err != nil {
	// 	panic(err)
	// }

	// defer func() {
	// 	if err := client.Disconnect(context.TODO()); err != nil {
	// 		panic(err)
	// 	}
	// }()

	// locationsColl = client.Database("db").Collection("locations")

	http.HandleFunc("/location/nearest", handleNearest)
	http.HandleFunc("/location", handleLocation)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
