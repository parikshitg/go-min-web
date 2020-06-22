package handlers

import (
	"log"
	"time"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"
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

var Db *sql.DB

// Create user function creates a new user in users table
func CreateUser(username, password string) {

	createdAt := time.Now()

	_, err := Db.Exec(`INSERT INTO users (username, password, created_at) VALUES (?, ?, ?)`, username, password, createdAt)
	if err != nil {
		log.Println("Insert Error : ", err)
	}
}

// Read user reads a user from users table
func ReadUser(username, password string) (string, string) {

	var dbusername, dbpassword string

	query := "SELECT username, password FROM users WHERE username = ?"
	if err := Db.QueryRow(query, username).Scan(&dbusername, &dbpassword); err != nil {
		log.Println("Read User Error : ", err)
	}

	return dbusername, dbpassword
}
