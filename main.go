package main

import (
	"github.com/gorilla/mux"
	"go-example/post"
	postRest "go-example/post/rest"
	"go-example/user"
	userRest "go-example/user/rest"
	"log"
	"net/http"
)

func main() {
	userRepository := user.NewRepository()
	userService := user.NewService(userRepository)
	userFacade := userRest.NewFacade(userService)
	userController := userRest.NewController(userFacade)

	postRepository := post.NewRepository()
	postService := post.NewService(postRepository)
	postFacade := postRest.NewFacade(postService)
	postController := postRest.NewController(postFacade)

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/users", userController.FindAll).Methods(http.MethodGet)
	router.HandleFunc("/users/{userId}/posts", postController.FindByUserId).Methods(http.MethodGet)
	router.HandleFunc("/posts/{id}", postController.FindById).Methods(http.MethodGet)

	err := http.ListenAndServe("localhost:8080", router)
	if err != nil {
		log.Panicln("Error starting server", err)
	}
}
