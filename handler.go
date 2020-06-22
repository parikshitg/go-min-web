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
	CreatedAt string
}

type Flash struct {
	Message string
}

var flash Flash

// Create user function creates a new user in users table
func CreateUser(username, password string) {

	createdAt := time.Now()

	_, err := db.Exec(`INSERT INTO users (username, password, created_at) VALUES (?, ?, ?)`, username, password, createdAt)
	if err != nil {
		log.Println("Insert Error : ", err)
	}
}

// Read user reads a user from users table
func ReadUser(username, password string) (string, string) {

	var usernamedb, passworddb string

	query := "SELECT username, password FROM users WHERE username = ?"
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

		if username == "" || password == "" {

			flash.Message = "Fields can not be empty!!"
			log.Println(flash.Message)
		} else {

			dbusername, dbpassword := ReadUser(username, password)

			if username == dbusername && password == dbpassword {

				http.Redirect(w, r, "/dashboard", 301)

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

// Register Controller
func Register(w http.ResponseWriter, r *http.Request) {

	page, err := template.ParseFiles("templates/register.html")
	if err != nil {
		log.Fatal("ParseFiles :", err)
	}

	if r.Method == http.MethodPost {

		// post form values
		username := r.FormValue("username")
		password := r.FormValue("password")
		password2 := r.FormValue("password2")

		// Checking if form fields are empty
		if username == "" || password == "" || password2 == "" {

			flash.Message = "Fields Can not be empty!!"
			log.Println(flash.Message)
		} else {

			// Check if user already present in database
			dbusername, dbpassword := ReadUser(username, password)

			if dbusername != "" || dbpassword != "" {

				flash.Message = "User Already Registered!!"
				log.Println(flash.Message)
			} else {

				if username != dbusername {

					if password == password2 {

						CreateUser(username, password)
						flash.Message = "Registered Successfully."
						log.Println(flash.Message)
					} else {
						flash.Message = "Passwords Doesn't Match !!"
						log.Println(flash.Message)
					}
				} else {
					flash.Message = "This username is taken."
					log.Println(flash.Message)
				}
			}
		}
	}

	err = page.Execute(w, flash)
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

	// Read All users
	rows, err := db.Query(`SELECT id, username, password, created_at FROM users`)
	if err != nil {
		log.Println("Error Read ALl : ", err)
	}

	var users []User

	for rows.Next() {
		var u User
		err := rows.Scan(&u.Id, &u.Username, &u.Password, &u.CreatedAt)
		if err != nil {
			log.Println("Errror Scan All: ", err)
		}
		users = append(users, u)
	}
	if err := rows.Err(); err != nil {
		log.Fatal("Rows Error : ", err)
	}

	// log.Printf("%#v", users)

	err = page.Execute(w, users)
	if err != nil {
		log.Fatal("Execute: ", err)
	}
}
