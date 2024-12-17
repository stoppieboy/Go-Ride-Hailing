package User

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/rapido/Database"
	"github.com/rapido/db"
	"golang.org/x/crypto/bcrypt"
)

var jwt_secret = "secret"

func registerUser(w http.ResponseWriter, r *http.Request) {
	client := Database.PrismaClient.Client
	ctx := Database.PrismaClient.Context
	var user User
	json.NewDecoder(r.Body).Decode(&user)
	fmt.Printf("Name: %v, Email: %v, Password: %v\n", user.Name, user.Email, user.Password)
	hash, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hash)

	createdUser, err := client.User.CreateOne(
		db.User.Name.Set(user.Name),
		db.User.Email.Set(user.Email),
		db.User.Password.Set(user.Password),
	).Exec(ctx)
	if err != nil {
		fmt.Println(err)
		fmt.Fprint(w, err)
		return
	}
	result, _ := json.MarshalIndent(createdUser, "", "  ")
	fmt.Println(string(result))
	fmt.Fprint(w, string(result))
}

func loginUser(w http.ResponseWriter, r *http.Request) {
	client := Database.PrismaClient.Client
	ctx := Database.PrismaClient.Context
	var user User
	json.NewDecoder(r.Body).Decode(&user)

	dbResult, err := client.User.FindUnique(
		db.User.Email.Equals(user.Email),
	).Exec(ctx)

	if err != nil {
		fmt.Println(err)
		fmt.Fprint(w, err)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(dbResult.Password), []byte(user.Password)); err != nil {
		fmt.Println(err)
		fmt.Fprint(w, err)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": dbResult.ID,
	})

	tokenString, err := token.SignedString([]byte(jwt_secret))

	if err != nil {
		fmt.Println(err)
		fmt.Fprint(w, err)
		return
	}

	response, _ := json.MarshalIndent(LoginResponse{
		Token:   tokenString,
		Success: true,
	}, "", "	")

	fmt.Println(string(response))
	fmt.Println("Token: %v", tokenString)
	fmt.Fprint(w, string(response))

}

func test(w http.ResponseWriter, r *http.Request) {

}