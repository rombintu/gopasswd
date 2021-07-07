package routes

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/rombintu/gopassic.git/crypt"
	"github.com/rombintu/gopassic.git/database"
	"github.com/rombintu/gopassic.git/models"
)

func Index(res http.ResponseWriter, req *http.Request) {
	template, err := template.ParseFiles("templates/index.html", "templates/header.html")
	if err != nil {
		fmt.Fprintf(res, err.Error())
	}

	db := database.Get_db()
	var passwords []models.Passwords
	db.Find(&passwords)
	for i := 0; i < len(passwords); i++ {
		passwords[i].Pass = string(crypt.Decode_pass(passwords[i].Pass))
	}
	template.ExecuteTemplate(res, "index", passwords)
}

func Login(res http.ResponseWriter, req *http.Request) {
	template, err := template.ParseFiles("templates/login.html", "templates/header.html")
	if err != nil {
		fmt.Fprintf(res, err.Error())
	}

	template.ExecuteTemplate(res, "login", nil)
}

func Push_create(res http.ResponseWriter, req *http.Request) {

	var service_name string = req.FormValue("service")
	var email string = req.FormValue("email")
	var pass string = req.FormValue("pass")

	if service_name == "" || email == "" || pass == "" {
		http.Redirect(res, req, "/create", http.StatusSeeOther)
		return
	}

	enpass := crypt.Encode_pass([]byte(pass))
	db := database.Get_db()
	db.Create(&models.Passwords{Service: service_name, Email: email, Pass: enpass})

	http.Redirect(res, req, "/", http.StatusSeeOther)
}
func Create(res http.ResponseWriter, req *http.Request) {
	template, err := template.ParseFiles("templates/create.html", "templates/header.html")
	if err != nil {
		fmt.Fprintf(res, err.Error())
	}

	template.ExecuteTemplate(res, "create", nil)
}
