package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/kosyagut/userdata/internal/createU"
	"github.com/kosyagut/userdata/internal/deleteU"
	"github.com/kosyagut/userdata/internal/readU"
	"github.com/kosyagut/userdata/internal/storage"
	"github.com/kosyagut/userdata/internal/updateU"
)

func main() {
	cfg := storage.Config{
		Host:     "localhost",
		Port:     "5432",
		Username: "postgres",
		Password: "6189",
		DBName:   "userdata",
		SSLMode:  "disable",
	}

	db := storage.NewPostgresDB(cfg)
	defer db.Close()

	if err := storage.InitSchema(db); err != nil {
		log.Fatalf("Ошибка миграции: %v", err)
	}

	log.Println("База данных готова таблица создана!")
	// Инициализируем хранилище
	userStorage := storage.NewUserStorage(db)

	router := gin.Default()
	// Передаем хранилище в обработчики
	router.GET("/users", readU.GetUsers(userStorage))
	router.GET("/users/:id", readU.GetIDUser(userStorage))
	router.POST("/users", createU.PostUser(userStorage))
	router.PUT("/users/:id", updateU.UpdateIDUser(userStorage))
	router.DELETE("/users/:id", deleteU.DeleteUser(userStorage))

	err := router.Run("localhost:8080")
	if err != nil {
		return
	}
}
