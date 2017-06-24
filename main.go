package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jbelmont/docker-workshop/model"
	"github.com/jbelmont/docker-workshop/routes"
)

func getRouter() *mux.Router {
	return routes.NewRouter()
}

func init() {
	session := model.InitDB()
	model.CreateInitDocument(session)
}

func main() {
	router := getRouter()

	log.Fatal(http.ListenAndServe(":3000", router))
}
