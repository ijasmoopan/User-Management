package adminHandlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	e "github.com/ijasmoopan/usermanagement/entities"
	d "github.com/ijasmoopan/usermanagement/repository"
)
func AccessingUserList(w http.ResponseWriter, r *http.Request){
	fmt.Println("Got the user list page")
	db := d.GetDatabase()
	defer d.CloseDatabase(db)

	var user e.UserList
	userDetails := []interface{}{}

	rows, err := db.Query("SELECT id, username, status FROM userregistration ORDER BY id")
	if err != nil {
		http.Error(w, "Can't fetch database values", http.StatusBadRequest)
		return
	}
	defer rows.Close()
	for rows.Next(){
		if err = rows.Scan(&user.UserID, &user.Username, &user.Status); err != nil {
			http.Error(w, "Can't fetch database values", http.StatusBadRequest)
			return
		}
		userDetails = append(userDetails, user)
	}
	tpl, err := template.ParseFiles("admin/templates/userList.html")
	if err != nil {
		log.Fatal(err)
	}
	err = tpl.Execute(w, map[string]interface{}{"user": userDetails})
	if err != nil {
		log.Fatal(err)
	}
}
