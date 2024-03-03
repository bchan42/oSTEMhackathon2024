package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var errRange = errors.New("Location out of range")

const mongoURI = "mongodb://user:pass@mongo.lone-faerie.xyz:27017"

func main() {
	// locations = make(map[string]Location)

	log.SetOutput(os.Stdout)

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoURI))
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	db := client.Database("db")
	locations = db.Collection("locations")
	users = db.Collection("users")

	http.HandleFunc("/location/nearest", handleNearest)
	http.HandleFunc("/location", handleLocation)

	http.HandleFunc("/user/add", handleAddUser)
	http.HandleFunc("/user/remove", handleRemoveUser)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
