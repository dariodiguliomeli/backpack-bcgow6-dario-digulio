package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

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
	router.GET("/GetAll", getUsersHandler())
	router.GET("/users/:id", getUserById())
	err := router.Run()
	check(err)
}
