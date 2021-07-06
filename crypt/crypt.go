package crypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
)

func encode_pass(pass []byte, key []byte) string {
	c, err := aes.NewCipher(key)

	if err != nil {
		fmt.Println(err)
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		fmt.Println(err)
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		fmt.Println(err)
	}

	enpass := gcm.Seal(nonce, nonce, pass, nil)
	return string(enpass)

}

func decode_pass(enpass string, key []byte) []byte {

	c, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(err)
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		fmt.Println(err)
	}

	nonceSize := gcm.NonceSize()
	if len(enpass) < nonceSize {
		fmt.Println(err)
	}

	nonce, enpass := enpass[:nonceSize], enpass[nonceSize:]
	depass, err := gcm.Open(nil, []byte(nonce), []byte(enpass), nil)
	if err != nil {
		fmt.Println(err)
	}
	return depass
}

// func db_manage() {

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
// }
