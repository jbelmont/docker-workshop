package redis

import (
	"encoding/json"
	"log"
	"os"

	"github.com/jbelmont/docker-workshop/model"
	"github.com/keimoon/gore"
)

// User Struct is the model for the app
type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Gender    string `json:"gender"`
}

var connect *gore.Conn

// ConnectRedis connects to redis and sets users keys
func ConnectRedis() {
	var err error
	connect, err = gore.Dial(os.Getenv("REDIS_URL"))
	if err != nil {
		return
	}
	setUsers()
}

func setUsers() {
	data, err := json.Marshal(model.GetUsers())
	if err != nil {
		panic(err)
	}
	rep, err := gore.NewCommand("SET", "users", data).Run(connect)
	if err != nil || !rep.IsOk() {
		log.Fatal(err, "not ok")
	}
	defer connect.Close()
}

// GetConnection returns redis url
func GetConnection() *gore.Conn {
	var err error
	connect, err = gore.Dial("localhost:6379")
	if err != nil {
		return nil
	}
	return connect
}
