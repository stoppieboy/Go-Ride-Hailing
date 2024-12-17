package User

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

func InitializeUserService(userRouter *mux.Router) {
	fmt.Println("User service starting...")
	userRouter.HandleFunc("/register", registerUser).Methods("POST")
	userRouter.HandleFunc("/login", loginUser).Methods("POST")
	userRouter.HandleFunc("/test", test).Methods("GET")

	http.Handle("/", userRouter)
	fmt.Println("User service listening on port 3000")
	http.ListenAndServe(":3000", nil)
}