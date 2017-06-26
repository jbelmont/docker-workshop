package redis

import (
	"os"

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

func ConnectRedis() {
	dial := os.Getenv("REDIS_URL")
	conn, err := gore.Dial(dial)
	if err != nil {
		return
	}
	defer conn.Close()
}
