package models

type Users struct {
	Id       uint
	Login    string `json:"login"`
	Password string `json:"pass"`
}

type Passwords struct {
	Id      uint
	Service string `json:"service"`
	Url     string `json:"url"`
	Email   string `json:"email"`
	Pass    string `json:"pass"`
	User_id uint
}
