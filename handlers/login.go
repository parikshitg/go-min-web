package handlers

import (
	"html/template"
	"log"
	"net/http"

	s "github.com/parikshitg/go-min-web/session"
)

// Login Controller
func Login(w http.ResponseWriter, r *http.Request) {

	page, err := template.ParseFiles("templates/login.html")
	if err != nil {
		log.Fatal("ParseFiles: ", err)
	}

	if r.Method == http.MethodPost {

		username := r.FormValue("username")
		password := r.FormValue("password")

		if username == "" || password == "" {

			flash.Message = "Fields can not be empty!!"
			log.Println(flash.Message)
		} else {

			dbusername, dbpassword := ReadUser(username, password)

			if username == dbusername && password == dbpassword {

				// setting up a session
				session, err := s.Store.Get(r, "cookie-name")
				if err != nil {
					log.Println("Session Error:", err)
				}
				session.Values["authenticated"] = true
				session.Save(r, w)

				http.Redirect(w, r, "/dashboard", http.StatusSeeOther)

				log.Println("You have been logged in Successfully.")

			} else {
				flash.Message = "Invalid username or password!!"
				log.Println(flash.Message)
			}
		}
	}

	err = page.Execute(w, flash)
	if err != nil {
		log.Fatal("Execute:", err)
	}
}
