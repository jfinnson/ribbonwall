package endpoints

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type fileUpload struct {
	//OrganizationName string                `form:"organization" binding:"required"`
	Src string `json:"src" binding:"required"`
}

// Binding for validation
type competitionResultsParams struct {
	//OrganizationName string                `form:"organization" binding:"required"`
	ResultsFile fileUpload `json:"competition_results" binding:"required"`
}

func UploadCompetitionResults(c *gin.Context) {
	var competitionResultsParams competitionResultsParams
	if err := c.ShouldBind(&competitionResultsParams); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//file, err := dataurl.Decode(competitionResultsParams.ResultsFile.Src)
	//if err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	//}
	//
	//// Source
	////file, _ := competitionResultsParams.ResultsFile.Open()
	////services := c.Keys["services"].(*Endpoints).Services
	////competitionResults, err := services.UploadCompetitionResults(file)
	////if err != nil {
	////	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	////	return
	////}
	//
	//c.JSON(http.StatusOK, file)

}
