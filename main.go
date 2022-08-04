package main

import (
	"database/sql"
	"github.com/Hikarikun92/go-example/post"
	postRest "github.com/Hikarikun92/go-example/post/rest"
	"github.com/Hikarikun92/go-example/security"
	securityRest "github.com/Hikarikun92/go-example/security/rest"
	"github.com/Hikarikun92/go-example/user"
	userRest "github.com/Hikarikun92/go-example/user/rest"
	"github.com/Hikarikun92/go-example/util"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	log.Println("Initializing application")
	config := util.LoadConfigFromEnvironment()

	log.Println("Connecting to the database")
	db, err := sql.Open("mysql", config.GetDataSourceName())
	if err != nil {
		log.Panicln("Error connecting to database", err)
	}
	if err := db.Ping(); err != nil {
		log.Panicln("Error connecting to database", err)
	}

	log.Println("Creating service handlers")
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

	log.Println("Application initialized")
	if err := http.ListenAndServe(config.GetServerAddress(), router); err != nil {
		log.Panicln("Error starting server", err)
	}
}
