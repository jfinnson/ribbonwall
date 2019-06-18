package services

import (
	"github.com/ribbonwall/domains/be_ribbonwall/models"
)

func (services *RibbonwallServices) UpdateCompetitor(
	competitor models.Competitor,
	firstName string,
	lastName string,
	externalId string,
	teamName string,
) (*models.Competitor, error) {
	db := services.DB
	tx := db.Begin()

	competitor.FirstName = firstName
	competitor.LastName = lastName
	competitor.ExternalId = externalId
	competitor.TeamName = teamName

	if err := tx.Save(&competitor).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return &competitor, nil
}
