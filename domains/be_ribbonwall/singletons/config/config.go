package config

import (
	"fmt"
	"github.com/jfinnson/ribbonwall/domains/be_ribbonwall/config"
	"os"
	"sync"

	"github.com/jinzhu/configor"
)

var once sync.Once
var conf config.Config

// Get returns a singleton instance of the application configs
func Get() config.Config {
	// Make sure initialization is thread safe
	once.Do(func() {
		// Load config
		conf = config.Config{}
		_ = configor.Load(&conf, fmt.Sprintf("domains/be_ribbonwall/config/config.%s.yaml", os.Getenv("SERVICE_CONFIG")))
	})

	return conf
}
