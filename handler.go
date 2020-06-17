package main

import (
	"html/template"
	"log"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {

	page, err := template.ParseFiles("templates/login.html")
	if err != nil {
		log.Fatal("ParseFiles: ", err)
	}

	err = page.Execute(w, nil)
	if err != nil {
		log.Fatal("Execute:", err)
	}
}

func Register(w http.ResponseWriter, r *http.Request) {

	page, err := template.ParseFiles("templates/register.html")
	if err != nil {
		log.Fatal("ParseFiles :", err)
	}

	err = page.Execute(w, nil)
	if err != nil {
		log.Fatal("Execute:", err)
	}
}

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
