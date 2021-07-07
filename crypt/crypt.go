package crypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"os"
)

func Encode_pass(pass []byte) string {
	key := []byte(os.Getenv("KEY"))
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

func Decode_pass(enpass string) []byte {
	key := []byte(os.Getenv("KEY"))
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
		depass = []byte("Error key")
	}
	return depass
}
