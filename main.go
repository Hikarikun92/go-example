package main

import (
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

	http.HandleFunc("/users", controller.FindAll)

	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Panicln("Error starting server", err)
	}
}
