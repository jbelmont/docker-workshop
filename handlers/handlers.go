package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/jbelmont/docker-workshop/model"
	"github.com/jbelmont/docker-workshop/redis"
)

// UsersIndex returns index page with users
func UsersIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "")
}

// User Struct is the model for the app
type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Gender    string `json:"gender"`
}

// GetUsers returns json payload with users
func GetUsers(w http.ResponseWriter, r *http.Request) {
	var users []model.Users
	var payload []byte
	var err error
	reply, err := redis.GetKey("users")
	if err != nil {
		users = model.GetUsers()
		payload, err = json.Marshal(users)
	} else {
		value, _ := reply.String()
		json.Unmarshal([]byte(value), &users)
		payload, err = json.Marshal(users)
	}

	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}
