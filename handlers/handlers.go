package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

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
	reply, err := redis.GetKeys("user:*")
	if err != nil {
		users = model.GetUsers()
		payload, err = json.Marshal(users)
	} else {
		users2, _ := reply.Array()
		for _, k := range users2 {
			key, _ := k.String()
			rep, _ := redis.GetHashAll(key)
			m, err := rep.Map()
			if err != nil {
				log.Fatal("something went wrong")
			}
			id, _ := strconv.Atoi(m["id"])
			user := model.Users{
				ID:        id,
				FirstName: m["firstname"],
				LastName:  m["lastname"],
				Email:     m["email"],
				Gender:    m["gender"],
			}
			users = append(users, user)
		}
		payload, err = json.Marshal(users)
	}

	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}

// AddUser add a new user to database
func AddUser(w http.ResponseWriter, r *http.Request) {
	payload, err := json.Marshal([]string{
		"hey", "whats", "up",
	})
	if err != nil {
		fmt.Println("something went wrong")
	}
	w.Write(payload)
}
