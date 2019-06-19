package services

import (
	log "github.com/jfinnson/ribbonwall/common/logging"
	"github.com/jfinnson/ribbonwall/domains/be_ribbonwall/models"
)

func (services *RibbonwallServices) DeleteCompetitor(
	competitor models.Competitor,
) {
	db := services.DB
	deleteResult := db.Delete(competitor)

	log.Infof("Deleted: %s", deleteResult)
}
