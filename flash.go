package main

import (
	"encoding/base64"
	"net/http"
	"time"
)

func SetFlash(res http.ResponseWriter, name string, value []byte) {
	c := &http.Cookie{Name: name, Value: encode(value)}
	http.SetCookie(res, c)
}

func GetFlash(res http.ResponseWriter, req *http.Request, name string) ([]byte, error) {
	c, err := req.Cookie(name)
	if err != nil {
		switch err {
		case http.ErrNoCookie:
			return nil, nil
		default:
			return nil, err
		}
	}
	value, err := decode(c.Value)
	if err != nil {
		return nil, err
	}
	dc := &http.Cookie{Name: name, MaxAge: -1, Expires: time.Unix(1, 0)}
	http.SetCookie(res, dc)
	return value, nil
}

func encode(src []byte) string {
	return base64.URLEncoding.EncodeToString(src)
}

func decode(src string) ([]byte, error) {
	return base64.URLEncoding.DecodeString(src)
}
