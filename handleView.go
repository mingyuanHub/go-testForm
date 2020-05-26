package main

import (
	"goTestForm/models"
	"html/template"
	"net/http"
)

func HomeView(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("views/home.html")
	t.Execute(w, "")
}

func ListView(w http.ResponseWriter, r *http.Request) {
	list := models.List()
	t, _ := template.ParseFiles("views/list.html")
	// fmt.Println(list)
	t.Execute(w, list)
}
