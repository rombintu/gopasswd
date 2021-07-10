package routes

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/gorilla/sessions"
	"github.com/rombintu/gopasswd.git/database"
	"github.com/rombintu/gopasswd.git/models"
	"gorm.io/gorm"
)

var cookieStore = sessions.NewCookieStore([]byte(os.Getenv("SECRET")))

const cookieName = "LocalCookiePasswds"

type sesKey int

const (
	sesKeyLogin sesKey = iota
)

func Sign(res http.ResponseWriter, req *http.Request) {
	gob.Register(sesKey(0))
	switch req.Method {
	case "GET":
		log.Println(req.Method, "/sign")
		template, err := template.ParseFiles("templates/sign.html", "templates/header.html")
		if err != nil {
			fmt.Fprintf(res, err.Error(), http.StatusBadRequest)
		}
		template.ExecuteTemplate(res, "sign", nil)

	case "POST":
		log.Println(req.Method, "/sign")
		var username string = req.FormValue("username")
		var password string = req.FormValue("password")

		if username == "" || password == "" || username == "nil" {
			http.Redirect(res, req, "/reg", http.StatusSeeOther)
			return
		}

		var user []models.Users

		h := sha256.New()
		h.Write([]byte(username))

		db := database.Get_db()

		err := db.First(&user, "Login = ?", username).Error
		if err == gorm.ErrRecordNotFound {
			log.Println("User not exists")
			http.Redirect(res, req, "/sign", http.StatusSeeOther)
			return
		}

		compare := bytes.Compare([]byte(user[0].Password), h.Sum([]byte(password)))

		if compare == 0 {
			sesStatus, err := cookieStore.Get(req, cookieName)
			if err != nil {
				http.Error(res, err.Error(), http.StatusBadRequest)
				return
			}

			sesStatus.Values[sesKeyLogin] = username
			err = cookieStore.Save(req, res, sesStatus)
			if err != nil {
				http.Error(res, err.Error(), http.StatusBadRequest)
				return
			}

			http.Redirect(res, req, "/", http.StatusSeeOther)
		} else {
			log.Println("Error")
			http.Redirect(res, req, "/sign", http.StatusSeeOther)
		}
	}

}

func Reg(res http.ResponseWriter, req *http.Request) {
	gob.Register(sesKey(0))
	switch req.Method {
	case "GET":
		log.Println(req.Method, "/reg")
		template, err := template.ParseFiles("templates/reg.html", "templates/header.html")
		if err != nil {
			fmt.Fprintf(res, err.Error(), http.StatusBadRequest)
		}
		template.ExecuteTemplate(res, "reg", nil)

	case "POST":
		log.Println(req.Method, "/reg")

		var username string = req.FormValue("username")
		var password string = req.FormValue("password")

		if username == "" || password == "" || username == "nil" {
			http.Redirect(res, req, "/reg", http.StatusSeeOther)
			return
		}

		var user models.Users

		h := sha256.New()
		h.Write([]byte(username))

		db := database.Get_db()

		err := db.Where("Login = ?", username).First(&user).Error

		if err != gorm.ErrRecordNotFound {
			log.Println("User exists")
			http.Redirect(res, req, "/sign", http.StatusSeeOther)
		} else {
			db.Create(&models.Users{Login: username, Password: string(h.Sum([]byte(password)))})

			sesStatus, err := cookieStore.Get(req, cookieName)
			if err != nil {
				http.Error(res, err.Error(), http.StatusBadRequest)
				return
			}

			sesStatus.Values[sesKeyLogin] = username
			err = cookieStore.Save(req, res, sesStatus)
			if err != nil {
				http.Error(res, err.Error(), http.StatusBadRequest)
				return
			}

			http.Redirect(res, req, "/", http.StatusSeeOther)

		}

	}

}

func Logout(res http.ResponseWriter, req *http.Request) {
	log.Println(req.Method, "/logout")
	sesStatus, err := cookieStore.Get(req, cookieName)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	sesStatus.Values[sesKeyLogin] = "nil"
	err = cookieStore.Save(req, res, sesStatus)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	http.Redirect(res, req, "/sign", http.StatusSeeOther)
}
