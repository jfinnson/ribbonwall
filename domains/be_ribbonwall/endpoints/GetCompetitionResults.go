package endpoints

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jfinnson/ribbonwall/domains/be_ribbonwall/models"
	"net/http"
)

func GetCompetitionResults(c *gin.Context) {
	s := c.Keys["services"].(*Endpoints)

	// This triggers two queries, not too bad but to use a join i would have to learn more about gorm or
	// use db.Join and write a custom mapper.
	var competitionResults []models.CompetitionResults
	if err := s.Services.DB.
		Preload("Competitor").
		Find(&competitionResults).
		Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": competitionResults})
}
