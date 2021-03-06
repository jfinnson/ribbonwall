package endpoints

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jfinnson/ribbonwall/domains/be_ribbonwall/models"
	"net/http"
	"time"
)

// Binding for validation
type updateCompetitionResultParams struct {
	CompetitorUUID   string    `form:"competitor_uuid" binding:"required"`
	OrganizationName string    `form:"organization_name" binding:"required"`
	HorseName        string    `form:"horse_name" binding:"required"`
	CompetitionName  string    `form:"competition_name" binding:"required"`
	CompetitionDate  time.Time `form:"competition_date" binding:"required"`
	DivisionName     string    `form:"division_name" binding:"required"`
	ClassName        string    `form:"class_name" binding:"required"`
	Placing          int       `form:"placing" binding:"required"`
	Score            string    `form:"score" binding:"required"`
}

func UpdateCompetitionResult(c *gin.Context) {
	s := c.Keys["services"].(*Endpoints).Services

	var fields updateCompetitionResultParams
	if err := c.ShouldBind(&fields); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	uuid := c.Params.ByName("uuid")
	var competitionResult models.CompetitionResults
	if err := s.DB.Where("uuid = ?", uuid).First(&competitionResult).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("competition result not found (%s)", uuid)})
		return
	}

	// Create new competition result
	competitionResultUpdated, err := s.UpdateCompetitionResult(
		&competitionResult,
		fields.OrganizationName,
		fields.HorseName,
		fields.CompetitionName,
		fields.CompetitionDate,
		fields.DivisionName,
		fields.ClassName,
		fields.Placing,
		fields.Score,
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, competitionResultUpdated)
}
