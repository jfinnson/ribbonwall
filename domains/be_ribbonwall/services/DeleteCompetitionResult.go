package services

import (
	log "github.com/jfinnson/ribbonwall/common/logging"
	"github.com/jfinnson/ribbonwall/domains/be_ribbonwall/models"
)

func (services *RibbonwallServices) DeleteCompetitionResult(
	competitionResult models.CompetitionResults,
) {
	db := services.DB
	deleteResult := db.Delete(competitionResult)

	log.Infof("Deleted: %s", deleteResult)
}
