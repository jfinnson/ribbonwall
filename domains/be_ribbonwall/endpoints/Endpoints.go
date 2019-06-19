package endpoints

import (
	"github.com/jfinnson/ribbonwall/domains/be_ribbonwall/config"
	"github.com/jfinnson/ribbonwall/domains/be_ribbonwall/services"
	"github.com/jinzhu/gorm"
)

// Endpoints --
type Endpoints struct {
	Services services.RibbonwallServices
}

// NewEndpointService --
func NewEndpointService(
	dbClient *gorm.DB,
	config config.Config,
) *Endpoints {
	return &Endpoints{
		Services: *services.NewService(
			dbClient,
			config,
		),
	}
}
