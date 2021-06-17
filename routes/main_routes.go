package routes

import (
	"fmt"
	"net/http"
	"text/template"
)

func Index(res http.ResponseWriter, req *http.Request) {
	template, err := template.ParseFiles("templates/index.html", "templates/header.html")
	if err != nil {
		fmt.Fprintf(res, err.Error())
	}

	template.ExecuteTemplate(res, "index", nil)
}

func Login(res http.ResponseWriter, req *http.Request) {
	template, err := template.ParseFiles("templates/login.html", "templates/header.html")
	if err != nil {
		fmt.Fprintf(res, err.Error())
	}

	template.ExecuteTemplate(res, "login", nil)
}

func Create(res http.ResponseWriter, req *http.Request) {
	template, err := template.ParseFiles("templates/create.html", "templates/header.html")
	if err != nil {
		fmt.Fprintf(res, err.Error())
	}

	template.ExecuteTemplate(res, "create", nil)
}
