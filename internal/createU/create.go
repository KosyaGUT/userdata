package createU

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kosyagut/userdata/internal/handler"
)

func PostUser(c *gin.Context) {
	var newUser handler.User

	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, handler.Error{Error: "Bad Request"})
	}

	newUser.ID = uuid.New()
	newUser.DateReg = time.Now()

	handler.Users = append(handler.Users, newUser)
	c.IndentedJSON(http.StatusCreated, newUser)
}
