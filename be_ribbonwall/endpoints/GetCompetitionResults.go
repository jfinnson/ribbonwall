package endpoints

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/ribbonwall/be_ribbonwall/models"
	log "github.com/ribbonwall/common/logging"
	"net/http"
)

func GetCompetitionResults(c *gin.Context) {
	s := c.Keys["services"].(*Endpoints)

	var competition_results []models.CompetitionResults
	s.Services.DB.Debug().Find(&competition_results)

	data, err := json.Marshal(competition_results)
	if err != nil {
		log.Errorf("error marshalling entity, %v", err)
		return
	}

	c.String(http.StatusOK, fmt.Sprintf("Database!\n Test results:\n%s", data))
}
