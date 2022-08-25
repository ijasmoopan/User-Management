package adminHandlers

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
)

const key = "secretKey"

func AdminLogin(w http.ResponseWriter, r *http.Request) {
	
	tpl, _ := template.ParseFiles("admin/templates/adminLoginPage.html")
	tpl.Execute(w, nil)
}
func AdminLoginValidation(w http.ResponseWriter, r *http.Request) {

	db := d.GetDatabase()
	defer d.CloseDatabase(db)

	r.ParseForm()
	var adminAuth e.Admin
	adminAuth.Username = r.Form.Get("username")
	adminAuth.Password = r.Form.Get("password")
	// adminAuth.Password = u.GenerateMD5HashPassword(adminAuth.Password)

	var admin e.UserEdit
	row := db.QueryRow("SELECT id, username, password FROM userregistration WHERE id=$1", 80)
	if err := row.Scan(&admin.UserID, &admin.Username, &admin.Password); err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No row were returned", err)
			data := map[string]interface{}{
				"msg": "Username or Password is incorrect",
			}
			tpl, _ := template.ParseFiles("admin/templates/adminLoginPage.html")
			tpl.Execute(w, data)
		}
		tpl, _ := template.ParseFiles("admin/templates/adminLoginPage.html")
		tpl.Execute(w, nil)
	}
	if admin.UserID == 0 {
		log.Fatal("User not found: id = 0")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(adminAuth.Password)); err != nil {
		data := map[string]interface{}{
			"msg": "Username or Password is incorrect",
		}
		tpl, _ := template.ParseFiles("admin/templates/adminLoginPage.html")
		tpl.Execute(w, data)
	}

	fmt.Println("NO ERROR")
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(admin.UserID)),
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
	http.Redirect(w, r, "/userlist", http.StatusSeeOther)
}
