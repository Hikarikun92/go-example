package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"go-example/post"
	postRest "go-example/post/rest"
	"go-example/security"
	"go-example/user"
	userRest "go-example/user/rest"
	"go-example/util"
	"log"
	"net/http"
)

func main() {
	config := util.LoadConfigFromEnvironment()

	db, err := sql.Open("mysql", config.GetDataSourceName())
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

	router.Use(security.AuthenticationMiddleware())

	router.HandleFunc("/users", userController.FindAll).Methods(http.MethodGet)
	router.HandleFunc("/users/{userId}/posts", postController.FindByUserId).Methods(http.MethodGet)
	router.HandleFunc("/posts/{id}", postController.FindById).Methods(http.MethodGet)

	if err := http.ListenAndServe(config.GetServerAddress(), router); err != nil {
		log.Panicln("Error starting server", err)
	}
}
