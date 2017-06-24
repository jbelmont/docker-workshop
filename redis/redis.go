package redis

import "github.com/keimoon/gore"

// User Struct is the model for the app
type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Gender    string `json:"gender"`
}

func ConnectRedis() {
	conn, err := gore.Dial("localhost:6379")
	if err != nil {
		return
	}
	defer conn.Close()
}
