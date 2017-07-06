package routes

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/gorilla/mux"

	"github.com/jbelmont/docker-workshop/handlers"
)

type Route struct {
	Method      string
	Pattern     string
	Name        string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

type MyHandler struct {
	sync.Mutex
	count int
}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var count int
	h.Lock()
	h.count++
	count = h.count
	h.Unlock()

	fmt.Fprintf(w, "Visitor count: %d.", count)
}

// NewRouter retrieves a new router
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			PathPrefix("/api/v1/").
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}

var routes = Routes{
	Route{
		"GET",
		"/",
		"Index",
		handlers.UsersIndex,
	},
	Route{
		"GET",
		"/users",
		"Users",
		handlers.GetUsers,
	},
	Route{
		"POST",
		"/addUser",
		"AddUser",
		handlers.AddUser,
	},
}
