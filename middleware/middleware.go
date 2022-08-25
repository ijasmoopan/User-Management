package middleware

import (
	"fmt"
	"log"
	"net/http"
	// "time"

	d "github.com/ijasmoopan/usermanagement/repository"
	e "github.com/ijasmoopan/usermanagement/entities"

	"github.com/golang-jwt/jwt"
	// "github.com/go-chi/jwtauth/v5"
)
const key = "secretKey"

func IsAuthorized(handler http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){

		cookie, err := r.Cookie("jwt")
		if err != nil {
			fmt.Println("Error while fetching cookie.. Redirecting to login")
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		fmt.Println("Cookie: ", cookie)
		token, err := jwt.ParseWithClaims(cookie.Value, &jwt.StandardClaims{}, func(token *jwt.Token)(interface{}, error){
			return []byte(key), nil
		})
		if err != nil {
			log.Fatal("Unauthorized")
		}
		claims := token.Claims.(*jwt.StandardClaims)

		var user e.User

		db := d.GetDatabase()
		defer d.CloseDatabase(db)

		row := db.QueryRow("SELECT id, username, password FROM userregistration WHERE id=$1", claims.Issuer)
		if err := row.Scan(&user.UserID, &user.Username, &user.Password); err != nil {
			log.Fatal("Unauthenticated while querying")
		}
		handler.ServeHTTP(w	, r)
	})
}

func HaveToken( handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){

		fmt.Println("Have Token Middleware")
		cookie, err := r.Cookie("jwt")
		if err != nil {
			handler.ServeHTTP(w, r)
			return
		}
		token, err := jwt.ParseWithClaims(cookie.Value, &jwt.StandardClaims{}, func(token *jwt.Token)(interface{}, error){
			return []byte(key), nil
		})
		if err != nil {
			handler.ServeHTTP(w, r)
			return
		}
		claims := token.Claims.(*jwt.StandardClaims)

		var user e.UserEdit
		db := d.GetDatabase()
		defer d.CloseDatabase(db)

		row := db.QueryRow("SELECT id, username, password FROM userregistration WHERE id=$1", claims.Issuer)
		if err := row.Scan(&user.UserID, &user.Username, &user.Password); err != nil {
			handler.ServeHTTP(w, r)
			return
		}
		http.Redirect(w, r, "/home/"+user.Username, http.StatusSeeOther)
		return
	})
}

func DeleteToken(handler http.Handler) http.Handler {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request){
		fmt.Println("Middleware- Delete Token")
		t := http.Cookie{
			Name: "jwt",
			Value: "",
			MaxAge: -1,
			Path: "/login",
		}
		http.SetCookie(w, &t)
		ms := http.Cookie{
			Name: "mysession",
			Value: "",
			MaxAge: -1,
			Path: "/login",
		}
		http.SetCookie(w, &ms)
		sn := http.Cookie{
			Name: "session-name",
			Value: "",
			MaxAge: -1,
			Path: "/login",
		}
		http.SetCookie(w, &sn)
		fmt.Println("Cookie deleted")
		handler.ServeHTTP(w, r)
	})
}

func NoCache(handler http.Handler) http.Handler{
	fn := func(w http.ResponseWriter, r *http.Request){
		// Setting headers
		// handler.ServeHTTP(w, r)

		w.Header().Set("Cache-Control", "no-cache, private, max-age=0 no-store, must-revalidate, post-check=0, pre-check=0")
		w.Header().Set("Expires", "0")
		w.Header().Set("Pragma", "no-cache")
		w.Header().Set("X-Accel-Expires", "0")
		fmt.Println("Cache cleared")
		handler.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}