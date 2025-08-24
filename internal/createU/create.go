package createU

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kosyagut/userdata/internal/handler"
	"github.com/kosyagut/userdata/internal/storage"
)

func PostUser(userStorage *storage.UserStorage) gin.HandlerFunc {
	return func(c *gin.Context) {
		var newUser storage.User

		if err := c.BindJSON(&newUser); err != nil {
			c.JSON(http.StatusBadRequest, handler.Error{Error: "Bad Request"})
			return
		}

		newUser.ID = uuid.New()
		newUser.DateReg = time.Now()

		if err := userStorage.CreateUser(newUser); err != nil {
			c.JSON(http.StatusInternalServerError, handler.Error{Error: err.Error()})
			return
		}

		c.JSON(http.StatusCreated, newUser)
	}
}
