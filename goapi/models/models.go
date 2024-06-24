package models

type User struct {
	UserID   int    `json:"userid"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Instance struct {
	UserID   int    `json:"userid"`
	Password string `json:"password"`
	Port     int    `json:"port"`
}

