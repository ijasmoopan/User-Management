package adminHandlers

import (
	// "fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
	e "github.com/ijasmoopan/usermanagement/entities"
	d "github.com/ijasmoopan/usermanagement/repository"
)

func SearchHandler(w http.ResponseWriter, r *http.Request){
	db := d.GetDatabase()
	defer d.CloseDatabase(db)

	var user e.UserList
	userDetails := []interface{}{}

	r.ParseForm()
	username := r.Form.Get("search")
	// username = fmt.Sprintf(username, "%")
	var sb strings.Builder
	sb.WriteString("%")
	sb.WriteString(username)
	sb.WriteString("%")
	username = sb.String()

	// -----------------------sprintf--------------------------

	rows, err := db.Query("SELECT id, username, status FROM userregistration WHERE username ILIKE $1", username)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next(){
		if err := rows.Scan(&user.UserID, &user.Username, &user.Status); err != nil {
			log.Fatal(err)
		}
		userDetails = append(userDetails, user)
	}
	tpl, _ := template.ParseFiles("admin/templates/userList.html")
	tpl.Execute(w, map[string]interface{}{"user": userDetails})
}
