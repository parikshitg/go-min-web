package main

import (
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/sessions"
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

// initialize session key
var key = []byte("super-secret-key")
var store = sessions.NewCookieStore(key)

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

				// setting up a session
				session, err := store.Get(r, "cookie-name")
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

// Logout function
func Logout(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")

	session.Values["authenticated"] = false
	session.Save(r, w)
}

// Checks if user is authenticated (middleware)
func AuthenticatedUser(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		session, _ := store.Get(r, "cookie-name")

		cookie, ok := session.Values["authenticated"]

		log.Println("AUTHENTICATEDUSER", r.URL.Path, "cookie : ", cookie, "ok : ", ok)

		if cookie == true && ok {
			f(w, r)
		} else {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		}

	}
}

// checks unauthenticated user (middleware)
func UnauthenticatedUser(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		session, _ := store.Get(r, "cookie-name")

		cookie, ok := session.Values["authenticated"]

		log.Println("UNAUTHENTICATEDUSER", r.URL.Path, " cookie : ", cookie, "ok : ", ok)

		if cookie == true && ok {
			http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
		} else {
			f(w, r)
		}
	}
}
