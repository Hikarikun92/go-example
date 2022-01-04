package main

import (
	"fmt"
	"go-example/user"
)

func main() {
	repository := user.RepositoryImpl{}
	service := user.Service{Repository: repository}

	fmt.Println(service.FindAll())
	fmt.Println(service.FindCredentialsByUsername("user1"))
	fmt.Println(service.FindCredentialsByUsername("user2"))
	fmt.Println(service.FindCredentialsByUsername("user3"))
}
