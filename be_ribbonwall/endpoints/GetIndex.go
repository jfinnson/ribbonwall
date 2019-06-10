package endpoints

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetIndex(c *gin.Context) {
	c.String(http.StatusOK, "Hello, World!!")
}
