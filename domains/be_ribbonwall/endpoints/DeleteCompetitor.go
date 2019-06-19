package endpoints

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jfinnson/ribbonwall/domains/be_ribbonwall/models"
	"net/http"
)

func DeleteCompetitor(c *gin.Context) {
	s := c.Keys["services"].(*Endpoints).Services

	uuid := c.Params.ByName("uuid")
	var competitor models.Competitor
	if err := s.DB.Where("uuid = ?", uuid).First(&competitor).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Competitor not found (%s)", uuid)})
		return
	}

	// Create competitor
	s.DeleteCompetitor(competitor)

	c.JSON(http.StatusOK, gin.H{"uuid #" + uuid: "deleted"})
}
