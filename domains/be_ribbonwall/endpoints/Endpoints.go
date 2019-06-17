package endpoints

import (
	"github.com/jinzhu/gorm"
	"github.com/ribbonwall/domains/be_ribbonwall/config"
	"github.com/ribbonwall/domains/be_ribbonwall/services"
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
