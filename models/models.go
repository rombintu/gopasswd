package models

type Users struct {
	Id       uint
	Login    string
	Password string
}

type Passwords struct {
	Id      uint
	Service string `json:"service"`
	Email   string `json:"email"`
	Pass    string `json:"pass"`
}
