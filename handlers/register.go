package handlers

import (
	"html/template"
	"log"
	"net/http"
)

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

		if username == "" || password == "" || password2 == "" {

			flash.Message = "Fields Can not be empty!!"
			log.Println(flash.Message)

			err = page.Execute(w, flash)
			if err != nil {
				log.Fatal("Execute:", err)
			}
			return
		}

		dbusername, dbpassword := ReadUser(username, password)

		if dbusername != "" || dbpassword != "" {

			flash.Message = "User Already Registered!!"
			log.Println(flash.Message)

			err = page.Execute(w, flash)
			if err != nil {
				log.Fatal("Execute:", err)
			}
			return
		}

		if username == dbusername {
			flash.Message = "This username is taken."
			log.Println(flash.Message)

			err = page.Execute(w, flash)
			if err != nil {
				log.Fatal("Execute:", err)
			}
			return
		}

		if password != password2 {
			flash.Message = "Passwords Doesn't Match !!"
			log.Println(flash.Message)

			err = page.Execute(w, flash)
			if err != nil {
				log.Fatal("Execute:", err)
			}
			return
		}

		CreateUser(username, password)
		flash.Message = "Registered Successfully."
		log.Println(flash.Message)

		// Checking if form fields are empty
		// if username == "" || password == "" || password2 == "" {

		// 	flash.Message = "Fields Can not be empty!!"
		// 	log.Println(flash.Message)
		// } else {

		// 	// Check if user already present in database
		// 	dbusername, dbpassword := ReadUser(username, password)

		// 	if dbusername != "" || dbpassword != "" {

		// 		flash.Message = "User Already Registered!!"
		// 		log.Println(flash.Message)
		// 	} else {

		// 		if username != dbusername {

		// 			if password == password2 {

		// 				CreateUser(username, password)
		// 				flash.Message = "Registered Successfully."
		// 				log.Println(flash.Message)
		// 			} else {
		// 				flash.Message = "Passwords Doesn't Match !!"
		// 				log.Println(flash.Message)
		// 			}
		// 		} else {
		// 			flash.Message = "This username is taken."
		// 			log.Println(flash.Message)
		// 		}
		// 	}
		// }
	}

	err = page.Execute(w, flash)
	if err != nil {
		log.Fatal("Execute:", err)
	}
}
