package crypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"log"
	"os"
)

func Encode_pass(pass []byte) string {
	key := []byte(os.Getenv("KEY"))
	if string(key) == "" || len(key) != 32 {
		log.Fatal("EXAMPLE:\n    export KEY=passphrasewhichneedstobe32bytes1")
	}
	c, err := aes.NewCipher(key)

	if err != nil {
		log.Panic(err)
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		log.Panic(err)
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		log.Panic(err)
	}

	enpass := gcm.Seal(nonce, nonce, pass, nil)
	return string(enpass)

}

func Decode_pass(enpass string) []byte {
	key := []byte(os.Getenv("KEY"))
	if string(key) == "" || len(key) != 32 {
		log.Fatal("EXAMPLE:\n    export KEY=passphrasewhichneedstobe32bytes1")
	}

	c, err := aes.NewCipher(key)
	if err != nil {
		log.Panic(err)
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		log.Panic(err)
	}

	nonceSize := gcm.NonceSize()
	if len(enpass) < nonceSize {
		log.Panic(err)
	}

	nonce, enpass := enpass[:nonceSize], enpass[nonceSize:]
	depass, err := gcm.Open(nil, []byte(nonce), []byte(enpass), nil)
	if err != nil {
		log.Panic(err)
	}
	return depass
}
