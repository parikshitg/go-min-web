package main

import (
	"html/template"
	"log"
	"net/http"
)

func Dashboard(w http.ResponseWriter, r *http.Request) {

	page, err := template.ParseFiles("dashboard.html")
	if err != nil {
		log.Fatal("ParseFiles: ", err)
	}

	err = page.Execute(w, nil)
	if err != nil {
		log.Fatal("Execute: ", err)
	}
}
