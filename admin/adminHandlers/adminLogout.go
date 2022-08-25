package adminHandlers

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
)

func AdminLogOut(w http.ResponseWriter, r *http.Request){

	fmt.Println("Deleting cookie..")
	cookie := http.Cookie{
		Name: "jwt",
		Value: "",
		Expires: time.Now().Add(-time.Hour),
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)
	fmt.Println("Cookie deleted, Logouting..")
	// http.Redirect(w, r, "/adminlogin", http.StatusSeeOther)

	tmp, _ := template.ParseFiles("admin/templates/adminLoginPage.html")
	tmp.Execute(w, nil)
	return
}