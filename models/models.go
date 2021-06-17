package models

type Users struct {
	Id       uint
	Login    string
	Password string
}

type Passwords struct {
	Id      uint
	Title   string `json:"title"`
	content string `json:"content"`
}
