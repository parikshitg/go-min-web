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

type Data struct {
	Message string
}

var data Data

// Create user function creates a new user in users table
func CreateUser(username, password string) {

	createdAt := time.Now()

	_, err := db.Exec(`INSERT INTO userss (username, password, created_at) VALUES (?, ?, ?)`, username, password, createdAt)
	if err != nil {
		log.Println("Insert Error : ", err)
	}
}

// Read user reads a user from users table
func ReadUser(username, password string) (string, string) {

	var (
		usernamedb string
		passworddb string
	)

	query := "SELECT username, password FROM userss WHERE username = ?"
	if err := db.QueryRow(query, username).Scan(&usernamedb, &passworddb); err != nil {
		log.Println("Read User Error : ", err)
	}

	return usernamedb, passworddb
}

// Login Controller
func Login(w http.ResponseWriter, r *http.Request) {

	page, err := template.ParseFiles("templates/login.html")
	if err != nil {
		log.Fatal("ParseFiles: ", err)
	}

	if r.Method == http.MethodPost {

		username := r.FormValue("username")
		password := r.FormValue("password")

		dbusername, dbpassword := ReadUser(username, password)

		log.Println("NOt exists : ", dbusername, dbpassword)

		if username == dbusername && password == dbpassword {

			http.Redirect(w, r, "/dashboard", 301)

			data.Message = "You have been logged in Successfully."
			log.Println(data.Message)

		} else {
			data.Message = "Invalid username or password!!"
			log.Println(data.Message)
		}
	}

	err = page.Execute(w, data)
	if err != nil {
		log.Fatal("Execute:", err)
	}
}

// Register Controller
func Register(w http.ResponseWriter, r *http.Request) {

	page, err := template.ParseFiles("templates/register.html")
	if err != nil {
		log.Fatal("ParseFiles :", err)
	}

	if r.Method == http.MethodPost {

		username := r.FormValue("username")
		password := r.FormValue("password")
		password2 := r.FormValue("password2")

		dbusername, _ := ReadUser(username, password)

		if username != dbusername {

			if password == password2 {

				CreateUser(username, password)

				data.Message = "Registered Successfully."
				log.Println(data.Message)

			} else {
				data.Message = "Passwords Doesn't Match !!"
				log.Println(data.Message)
			}
		} else {
			data.Message = "Already Registered!"
			log.Println(data.Message)
		}

	}

	err = page.Execute(w, data)
	if err != nil {
		log.Fatal("Execute:", err)
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
