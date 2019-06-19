package endpoints

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jfinnson/ribbonwall/domains/be_ribbonwall/models"
	"net/http"
)

// Binding for validation
type updateCompetitorParams struct {
	FirstName  string `form:"first_name" binding:"required"`
	LastName   string `form:"last_name" binding:"required"`
	ExternalID string `form:"external_id" binding:"required"`
	TeamName   string `form:"team_name" binding:"required"`
}

func UpdateCompetitor(c *gin.Context) {
	s := c.Keys["services"].(*Endpoints).Services

	uuid := c.Params.ByName("uuid")
	var fields updateCompetitorParams
	if err := c.ShouldBind(&fields); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var competitor models.Competitor
	if err := s.DB.Where("uuid = ?", uuid).First(&competitor).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Competitor %s %s not found (%s)",
			fields.FirstName, fields.LastName, uuid)})
		return
	}

	// Create competitor
	updatedCompetitor, err := s.UpdateCompetitor(&competitor, fields.FirstName, fields.LastName,
		fields.ExternalID, fields.TeamName)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedCompetitor)
}
