package endpoints

import (
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"net/http"
	//log "github.com/jfinnson/ribbonwall/common/logging"
)

// Binding for validation
type competitionResultsParams struct {
	OrganizationName string                `form:"organization" binding:"required"`
	ResultsFile      *multipart.FileHeader `form:"competition_results" binding:"required"`
}

func UploadCompetitionResults(c *gin.Context) {
	var competitionResultsParams competitionResultsParams
	if err := c.ShouldBind(&competitionResultsParams); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Source
	file, _ := competitionResultsParams.ResultsFile.Open()
	services := c.Keys["services"].(*Endpoints).Services
	competitionResults, err := services.UploadCompetitionResults(file)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, competitionResults)

}
