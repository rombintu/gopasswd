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

func db_manage() {

	// Create
	// db.Create(&models.Users{1, "Login", "Passw"})

	// Read
	// var users models.Users
	// db.First(&users, "Id = ?", 1) // find product with integer primary key
	// db.First(&product, "code = ?", "D42") // find product with code D42
	// fmt.Println(users)
	// // Update - update product's price to 200
	// db.Model(&product).Update("Price", 200)
	// db.Model(&usqers).Update("Login", "Hello")
	// // Update - update multiple fields
	// db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
	// db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// // Delete - delete product
	// db.Delete(&product, 1)
}

func main() {
	database.Init()
	fmt.Println("Go to: http://localhost:8080")
	// db_manage()
	listen()
}
