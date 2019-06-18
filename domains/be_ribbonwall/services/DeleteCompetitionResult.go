package services

import (
	log "github.com/ribbonwall/common/logging"
	"github.com/ribbonwall/domains/be_ribbonwall/models"
)

func (services *RibbonwallServices) DeleteCompetitionResult(
	competitionResult models.CompetitionResults,
) {
	db := services.DB
	deleteResult := db.Delete(competitionResult)

	log.Infof("Deleted: %s", deleteResult)
}
