package main

import (
	"context"
	"encoding/json"
	"log"
	"math"
	"net/http"
	"sort"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Location struct {
	Latitude  float64   `bson:"latitude" json:"latitude"`
	Longitude float64   `bson:"longitude" json:"longitude"`
	Timestamp time.Time `bson:"timestamp" json:"timestamp"`
	UserID    string    `bson:"user_id,omitempty" json:"user_id,omitempty"`
	distance  float64   `bson:"-" json:"-"`
}

func (l Location) Distance(o Location, maxRange float64) (float64, error) {
	dlat := math.Abs(l.Latitude - o.Latitude)
	dlon := math.Abs(l.Longitude - o.Longitude)
	if dlat > maxRange || dlon > maxRange {
		return 0, errRange
	}
	return math.Sqrt(math.Pow(dlat, 2) + math.Pow(dlon, 2)), nil
}

var locations *mongo.Collection

func handleNearest(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	var uid string
	if uid = q.Get("user_id"); uid == "" {
		http.Error(w, "empty user_id", http.StatusBadRequest)
		return
	}
	log.Println("Received request for /location/nearest")
	var l Location
	err := locations.FindOne(context.TODO(), bson.D{{"user_id", uid}}).Decode(&l)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
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

	w.Header().Set("Access-Control-Allow-Origin", "*")

	nearest := make([]Location, 0, n)
	var o Location
	cursor, err := locations.Find(context.TODO(), bson.D{{"user_id", bson.D{{"$ne", uid}}}})
	for cursor.Next(context.TODO()) {
		cursor.Decode(&o)
		if o.UserID == uid {
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
	log.Println("Found", len(nearest), "nearest locations")
	json.NewEncoder(w).Encode(nearest)
}

func handleLocation(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	var uid string
	if uid = q.Get("user_id"); uid == "" {
		http.Error(w, "empty user_id", http.StatusBadRequest)
		return
	}
	log.Println("Received /location request for", uid)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var l Location
	switch r.Method {
	case http.MethodGet:
		// l, ok := locations[uid]
		err := locations.FindOne(context.TODO(), bson.D{{"user_id", uid}}).Decode(&l)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
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
		if res, err := locations.ReplaceOne(context.TODO(), bson.D{{"user_id", uid}}, l); err != nil {
			log.Println("Error:", err.Error())
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		} else if res.MatchedCount == 0 {
			if _, err := locations.InsertOne(context.TODO(), l); err != nil {
				log.Println("Error:", err.Error())
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
		}
		log.Println("Post success")
		w.WriteHeader(http.StatusOK)
	}
}
