package endpoints

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/ribbonwall/domains/be_ribbonwall/models"
	"net/http"
)

func DeleteCompetitionResult(c *gin.Context) {
	s := c.Keys["services"].(*Endpoints).Services

	uuid := c.Params.ByName("uuid")
	var competitionResult models.CompetitionResults
	if err := s.DB.Where("uuid = ?", uuid).First(&competitionResult).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Competition Result not found (%s)", uuid)})
		return
	}

	// Create competitor
	s.DeleteCompetitionResult(competitionResult)

	c.JSON(http.StatusOK, gin.H{"uuid #" + uuid: "deleted"})
}
