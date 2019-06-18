package services

import (
	"github.com/ribbonwall/domains/be_ribbonwall/models"
	"time"
)

func (services *RibbonwallServices) UpdateCompetitionResult(
	competitionResult *models.CompetitionResults,
	organizationName string,
	horseName string,
	competitionName string,
	competitionDate time.Time,
	divisionName string,
	className string,
	placing int,
	score string,
) (*models.CompetitionResults, error) {
	db := services.DB
	tx := db.Begin()

	competitionResult.OrganizationName = organizationName
	competitionResult.HorseName = horseName
	competitionResult.CompetitionName = competitionName
	competitionResult.CompetitionDate = competitionDate
	competitionResult.DivisionName = divisionName
	competitionResult.ClassName = className
	competitionResult.Placing = placing
	competitionResult.Score = score

	if err := tx.Save(&competitionResult).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return competitionResult, nil
}
