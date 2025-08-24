package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/kosyagut/userdata/internal/createU"
	"github.com/kosyagut/userdata/internal/deleteU"
	"github.com/kosyagut/userdata/internal/readU"
	"github.com/kosyagut/userdata/internal/storage"
	"github.com/kosyagut/userdata/internal/updateU"
)

func main() {
	cfg := storage.Config{
		Host:     getEnv("DB_HOST", "localhost"),
		Port:     getEnv("DB_PORT", "5432"),
		Username: getEnv("DB_USER", "postgres"),
		Password: getEnv("DB_PASSWORD", ""),
		DBName:   getEnv("DB_NAME", "userdata"),
		SSLMode:  getEnv("DB_SSLMODE", "disable"),
	}

	db := storage.NewPostgresDB(cfg)
	defer func(db *sqlx.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(db)

	if err := storage.InitSchema(db); err != nil {
		log.Fatalf("Ошибка миграции: %v", err)
	}

	log.Println("База данных готова таблица создана!")
	userStorage := storage.NewUserStorage(db)

	router := gin.Default()

	router.GET("/users", readU.GetUsers(userStorage))
	router.GET("/users/:id", readU.GetIDUser(userStorage))
	router.POST("/users", createU.PostUser(userStorage))
	router.PUT("/users/:id", updateU.UpdateIDUser(userStorage))
	router.DELETE("/users/:id", deleteU.DeleteUser(userStorage))

	err := router.Run(":8080") // Слушаем на всех интерфейсах
	if err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}

}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
