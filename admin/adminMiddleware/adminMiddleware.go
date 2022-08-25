package adminMiddleware

import (
	"fmt"
	"log"
	"net/http"
	"time"

	e "github.com/ijasmoopan/usermanagement/entities"
	d "github.com/ijasmoopan/usermanagement/repository"

	"github.com/golang-jwt/jwt"
)
const key = "secretKey"

func IsAuthorized(handler http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){

		cookie, err := r.Cookie("jwt")
		if err != nil {
			fmt.Println("Error while fetching cookie.. Redirecting to login")
			http.Redirect(w, r, "/adminlogin", http.StatusSeeOther)
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

		var user e.UserEdit

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
		http.Redirect(w, r, "/userlist", http.StatusSeeOther)
	})
}

func DeleteToken(handler http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		t := http.Cookie{
			Name: "jwtAdminToken",
			Value:"",
			MaxAge: -1,
		}
		http.SetCookie(w, &t)
		fmt.Println("Cookie deleted")
		handler.ServeHTTP(w, r)
	})
}

func NoCache(handler http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Cache-Control", "no-cache, private, max-age=0 no-store, must-revalidate, post-check=0, pre-check=0")
		w.Header().Set("Expires", time.Unix(0, 0).Format(http.TimeFormat))
		w.Header().Set("Pragma", "no-cache")
		w.Header().Set("X-Accel-Expires", "0")
		fmt.Println("Cache cleared")
		handler.ServeHTTP(w, r)
	})
}