package redis

import (
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

const (
	localDial = "localhost:6379"
	redisURL  = "REDIS_URL"
	redis     = "redis"
	users     = "users"
)

var (
	dockerDial = os.Getenv(redisURL)
)

func getDial() string {
	var dial string
	if os.Getenv(redisURL) == redis {
		dial = os.Getenv(redisURL)
	} else {
		dial = localDial
	}
	return dial
}

// Users is a slice of User
type Users []User

var connect *gore.Conn

// SetKey sets a redis key
func SetKey(key string, data interface{}) {
	connect := Connect()
	if data == nil {
		log.Println("should not be nil")
	}
	rep, err := gore.NewCommand("SET", key, data).Run(connect)
	if err != nil || !rep.IsOk() {
		log.Fatal(err, "not ok")
	}
	defer connect.Close()
}

// GetKey returns a redis value
func GetKey(key string) (*gore.Reply, error) {
	connect := Connect()
	reply, err := gore.NewCommand("GET", key).Run(connect)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

// GetKeys using KEYS redis command with pattern
func GetKeys(pattern string) (*gore.Reply, error) {
	connect := Connect()
	reply, err := gore.NewCommand("KEYS", pattern).Run(connect)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

// SetHash sets a redis hash set
func SetHash(key string, data model.UserModel) {
	connect := Connect()
	rep, err := gore.NewCommand("HMSET", key, "id", data.ID, "firstname", data.FirstName, "lastname", data.LastName, "email", data.Email, "gender", data.Gender).Run(connect)
	if err != nil || !rep.IsOk() {
		log.Fatal(err, "not ok")
	}
	defer connect.Close()
}

// GetHashAll returns redis hash
func GetHashAll(key string) (*gore.Reply, error) {
	connect := Connect()
	rep, err := gore.NewCommand("HGETALL", key).Run(connect)
	if err != nil {
		log.Fatal(err, "not ok")
	}
	defer connect.Close()
	return rep, err
}

// Connect returns redis connection
func Connect() *gore.Conn {
	var err error
	connect, err = gore.Dial(getDial())
	if err != nil {
		return nil
	}
	return connect
}
