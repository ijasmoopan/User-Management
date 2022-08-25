package adminHandlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	e "github.com/ijasmoopan/usermanagement/entities"
	d "github.com/ijasmoopan/usermanagement/repository"

	"github.com/go-chi/chi/v5"
)

func EditHandler(w http.ResponseWriter, r *http.Request){

	db := d.GetDatabase()
	defer d.CloseDatabase(db)

	var user e.UserEdit
	userid := chi.URLParam(r, "id")
	id, err := strconv.Atoi(userid)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("got parameter")
	row := db.QueryRow("SELECT id, username, password FROM userregistration WHERE id=$1", id)
	if err := row.Scan(&user.UserID, &user.Username, &user.Password); err != nil {
		log.Fatal("Error: ", err)
	}
	fmt.Println("execute query")
	tpl, err := template.ParseFiles("admin/templates/editUser.html")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("template parsed")
	tpl.Execute(w, user)
}

func EditingUserHandler(w http.ResponseWriter, r *http.Request){
	db := d.GetDatabase()
	defer d.CloseDatabase(db)

	userid := chi.URLParam(r, "id")
	id, _ := strconv.Atoi(userid)
	
	r.ParseForm()
	username := r.Form.Get("username")
	password := r.Form.Get("password")

	fmt.Println("Got form data")

	_, err := db.Exec("UPDATE userregistration SET username=$1, password=$2 WHERE id=$3", username, password, id)
	if err != nil {
		log.Fatal(err)
	}
	http.Redirect(w, r, "/userlist", http.StatusSeeOther)
	return
}


