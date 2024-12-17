package main

import (
	"fmt"

	"github.com/gorilla/mux"
	"github.com/rapido/Database"
	"github.com/rapido/User"
)

func main() {
	router := mux.NewRouter()
	Database.InitializeDB()
	User.InitializeUserService(router.PathPrefix("/user").Subrouter())
	fmt.Println("User service started")
}