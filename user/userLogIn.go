package user

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"

	e "github.com/ijasmoopan/usermanagement/entities"
	d "github.com/ijasmoopan/usermanagement/repository"
	// u "users/useCases"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	// "github.com/go-chi/jwtauth/v5"
)

const key = "secretKey"

func LogInIndex(w http.ResponseWriter, r *http.Request) {

	tmp, _ := template.ParseFiles("templates/loginpage.html")
	tmp.Execute(w, nil)
}

func LogInHandler(w http.ResponseWriter, r *http.Request) {

	db := d.GetDatabase()
	defer d.CloseDatabase(db)

	r.ParseForm()
	var userAuth e.Authentication
	userAuth.Username = r.Form.Get("username")
	userAuth.Password = r.Form.Get("password")
	fmt.Println(userAuth)

	row := db.QueryRow("SELECT * FROM userregistration WHERE username=$1", userAuth.Username)
	var user e.User
	if err := row.Scan(&user.UserID, &user.Username, &user.Password, &user.Status); err != nil {
		if err == sql.ErrNoRows {
			log.Fatal("User not found")
		}
		log.Fatal("User not found: ", err)
	}
	if user.UserID == 0 {
		log.Fatal("User not found: id = 0")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userAuth.Password)); err != nil {

		data := map[string]interface{}{
			"msg": "Username or Password is incorrect",
		}
		tpl, _ := template.ParseFiles("templates/loginpage.html")
		tpl.Execute(w, data)
	}
	// ---------------Checking status-----------------------
	if !user.Status {

		data := map[string]interface{}{
			"msg": "You are blocked by Admin",
		}
		tpl, _ := template.ParseFiles("templates/loginpage.html")
		tpl.Execute(w, data)
		return
	}

	fmt.Println("NO ERROR")
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.UserID)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // 1 day
	})
	token, err := claims.SignedString([]byte(key))
	if err != nil {
		log.Fatal("Error in generating token")
	}
	fmt.Println("Storing in cookie")
	cookie := http.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)
	fmt.Println("Success: Token generated successfully")
	http.Redirect(w, r, "/home/"+userAuth.Username, http.StatusSeeOther)

}
