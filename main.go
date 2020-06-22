package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {

	var err error
	// Open a database
	db, err = sql.Open("mysql", "root:parikshitg@tcp(127.0.0.1:3306)/testingdb")
	if err != nil {
		fmt.Println("Db Open Error:", err)
	}

	defer db.Close()

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

	if _, err := db.Exec(query); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully Create Table")

	http.Handle("/login", http.HandlerFunc(UnauthenticatedUser(Login)))
	http.Handle("/logout", http.HandlerFunc(AuthenticatedUser(Logout)))
	http.Handle("/register", http.HandlerFunc(UnauthenticatedUser(Register)))
	http.Handle("/dashboard", http.HandlerFunc(AuthenticatedUser(Dashboard)))

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	http.ListenAndServe(":8080", nil)
}
