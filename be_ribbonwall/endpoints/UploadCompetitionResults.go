package endpoints

import (
	"bufio"
	"encoding/csv"
	"github.com/gin-gonic/gin"
	"io"
	"mime/multipart"
	"net/http"
	"time"

	log "github.com/ribbonwall/common/logging"
)

// Binding for validation
type CompetitionResults struct {
	OrganizationName string                `form:"organization" binding:"required"`
	ResultsFile      *multipart.FileHeader `form:"competition_results" binding:"required"`
}

func UploadCompetitionResults(c *gin.Context) {
	var competitionResults CompetitionResults
	if err := c.ShouldBind(&competitionResults); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Source
	file, _ := competitionResults.ResultsFile.Open()
	parsedCsv, err := parseCsv(file)
	if err != nil {
		log.Errorf("error parsing csv, %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, parsedCsv)

}

type CompetitionResultsRaw struct {
	CompetitorExternalID string    `json:"competitorExternalID"`
	CompetitorFirstName  string    `json:"competitorFirstName"`
	CompetitorLastName   string    `json:"competitorLastName"`
	HorseName            string    `json:"horseName"`
	TeamName             string    `json:"teamName"`
	Placing              string    `json:"placing"`
	Score                string    `json:"score"`
	OrganizationName     string    `json:"organizationName"`
	CompetitionName      string    `json:"competitionName"`
	DivisionName         string    `json:"divisionName"`
	ClassName            string    `json:"className"`
	CompetitionDate      time.Time `json:"competitionDate"`
}

func parseCsv(file multipart.File) ([]CompetitionResultsRaw, error) {
	reader := csv.NewReader(bufio.NewReader(file))
	var compResults []CompetitionResultsRaw
	rowCount := 0
	for {
		column, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Panicf("Error in csv parsing: %s", err)
		}

		rowCount++
		if rowCount == 1 {
			// Skip first row assuming its a header row
			continue
		}

		var CompetitionDate time.Time
		if column[11] != "" {
			CompetitionDate, err = time.Parse("2006-01-02", column[11])
			if err != nil {
				log.Errorf("error parsing competition date, %v", err)
				return nil, err
			}
		} else {
			CompetitionDate = time.Now()
		}

		compResults = append(compResults, CompetitionResultsRaw{
			CompetitorExternalID: column[0],
			CompetitorFirstName:  column[1],
			CompetitorLastName:   column[2],
			HorseName:            column[3],
			TeamName:             column[4],
			Placing:              column[5],
			Score:                column[6],
			OrganizationName:     column[7],
			CompetitionName:      column[8],
			DivisionName:         column[9],
			ClassName:            column[10],
			CompetitionDate:      CompetitionDate,
		})
	}
	return compResults, nil
}
