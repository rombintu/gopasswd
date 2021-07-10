package models

type Users struct {
	Id       uint
	Login    string `json:"login"`
	Password string `json:"pass"`
}

type Passwords struct {
	Id         uint
	Service    string `json:"service"`
	Url        string `json:"url"`
	Email      string `json:"email"`
	Pass       string `json:"pass"`
	User_login string `json:"user_login"`
}

type Index_page struct {
	Passwords []Passwords
	Status    string
}

type Other_page struct {
	Status string
}

type Export_data struct {
	Service []string
	Url     []string
	Email   []string
	Pass    []string
}
