package endpoints

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jfinnson/ribbonwall/domains/be_ribbonwall/models"
	"net/http"
)

func GetCompetitors(c *gin.Context) {
	s := c.Keys["services"].(*Endpoints)

	var competitors []models.Competitor
	if err := s.Services.DB.
		Find(&competitors).
		Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, competitors)
}
