package readU

import (
	"github.com/kosyagut/userdata/internal/handler"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)
}
