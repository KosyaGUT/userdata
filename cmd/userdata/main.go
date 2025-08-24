package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kosyagut/userdata/internal/createU"
	"github.com/kosyagut/userdata/internal/deleteU"
	"github.com/kosyagut/userdata/internal/readU"
	"github.com/kosyagut/userdata/internal/updateU"
)

func main() {
	router := gin.Default()
	router.GET("/users", readU.GetUsers)
	router.GET("/users/:id", readU.GetIDUser)

	router.POST("/users", createU.PostUser)

	router.PUT("users/:id", updateU.UpdateIDUser)

	router.DELETE("users/:id", deleteU.DeleteUser)

	err := router.Run("localhost:8080")
	if err != nil {
		return
	}
}
