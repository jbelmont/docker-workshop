package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/jbelmont/docker-workshop/model"
)

// UsersIndex returns index page with users
func UsersIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "")
}

// GetUsers returns json payload with users
func GetUsers(w http.ResponseWriter, r *http.Request) {
	users := model.GetUsers()
	payload, err := json.Marshal(users)
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}
