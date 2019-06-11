package mysql

import (
	"fmt"
	"github.com/jinzhu/gorm"
	log "github.com/ribbonwall/common/logging"
)

// Initialize --
func Initialize(config Config) (*gorm.DB, error) {
	dbStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.Name,
	)
	log.Infof("dbStr %s", dbStr)
	// Open connection
	dbClient, err := gorm.Open(
		"mysql",
		dbStr,
	)
	dbClient.LogMode(true) // For debugging
	if err != nil {
		log.Error(err)
		return nil, err
	}

	// Connection Pool Settings
	dbClient.DB().SetMaxIdleConns(6)
	dbClient.DB().SetMaxOpenConns(12)

	// Enable Logger
	dbClient.LogMode(config.DebugLog)
	dbClient.SetLogger(log.GetLogger())

	log.Infof("Connected to MySQL at %s/%s", config.Host, config.Name)

	return dbClient, nil
}
