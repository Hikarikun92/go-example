package main

import (
	"github.com/gorilla/mux"
	"go-example/user"
	"go-example/user/rest"
	"log"
	"net/http"
)

func main() {
	repository := user.NewRepository()
	service := user.NewService(repository)
	facade := rest.NewFacade(service)
	controller := rest.NewController(facade)

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/users", controller.FindAll).Methods(http.MethodGet)

	err := http.ListenAndServe("localhost:8080", router)
	if err != nil {
		log.Panicln("Error starting server", err)
	}
}
