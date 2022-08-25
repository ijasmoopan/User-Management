package user

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	d "github.com/ijasmoopan/usermanagement/repository"
	m "github.com/ijasmoopan/usermanagement/entities"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/chi/v5"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	
	//-------------Authenticating-----------------
	const key = "secretKey"

	cookie, err := r.Cookie("jwt")
	if err != nil {
		fmt.Println("Error while fetching cookie.. Redirecting to login")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
	fmt.Println("Cookie: ", cookie)
	token, err := jwt.ParseWithClaims(cookie.Value, &jwt.StandardClaims{}, func(token *jwt.Token)(interface{}, error){
		return []byte(key), nil
	})
	if err != nil {
		log.Fatal("Unauthorized")
	}
	claims := token.Claims.(*jwt.StandardClaims)

	var user m.User

	db := d.GetDatabase()
	defer d.CloseDatabase(db)

	row := db.QueryRow("SELECT id, username, password FROM userregistration WHERE id=$1", claims.Issuer)
	if err := row.Scan(&user.UserID, &user.Username, &user.Password); err != nil {
		log.Fatal("Unauthenticated while querying")
	}

	name := chi.URLParam(r, "user")
	data := map[string]interface{}{
		"user": name,
	}

	tmp, _ := template.ParseFiles("templates/homepage.html")
	tmp.Execute(w, data)
}