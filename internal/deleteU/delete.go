package deleteU

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kosyagut/userdata/internal/handler"
	"github.com/kosyagut/userdata/internal/storage"
)

func DeleteUser(userStorage *storage.UserStorage) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		userID, err := uuid.Parse(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, handler.Error{Error: "Bad Request"})
			return
		}

		if err := userStorage.DeleteUser(userID); err != nil {
			c.JSON(http.StatusInternalServerError, handler.Error{Error: err.Error()})
			return
		}

		c.Status(http.StatusNoContent)
	}
}
