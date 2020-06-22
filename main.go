package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	h "github.com/parikshitg/go-min-web/handlers"
	"github.com/parikshitg/go-min-web/middlewares"
)

func main() {

	var err error
	// Open a database
	h.Db, err = sql.Open("mysql", "root:parikshitg@tcp(127.0.0.1:3306)/testingdb")
	if err != nil {
		fmt.Println("Db Open Error:", err)
	}

	defer h.Db.Close()

	fmt.Println("Successfully connected to Database.")

	// Create a database table
	query := `
            CREATE TABLE IF NOT EXISTS users (
                id INT AUTO_INCREMENT,
                username TEXT NOT NULL,
                password TEXT NOT NULL,
                created_at DATETIME,
                PRIMARY KEY (id)
            );`

	if _, err := h.Db.Exec(query); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully Create Table")

	http.Handle("/login", middlewares.UnauthenticatedUser(h.Login))
	http.Handle("/logout", middlewares.AuthenticatedUser(h.Logout))
	http.Handle("/register", middlewares.UnauthenticatedUser(h.Register))
	http.Handle("/dashboard", middlewares.AuthenticatedUser(h.Dashboard))

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	http.ListenAndServe(":8080", nil)
}
