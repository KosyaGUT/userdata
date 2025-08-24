package deleteU

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kosyagut/userdata/internal/handler"
)

func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	strID, err := uuid.Parse(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, handler.Error{Error: "Bad Request"})
	}

	for i, user := range handler.Users {
		if user.ID == strID {
			handler.Users = append(handler.Users[:i], handler.Users[i+1:]...)
			c.IndentedJSON(http.StatusNoContent, handler.Users)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, handler.Error{Error: "User not found"})
}
