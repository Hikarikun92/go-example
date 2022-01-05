package main

import (
	"fmt"
	"go-example/user"
)

func main() {
	repository := user.NewRepository()
	service := user.NewService(repository)

	fmt.Println(service.FindAll())
	fmt.Println(service.FindCredentialsByUsername("user1"))
	fmt.Println(service.FindCredentialsByUsername("user2"))
	fmt.Println(service.FindCredentialsByUsername("user3"))
}
