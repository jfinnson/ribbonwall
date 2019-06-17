package services

import "github.com/ribbonwall/domains/be_ribbonwall/models"

func (services *RibbonwallServices) CreateCompetitor(
	firstName string,
	lastName string,
	externalId string,
	teamName string,
) (*models.Competitor, error) {
	db := services.DB
	tx := db.Begin()

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
