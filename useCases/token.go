package useCases

import (
	// "fmt"
	"log"
	// "strconv"
	"time"

	// "github.com/go-chi/jwtauth/v5"
	"github.com/golang-jwt/jwt"
	_ "github.com/lib/pq"
)

func GenerateToken(username string) (string, error) {
	var key = []byte("secretKey")

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(key)
	if err != nil {
		log.Fatalln("Something bad happening while creating tokenString", err)
	}
	return tokenString, err
}

// var tokenAuth *jwtauth.JWTAuth

// func GenerateToken(username string) *jwtauth.JWTAuth {

// 	tokenAuth = jwtauth.New("HS256", []byte("secret"), nil)
// 	// useridStr := strconv.Itoa(userid)

// 	// For debugging/example purposes, we generate and print
// 	// a sample jwt token with claims `user_id:123` here:
// 	_, tokenString, err := tokenAuth.Encode(map[string]interface{}{"user_id": username})
// 	if err != nil {
// 		log.Fatal("something bad happening while creating tokenString", err)
// 	}
// 	fmt.Printf("DEBUG: a sample jwt is %s\n\n", tokenString)
// 	return tokenAuth
// }