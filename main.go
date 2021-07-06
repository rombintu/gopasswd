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
	http.HandleFunc("/push_create", routes.Push_create)
	http.HandleFunc("/login", routes.Login)
	http.ListenAndServe(":8080", nil)
}

func main() {
	database.Init()
	fmt.Println("Go to: http://localhost:8080")
	listen()
}
