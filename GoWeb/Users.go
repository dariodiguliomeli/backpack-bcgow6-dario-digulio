package main

type Users struct {
	Users []User `json:"users"`
}

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

type CreateUserRequest struct {
	Name    string  `json:"name" binding:"required"`
	Surname string  `json:"surname" binding:"required"`
	Email   string  `json:"email" binding:"required"`
	Age     int     `json:"age" binding:"required"`
	Height  float64 `json:"height" binding:"required"`
	Active  bool    `json:"active" binding:"required"`
	Date    string  `json:"date" binding:"required"`
}
