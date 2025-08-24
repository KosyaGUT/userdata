package updateU

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kosyagut/userdata/internal/handler"
	"github.com/kosyagut/userdata/internal/storage"
)

func UpdateIDUser(userStorage *storage.UserStorage) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		userID, err := uuid.Parse(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, handler.Error{Error: "Invalid ID format"})
			return
		}

		// Проверяем существование пользователя
		_, err = userStorage.GetUserByID(userID)
		if err != nil {
			c.JSON(http.StatusNotFound, handler.Error{Error: "User not found"})
			return
		}

		// Создаем карту для обновляемых полей
		updates := make(map[string]interface{})

		// Парсим JSON запроса в временную структуру
		var requestData map[string]interface{}
		if err := c.BindJSON(&requestData); err != nil {
			c.JSON(http.StatusBadRequest, handler.Error{Error: "Bad Request"})
			return
		}

		// Добавляем в updates только те поля, которые присутствуют в запросе
		if login, exists := requestData["login"]; exists {
			updates["login"] = login
		}
		if fcs, exists := requestData["fcs"]; exists {
			updates["fcs"] = fcs
		}
		if sex, exists := requestData["sex"]; exists {
			updates["sex"] = sex
		}
		if age, exists := requestData["age"]; exists {
			updates["age"] = age
		}
		if contacts, exists := requestData["contacts"]; exists {
			updates["contacts"] = contacts
		}
		if avatar, exists := requestData["avatar"]; exists {
			updates["avatar"] = avatar
		}
		if status, exists := requestData["status"]; exists {
			updates["status"] = status
		}

		// Выполняем частичное обновление
		if err := userStorage.PartialUpdateUser(userID, updates); err != nil {
			c.JSON(http.StatusInternalServerError, handler.Error{Error: err.Error()})
			return
		}

		// Получаем обновленного пользователя для ответа
		updatedUser, err := userStorage.GetUserByID(userID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, handler.Error{Error: err.Error()})
			return
		}

		c.JSON(http.StatusOK, updatedUser)
	}
}
