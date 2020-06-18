package main

import (
	"html/template"
	"log"
	"net/http"
	"time"
)

type User struct {
	Id        string
	Username  string
	Password  string
	CreatedAt time.Time
}

func CreateUser(username, password string) {

	createdAt := time.Now()

	_, err := db.Exec(`INSERT INTO userss (username, password, created_at) VALUES (?, ?, ?)`, username, password, createdAt)
	if err != nil {
		log.Fatal("Insert Error : ", err)
	}
}

func ReadUser(username, password string) (string, string) {

	query := "SELECT username, password FROM userss WHERE username = ?"
	if err := db.QueryRow(query, username).Scan(&username, &password); err != nil {
		log.Fatal("Read User Error : ", err)
	}

	return username, password
}

// Login Controller
func Login(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {

		username := r.FormValue("username")
		password := r.FormValue("password")

		dbusername, dbpassword := ReadUser(username, password)

		if username == dbusername && password == dbpassword {
			pages := template.Must(template.ParseFiles("templates/dashboard.html"))

			err := pages.Execute(w, nil)
			if err != nil {
				log.Fatal("Execute : ", err)
			}
		}

	} else {

		page, err := template.ParseFiles("templates/login.html")
		if err != nil {
			log.Fatal("ParseFiles: ", err)
		}

		err = page.Execute(w, nil)
		if err != nil {
			log.Fatal("Execute:", err)
		}
	}

}

// Register Controller
func Register(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {

		username := r.FormValue("username")
		password := r.FormValue("password")
		password2 := r.FormValue("password2")

		if password != password2 {
			log.Fatal("password not equal to password 2")
			return
		}

		CreateUser(username, password)

	} else {

		page, err := template.ParseFiles("templates/register.html")
		if err != nil {
			log.Fatal("ParseFiles :", err)
		}

		err = page.Execute(w, nil)
		if err != nil {
			log.Fatal("Execute:", err)
		}

	}
}

// Dashboard Controller
func Dashboard(w http.ResponseWriter, r *http.Request) {

	page, err := template.ParseFiles("templates/dashboard.html")
	if err != nil {
		log.Fatal("ParseFiles: ", err)
	}

	err = page.Execute(w, nil)
	if err != nil {
		log.Fatal("Execute: ", err)
	}
}
