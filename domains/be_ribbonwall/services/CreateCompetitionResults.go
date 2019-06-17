package services

import (
	"github.com/ribbonwall/domains/be_ribbonwall/models"
	"time"
)

func (services *RibbonwallServices) CreateCompetitionResults(
	competitor *models.Competitor,
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

	newCompetitionResult := &models.CompetitionResults{
		OrganizationName: organizationName,

		//Competitor:           	*competitor,
		CompetitorUUID:       competitor.UUID,
		CompetitorExternalId: competitor.ExternalId,
		HorseName:            horseName,

		CompetitionName: competitionName,
		CompetitionDate: competitionDate,
		DivisionName:    divisionName,
		ClassName:       className,

		Placing: placing,
		Score:   score,
	}

	if err := tx.Create(newCompetitionResult).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return newCompetitionResult, nil
}
