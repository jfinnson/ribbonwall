package config

import (
	"fmt"
	"os"
	"sync"

	"github.com/jinzhu/configor"
	"github.com/ribbonwall/be_ribbonwall/config"
)

var once sync.Once
var conf config.Config

// Get returns a singleton instance of the application configs
func Get() config.Config {
	// Make sure initialization is threadsafe
	once.Do(func() {
		// Load config
		conf = config.Config{}
		_ = configor.Load(&conf, fmt.Sprintf("%s/config.yaml", os.Getenv("SERVICE_CONFIG")))
	})

	return conf
}
