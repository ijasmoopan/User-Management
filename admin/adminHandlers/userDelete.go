package adminHandlers

import (
	"log"
	"net/http"
	d "github.com/ijasmoopan/usermanagement/repository"

	"github.com/go-chi/chi/v5"
)
func DeleteHandler(w http.ResponseWriter, r *http.Request) {

	db := d.GetDatabase()
	defer d.CloseDatabase(db)

	userid := chi.URLParam(r, "id")

	_, err := db.Exec("DELETE FROM userregistration WHERE id=$1", userid)
	if err != nil {
		log.Fatal(err)
	}
	http.Redirect(w, r, "/userlist", http.StatusSeeOther)
}