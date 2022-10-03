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

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	router := gin.Default()
	router.GET("/welcome/:name", welcomeHandler())
	router.GET("/GetAll", getUsersHandler())
	err := router.Run()
	check(err)
}
