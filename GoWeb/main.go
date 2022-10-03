package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func welcomeHandler() func(context *gin.Context) {
	return func(context *gin.Context) {
		name := context.Param("name")
		result := fmt.Sprintf("Hello %s", name)
		context.JSON(http.StatusOK, gin.H{"message": result})
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
	err := router.Run()
	check(err)
}
