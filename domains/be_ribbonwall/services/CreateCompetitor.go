package services

import (
	"errors"
	"fmt"
	log "github.com/jfinnson/ribbonwall/common/logging"
	"github.com/jfinnson/ribbonwall/domains/be_ribbonwall/models"
)

func (services *RibbonwallServices) CreateCompetitor(
	firstName string,
	lastName string,
	externalId string,
	teamName string,
) (*models.Competitor, error) {
	db := services.DB
	tx := db.Begin()

	var competitor models.Competitor
	db.Where("external_id = ?", externalId).First(&competitor)

	if competitor != (models.Competitor{}) {
		msg := fmt.Sprintf("Competitor already exists with external id: %s", externalId)
		log.Errorf(msg)
		return nil, errors.New(msg)
	}

	newCompetitor := &models.Competitor{
		FirstName:  firstName,
		LastName:   lastName,
		ExternalId: externalId,
		TeamName:   teamName,
	}

	if err := tx.Create(newCompetitor).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return newCompetitor, nil
}
