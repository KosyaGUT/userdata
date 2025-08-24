package readU

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/kosyagut/userdata/internal/handler"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, handler.Users)
}

func GetIDUser(c *gin.Context) {
	id := c.Param("id")

	strID, err := uuid.Parse(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid id format"})
		return
	}

	for _, user := range handler.Users {
		if strID == user.ID {
			c.IndentedJSON(http.StatusOK, user)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, handler.Error{Error: "Not Found"})
}
