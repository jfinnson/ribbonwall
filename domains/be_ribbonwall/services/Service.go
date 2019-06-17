package services

import (
	"github.com/jinzhu/gorm"
	"github.com/ribbonwall/domains/be_ribbonwall/config"
	"sync"
)

// RibbonwallServices implements the Service interface.
// This struct is responsible for establishing and maintaining connections with others services like the database
type RibbonwallServices struct {
	DB     *gorm.DB
	Config config.Config
}

var ribbonwallServices *RibbonwallServices
var once sync.Once

// NewService --
func NewService(
	dbClient *gorm.DB,
	config config.Config,
) *RibbonwallServices {

	once.Do(func() {
		ribbonwallServices = &RibbonwallServices{
			DB:     dbClient,
			Config: config,
		}
	})
	return ribbonwallServices
}
