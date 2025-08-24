package updateU

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kosyagut/userdata/internal/handler"
)

func UpdateIDUser(c *gin.Context) {
	id := c.Param("id")

	strID, err := uuid.Parse(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, handler.Error{Error: "Bad Request"})
		return
	}

	for i, user := range handler.Users {
		if strID == user.ID {
			err := c.BindJSON(&user)
			if err != nil {
				return
			}

			newParam := handler.Users[i+1:]
			handler.Users = append(handler.Users[:i], user)
			handler.Users = append(handler.Users, newParam...)
			c.IndentedJSON(http.StatusCreated, user)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, handler.Error{Error: "Not Found"})
}
