package services

import (
	"bufio"
	"encoding/csv"
	"errors"
	"fmt"
	"github.com/ribbonwall/domains/be_ribbonwall/models"
	"io"
	"mime/multipart"
	"strconv"
	"strings"
	"time"

	log "github.com/ribbonwall/common/logging"
)

func (services *RibbonwallServices) UploadCompetitionResults(file multipart.File) ([]competitionResultsRaw, error) {
	parsedCsv, err := parseCsv(file)
	if err != nil {
		log.Errorf("error parsing csv, %v", err)
		return nil, err
	}

	_, err = importCompetitionResults(services, parsedCsv)
	if err != nil {
		log.Errorf("error parsing csv, %v", err)
		return nil, err
	}

	return parsedCsv, nil
}

type competitionResultsRaw struct {
	CompetitorExternalID string    `json:"competitorExternalID"`
	CompetitorFirstName  string    `json:"competitorFirstName"`
	CompetitorLastName   string    `json:"competitorLastName"`
	HorseName            string    `json:"horseName"`
	TeamName             string    `json:"teamName"`
	Placing              int       `json:"placing"`
	Score                string    `json:"score"`
	OrganizationName     string    `json:"organizationName"`
	CompetitionName      string    `json:"competitionName"`
	DivisionName         string    `json:"divisionName"`
	ClassName            string    `json:"className"`
	CompetitionDate      time.Time `json:"competitionDate"`
}

func parseCsv(file multipart.File) ([]competitionResultsRaw, error) {
	reader := csv.NewReader(bufio.NewReader(file))
	var compResults []competitionResultsRaw
	rowCount := -1 // Ignore first row, assume its a header row.
	for {
		column, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Panicf("Error in csv parsing: %s", err)
		}

		rowCount++
		if rowCount == 0 {
			// Skip first row assuming its a header row
			continue
		}

		// Parse placing
		var Placing int
		if column[5] != "" {
			Placing, err = strconv.Atoi(column[5])
			if err != nil {
				log.Errorf("error parsing \"placing\", %v", err)
				return nil, err
			}
		} else {
			Placing = 0 // 0 is used as no-value atm.
		}

		// Parse competition date
		var CompetitionDate time.Time
		if column[11] != "" {
			CompetitionDate, err = time.Parse("2006-01-02", column[11])
			if err != nil {
				log.Errorf("error parsing \"competition date\", %v", err)
				return nil, err
			}
		} else {
			CompetitionDate = time.Now()
		}

		compResult := competitionResultsRaw{
			CompetitorExternalID: column[0],
			CompetitorFirstName:  column[1],
			CompetitorLastName:   column[2],
			HorseName:            column[3],
			TeamName:             column[4],
			Placing:              Placing,
			Score:                column[6],
			OrganizationName:     column[7],
			CompetitionName:      column[8],
			DivisionName:         column[9],
			ClassName:            column[10],
			CompetitionDate:      CompetitionDate,
		}

		// Check for mandatory fields. Eventually make this a convenience function.
		var missingFields []string
		if compResult.CompetitorExternalID == "" {
			missingFields = append(missingFields, "CompetitorExternalID")
		}
		if compResult.CompetitorFirstName == "" {
			missingFields = append(missingFields, "CompetitorFirstName")
		}
		if compResult.CompetitorLastName == "" {
			missingFields = append(missingFields, "CompetitorLastName")
		}
		if compResult.OrganizationName == "" {
			missingFields = append(missingFields, "OrganizationName")
		}
		if len(missingFields) > 0 {
			msg := fmt.Sprintf("Error, row %v: missing fields: %v", rowCount, strings.Join(missingFields, ", "))
			log.Errorf(msg)
			return nil, errors.New(msg)
		}

		compResults = append(compResults, compResult)
	}
	return compResults, nil
}

func importCompetitionResults(services *RibbonwallServices, resultsRaw []competitionResultsRaw) ([]*models.CompetitionResults, error) {
	db := services.DB

	// begin a transaction
	tx := db.Begin()

	var importedResults []*models.CompetitionResults
	for _, result := range resultsRaw {
		// Create competitor if one does not exist with the external ID
		competitor, err := services.CreateOrUpdateCompetitorByExternalID(result.CompetitorFirstName,
			result.CompetitorLastName, result.CompetitorExternalID, result.TeamName)

		// Check if not found. (awkward comparator, cleaner approach?)
		if err != nil {
			log.Errorf("error creating/updating new competitor entry, %v", err)
			return nil, err
		}

		// Create competition results
		// Warning this is not idempotent. This is because atm there is no good way to check if an entry is a duplicate.
		competitionResult, err := services.CreateCompetitionResult(
			competitor,
			result.OrganizationName,
			result.HorseName,
			result.CompetitionName,
			result.CompetitionDate,
			result.DivisionName,
			result.ClassName,
			result.Placing,
			result.Score,
		)
		if err != nil {
			log.Errorf("error creating new competitor entry, %v", err)
			return nil, err
		}

		importedResults = append(importedResults, competitionResult)
	}

	// All competitors and results have been created. Commit.
	tx.Commit()

	return importedResults, nil
}
