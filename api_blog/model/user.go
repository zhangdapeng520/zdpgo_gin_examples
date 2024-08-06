package model

type User struct {
	Id       int     `json:"id"`
	Username string  `json:"username"`
	Password string  `json:"password"`
	Money    float64 `json:"money"`
}