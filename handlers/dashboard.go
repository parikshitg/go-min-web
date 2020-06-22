package handlers

import (
	"html/template"
	"log"
	"net/http"
)

// Dashboard Controller
func Dashboard(w http.ResponseWriter, r *http.Request) {

	page, err := template.ParseFiles("templates/dashboard.html")
	if err != nil {
		log.Fatal("ParseFiles: ", err)
	}

	// Read All users
	rows, err := Db.Query(`SELECT id, username, password, created_at FROM users`)
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
