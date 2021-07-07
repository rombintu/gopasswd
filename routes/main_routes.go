package routes

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/rombintu/gopassic.git/crypt"
	"github.com/rombintu/gopassic.git/csvman"
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

func Create_push(res http.ResponseWriter, req *http.Request) {

	var service string = req.FormValue("service")
	var url string = req.FormValue("url")
	var email string = req.FormValue("email")
	var pass string = req.FormValue("pass")

	if service == "" || url == "" || email == "" || pass == "" {
		http.Redirect(res, req, "/create", http.StatusSeeOther)
		return
	}

	enpass := crypt.Encode_pass([]byte(pass))
	db := database.Get_db()
	db.Create(&models.Passwords{Service: service, Url: url, Email: email, Pass: enpass})

	http.Redirect(res, req, "/", http.StatusSeeOther)
}

func Create(res http.ResponseWriter, req *http.Request) {
	template, err := template.ParseFiles("templates/create.html", "templates/header.html")
	if err != nil {
		fmt.Fprintf(res, err.Error())
	}

	template.ExecuteTemplate(res, "create", nil)
}

func Import_passwords(res http.ResponseWriter, req *http.Request) {
	template, err := template.ParseFiles("templates/import.html", "templates/header.html")
	if err != nil {
		fmt.Fprintf(res, err.Error())
	}

	template.ExecuteTemplate(res, "import", nil)
}

func Import_passwords_push(res http.ResponseWriter, req *http.Request) {
	file, _, err := req.FormFile("file")
	if file == nil {
		http.Redirect(res, req, "/import", http.StatusSeeOther)
		return
	}
	if err != nil {
		panic(err)
	}
	defer file.Close()

	data := csvman.Parse_csv(file)
	db := database.Get_db()

	for i := 0; i < len(data); i++ {
		if data[i][0] == "name" {
			continue
		}
		enpass := crypt.Encode_pass([]byte(data[i][3]))
		db.Create(&models.Passwords{Service: data[i][0], Url: data[i][1], Email: data[i][2], Pass: enpass})
	}

	http.Redirect(res, req, "/", http.StatusSeeOther)
}
