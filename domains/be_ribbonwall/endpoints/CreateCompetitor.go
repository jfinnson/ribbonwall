package endpoints

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

// Binding for validation
type createCompetitorParams struct {
	FirstName  string `form:"first_name" binding:"required"`
	LastName   string `form:"last_name" binding:"required"`
	ExternalID string `form:"external_id" binding:"required"`
	TeamName   string `form:"team_name" binding:"required"`
}

func CreateCompetitor(c *gin.Context) {
	s := c.Keys["services"].(*Endpoints).Services

	var fields createCompetitorParams
	if err := c.ShouldBind(&fields); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create competitor
	newCompetitor, err := s.CreateCompetitor(fields.FirstName, fields.LastName,
		fields.ExternalID, fields.TeamName)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, newCompetitor)
}
