package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rombintu/gopasswd.git/database"
	"github.com/rombintu/gopasswd.git/routes"
)

func listen() {
	router := mux.NewRouter()

	router.HandleFunc("/", routes.Index)
	router.HandleFunc("/sign", routes.Sign)
	router.HandleFunc("/reg", routes.Reg)
	router.HandleFunc("/logout", routes.Logout)
	router.HandleFunc("/create", routes.Create)
	router.HandleFunc("/import", routes.Import)
	router.HandleFunc("/delete", routes.Delete)

	router.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.ListenAndServe(":8080", router)
}

func main() {
	database.Init()
	fmt.Println("Go to: http://localhost:8080")
	listen()
}
