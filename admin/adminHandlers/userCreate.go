package adminHandlers

import (
	"html/template"
	"log"
	"net/http"
	d "github.com/ijasmoopan/usermanagement/repository"
)

func CreateHandler(w http.ResponseWriter, r *http.Request){
	tpl, _ := template.ParseFiles("admin/templates/createUser.html")
	tpl.Execute(w, nil)
}
func CreateUserHandler(w http.ResponseWriter, r *http.Request){
	db := d.GetDatabase()
	defer d.CloseDatabase(db)

	r.ParseForm()
	username := r.Form.Get("username")
	password := r.Form.Get("password")

	_, err := db.Exec("INSERT INTO userregistration (username, password) VALUES ($1, $2)", username, password)
	if err != nil {
		log.Fatal(err)
	}
	http.Redirect(w, r, "/userlist", http.StatusSeeOther)
}