package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kosyagut/userdata/internal/readU"
)

func main() {
	router := gin.Default()
	router.GET("/users", getUsers)
	//router.POST("/user/:id", postIDUser)
	//router.PUT("user/:id", putUser)
	//router.DELETE("user/:id", deleteUser)

	err := router.Run("localhost:8080")
	if err != nil {
		return
	}
}
