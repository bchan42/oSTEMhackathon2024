package main

import (
	"context"
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	Name string `bson:"name" json:"name"`
}

var users *mongo.Collection

func handleAddUser(w http.ResponseWriter, r *http.Request) {
	var u User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	res := users.FindOne(context.TODO(), bson.D{{"name", u.Name}})
	if res.Err() != mongo.ErrNoDocuments {
		http.Error(w, "user already exists", http.StatusBadRequest)
		return
	}
	if _, err := users.InsertOne(context.TODO(), u); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusOK)
}

func handleRemoveUser(w http.ResponseWriter, r *http.Request) {
	var u User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	res, err := users.DeleteOne(context.TODO(), bson.D{{"name", u.Name}})
	if err != nil || res.DeletedCount == 0 {
		http.Error(w, "user does not exist", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}
