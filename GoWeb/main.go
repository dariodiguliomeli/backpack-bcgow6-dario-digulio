package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

var users []User

func welcomeHandler() func(context *gin.Context) {
	return func(context *gin.Context) {
		name := context.Param("name")
		result := fmt.Sprintf("Hello %s", name)
		context.JSON(http.StatusOK, gin.H{"message": result})
	}
}

func getUsersHandler() func(context *gin.Context) {
	return func(context *gin.Context) {
		file, err := os.ReadFile("./GoWeb/users.json")
		check(err)
		var users User
		err = json.Unmarshal(file, &users)
		check(err)
		context.JSON(http.StatusOK, gin.H{"users": users})
	}
}

func createUserHandler() func(context *gin.Context) {
	return func(context *gin.Context) {
		var req CreateUserRequest
		if err := context.ShouldBindJSON(&req); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		nextId := len(users) + 1
		newUser := User{string(rune(nextId)), req.Name, req.Surname, req.Email, req.Age, req.Height, req.Active, req.Date}
		users = append(users, newUser)
		context.JSON(http.StatusCreated, gin.H{"newUserId": nextId})
	}
}

func getUserById() func(context *gin.Context) {
	return func(context *gin.Context) {
		id := context.Param("id")
		file, err := os.ReadFile("./users.json")
		check(err)
		var users Users
		err = json.Unmarshal(file, &users)
		check(err)
		for _, user := range users.Users {
			check(err)
			if user.Id == id {
				context.JSON(http.StatusOK, gin.H{"user": user})
				return
			}
		}
		context.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	router := gin.Default()
	router.GET("/welcome/:name", welcomeHandler())
	users := router.Group("/users")
	{
		users.GET("/GetAll", getUsersHandler())
		users.GET("/:id", getUserById())
		users.POST("/create", createUserHandler())
	}
	err := router.Run()
	check(err)
}
