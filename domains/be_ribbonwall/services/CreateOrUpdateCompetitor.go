package services

import (
	"github.com/ribbonwall/domains/be_ribbonwall/models"
)

func (services *RibbonwallServices) CreateOrUpdateCompetitorByExternalID(
	firstName string,
	lastName string,
	externalId string,
	teamName string,
) (*models.Competitor, error) {
	db := services.DB

	var competitor models.Competitor
	if err := db.Where("external_id = ?", externalId).
		Assign(models.Competitor{
			FirstName:  firstName,
			LastName:   lastName,
			ExternalId: externalId,
			TeamName:   teamName}).
		FirstOrCreate(&competitor).Error; err != nil {
		return nil, err
	}

	return &competitor, nil
}
