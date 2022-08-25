package adminHandlers

import (
	"log"
	"net/http"
	d "github.com/ijasmoopan/usermanagement/repository"

	"github.com/go-chi/chi/v5"
)
func StatusHandler(w http.ResponseWriter, r *http.Request) {

	db := d.GetDatabase()
	defer d.CloseDatabase(db)

	userid := chi.URLParam(r, "id")
	status := chi.URLParam(r, "status")

	if status != "true"{
		_, err := db.Exec("UPDATE userregistration SET status=$1 WHERE id=$2", "true", userid)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		_, err := db.Exec("UPDATE userregistration SET status=$1 WHERE id=$2", "false", userid)
		if err != nil {
			log.Fatal(err)
		}
	}
	
	http.Redirect(w, r, "/userlist", http.StatusSeeOther)
}
