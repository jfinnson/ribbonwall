package db

import (
	"sync"

	"github.com/jinzhu/gorm"
	"github.com/ribbonwall/be_ribbonwall/singletons/config"
	commonMySQL "github.com/ribbonwall/common/mysql"
)

var once sync.Once
var client *gorm.DB

// Get returns a singleton instance of the MySQL DB client
func Get() (*gorm.DB, error) {

	// Make sure initialization is thread safe
	var err error
	once.Do(func() {

		client, err = commonMySQL.Initialize(config.Get().DB)
		if err != nil {
			return
		}
	})
	if err != nil {
		return nil, err
	}

	return client, nil
}
