package user

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	e "github.com/ijasmoopan/usermanagement/entities"
	d "github.com/ijasmoopan/usermanagement/repository"

	"github.com/go-chi/chi/v5"
	"golang.org/x/crypto/bcrypt"
)

func SignUpIndex(w http.ResponseWriter, r *http.Request) {

	msg := chi.URLParam(r, "msg")
	data := map[string]interface{}{
		"msg": msg,
	}
	tmp, _ := template.ParseFiles("templates/index.html")
	tmp.Execute(w, data)
}

func SignUpHandler(w http.ResponseWriter, r *http.Request) {

	// Connecting database
	db := d.GetDatabase()
	defer d.CloseDatabase(db)

	// Accessing username & password
	var user e.User
	// err := json.NewDecoder(r.Body).Decode(&user)

	r.ParseForm()
	user.Username = r.Form.Get("username")
	user.Password = r.Form.Get("password")
	fmt.Println("user: ", user.Username, user.Password)

	// Hashing password
	// user.Password = u.GenerateMD5HashPassword(user.Password)
	password, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 14)


	// Checking that username is already exist in database or not
	sqlStatement := `SELECT username FROM userregistration WHERE username=$1;`
	var username string
	row := db.QueryRow(sqlStatement, user.Username)
	err := row.Scan(&username)
	fmt.Println("username: ", username)

	if username == user.Username {
		fmt.Println("username: ", username)
		data := map[string]interface{}{
			"msg": "Username is already exist",
		}
		tmp, _ := template.ParseFiles("templates/index.html")
		tmp.Execute(w, data)
		return
	} else if err == sql.ErrNoRows {

		// Saving username & password to database
		sqlStatement = `INSERT INTO userregistration (username, password) VALUES ($1, $2) RETURNING id`
		err := db.QueryRow(sqlStatement, user.Username, password).Scan(&user.UserID)
		if err != nil {
			fmt.Println("Error: ", err)
			http.Error(w, "Can't fetch data from table of Database", http.StatusInternalServerError)
			return
		}
		fmt.Println("Registration completed... ID : ", user.UserID)

		data := map[string]interface{}{
			"msg": "Successfully Registered",
		}
		tpl, _ := template.ParseFiles("templates/loginpage.html")
		tpl.Execute(w, data)

		// http.Redirect(w, r, "/login", http.StatusSeeOther)
		// http.Redirect(w, r, "/home?"+userName, http.StatusSeeOther)
		return
	} else {
		log.Fatalln("Can't connect to table of Database", err)
	}
}

