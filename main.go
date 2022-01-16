package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"go-example/post"
	postRest "go-example/post/rest"
	"go-example/user"
	userRest "go-example/user/rest"
	"log"
	"net/http"
)

func main() {
	db, err := sql.Open("mysql", "blog_backend_user:blog123@tcp(localhost:3306)/blog_backend_go?parseTime=true")
	if err != nil {
		log.Panicln("Error connecting to database", err)
	}
	if err := db.Ping(); err != nil {
		log.Panicln("Error connecting to database", err)
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userFacade := userRest.NewFacade(userService)
	userController := userRest.NewController(userFacade)

	postRepository := post.NewRepository(db)
	postService := post.NewService(postRepository)
	postFacade := postRest.NewFacade(postService)
	postController := postRest.NewController(postFacade)

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/users", userController.FindAll).Methods(http.MethodGet)
	router.HandleFunc("/users/{userId}/posts", postController.FindByUserId).Methods(http.MethodGet)
	router.HandleFunc("/posts/{id}", postController.FindById).Methods(http.MethodGet)

	if err := http.ListenAndServe("localhost:8080", router); err != nil {
		log.Panicln("Error starting server", err)
	}
}
