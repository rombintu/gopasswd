package routes

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/rombintu/gopasswd.git/crypt"
	"github.com/rombintu/gopasswd.git/csvman"
	"github.com/rombintu/gopasswd.git/database"
	"github.com/rombintu/gopasswd.git/models"
)

func Check_login(res http.ResponseWriter, req *http.Request) string {
	gob.Register(sesKey(0))
	sesStatus, err := cookieStore.Get(req, cookieName)
	if err != nil {
		fmt.Fprintf(res, err.Error(), http.StatusBadRequest)
	}

	login, ok := sesStatus.Values[sesKeyLogin].(string)
	if !ok {
		login = "nil"
	}

	if login == "nil" {
		http.Redirect(res, req, "/sign", http.StatusSeeOther)
	}

	return login
}

func Index(res http.ResponseWriter, req *http.Request) {

	status := Check_login(res, req)

	log.Println(req.Method, "/")
	template, err := template.ParseFiles("templates/index.html", "templates/header.html")
	if err != nil {
		fmt.Fprintf(res, err.Error(), http.StatusBadRequest)
	}

	db := database.Get_db()

	var passwords []models.Passwords

	db.Find(&passwords, "User_login = ?", status)
	for i := 0; i < len(passwords); i++ {
		passwords[i].Pass = string(crypt.Decode_pass(passwords[i].Pass))
	}

	Data := &models.Index_page{Passwords: passwords, Status: status}

	template.ExecuteTemplate(res, "index", Data)
}

func Create(res http.ResponseWriter, req *http.Request) {

	status := Check_login(res, req)

	switch req.Method {
	case "GET":
		log.Println(req.Method, "/create")
		template, err := template.ParseFiles("templates/create.html", "templates/header.html")
		if err != nil {
			fmt.Fprintf(res, err.Error(), http.StatusBadRequest)
		}

		Data := &models.Other_page{Status: status}

		template.ExecuteTemplate(res, "create", Data)

	case "POST":
		log.Println(req.Method, "/create")
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
		db.Create(&models.Passwords{
			Service:    service,
			Url:        url,
			Email:      email,
			Pass:       enpass,
			User_login: status,
		})

		http.Redirect(res, req, "/", http.StatusSeeOther)
	}
}

func Import(res http.ResponseWriter, req *http.Request) {

	status := Check_login(res, req)

	switch req.Method {
	case "GET":
		log.Println(req.Method, "/import")
		template, err := template.ParseFiles("templates/import.html", "templates/header.html")
		if err != nil {
			fmt.Fprintf(res, err.Error(), http.StatusBadRequest)
		}

		Data := &models.Other_page{Status: status}

		template.ExecuteTemplate(res, "import", Data)
	case "POST":
		log.Println(req.Method, "/import")
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
			db.Create(&models.Passwords{
				Service:    data[i][0],
				Url:        data[i][1],
				Email:      data[i][2],
				Pass:       enpass,
				User_login: status,
			})
		}

		http.Redirect(res, req, "/", http.StatusSeeOther)
	}

}

// TODO
func Export(res http.ResponseWriter, req *http.Request) {
	status := Check_login(res, req)

	switch req.Method {
	case "GET":
		log.Println(req.Method, "/export")
		template, err := template.ParseFiles("templates/export.html", "templates/header.html")
		if err != nil {
			fmt.Fprintf(res, err.Error(), http.StatusBadRequest)
		}

		Data := &models.Other_page{Status: status}

		template.ExecuteTemplate(res, "export", Data)
	case "POST":
		db := database.Get_db()

		var chbox string = req.FormValue("chbox")
		var passwords []models.Passwords

		var depass string
		var Data_for_csv [][]string

		db.Find(&passwords, "User_login = ?", status)
		for i := 0; i < len(passwords); i++ {
			if chbox == "1" {
				depass = string(crypt.Decode_pass(passwords[i].Pass))
			} else {
				depass = passwords[i].Pass
			}

			Data_for_csv = append(Data_for_csv, []string{
				passwords[i].Service,
				passwords[i].Url,
				passwords[i].Email,
				depass,
			})
		}

		res.Header().Set("Content-Disposition", "attachment; filename=MyPasswords.csv")
		res.Header().Set("Content-Type", "application/octet-stream")
		res.WriteHeader(http.StatusOK)
		csvman.Export_csv(res, Data_for_csv)

		// http.Redirect(res, req, "/", http.StatusSeeOther)
	}
}

func Delete(res http.ResponseWriter, req *http.Request) {

	Check_login(res, req)

	log.Println(req.Method, "/delete")
	var id_pass string = req.FormValue("id_pass")

	db := database.Get_db()
	db.Delete(&models.Passwords{}, id_pass)

	http.Redirect(res, req, "/", http.StatusSeeOther)
}
