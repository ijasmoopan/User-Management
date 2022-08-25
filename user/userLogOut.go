package user

import (
	"fmt"
	"net/http"
	"time"
)

func LogOutHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Deleting cookie..")
	cookie := http.Cookie{
		Name: "jwt",
		Value: "",
		Expires: time.Now().Add(-time.Hour),
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)
	fmt.Println("Cookie deleted, Logouting..")
	// tmp, _ := template.ParseFiles("loginpage.html")
	// tmp.Execute(w, r)
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
