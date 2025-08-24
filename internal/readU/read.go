package readU

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kosyagut/userdata/internal/handler"
	"github.com/kosyagut/userdata/internal/storage"
)

func GetUsers(userStorage *storage.UserStorage) gin.HandlerFunc {
	return func(c *gin.Context) {
		users, err := userStorage.GetUsers()
		if err != nil {
			c.JSON(http.StatusInternalServerError, handler.Error{Error: err.Error()})
			return
		}
		c.JSON(http.StatusOK, users)
	}
}

func GetIDUser(userStorage *storage.UserStorage) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		userID, err := uuid.Parse(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, handler.Error{Error: "invalid id format"})
			return
		}

		user, err := userStorage.GetUserByID(userID)
		if err != nil {
			c.JSON(http.StatusNotFound, handler.Error{Error: "User not found"})
			return
		}

		c.JSON(http.StatusOK, user)
	}
}
