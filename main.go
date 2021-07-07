package main

import (
	"fmt"
	"net/http"

	"github.com/rombintu/gopassic.git/database"
	"github.com/rombintu/gopassic.git/routes"
)

func listen() {
	http.HandleFunc("/", routes.Index)
	http.HandleFunc("/create", routes.Create)
	http.HandleFunc("/create_push", routes.Create_push)
	http.HandleFunc("/import", routes.Import_passwords)
	http.HandleFunc("/import_push", routes.Import_passwords_push)

	http.ListenAndServe(":8080", nil)
}

func main() {
	database.Init()
	fmt.Println("Go to: http://localhost:8080")
	listen()
}
