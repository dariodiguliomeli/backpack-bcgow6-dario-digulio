package main

type User struct {
	Id      string  `json:"id"`
	Name    string  `json:"name"`
	Surname string  `json:"surname"`
	Email   string  `json:"email"`
	Age     int     `json:"age"`
	Height  float64 `json:"height"`
	Active  bool    `json:"active"`
	Date    string  `json:"date"`
}
