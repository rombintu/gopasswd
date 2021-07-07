package main

import (
	"fmt"
	"net/http"

	"github.com/rombintu/gopasswd.git/database"
	"github.com/rombintu/gopasswd.git/routes"
)

func listen() {
	http.HandleFunc("/", routes.Index)
	http.HandleFunc("/create", routes.Create)
	http.HandleFunc("/import", routes.Import)
	http.HandleFunc("/delete", routes.Delete)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.ListenAndServe(":8080", nil)
}

func main() {
	database.Init()
	fmt.Println("Go to: http://localhost:8080")
	listen()
}
