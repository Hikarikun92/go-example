package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"go-example/post"
	postRest "go-example/post/rest"
	"go-example/security"
	securityRest "go-example/security/rest"
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

	jwtService := security.NewJwtService(config, userRepository)
	securityService := security.NewService(userRepository, jwtService)
	securityFacade := securityRest.NewFacade(securityService)
	securityController := securityRest.NewController(securityFacade)

	router := mux.NewRouter().StrictSlash(true)
	router.Use(jwtService.AuthenticationMiddleware())

	router.HandleFunc("/users", userController.FindAll).Methods(http.MethodGet)
	router.HandleFunc("/users/{userId}/posts", postController.FindByUserId).Methods(http.MethodGet)
	router.HandleFunc("/posts", postController.Create).Methods(http.MethodPost).Headers("Content-Type", "application/json")
	router.HandleFunc("/posts/{id}", postController.FindById).Methods(http.MethodGet)
	router.HandleFunc("/login", securityController.Login).Methods(http.MethodPost).Headers("Content-Type", "application/json")

	if err := http.ListenAndServe(config.GetServerAddress(), router); err != nil {
		log.Panicln("Error starting server", err)
	}
}
