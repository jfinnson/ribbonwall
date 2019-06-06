package mysql

// Config --
type Config struct {
	Host           string `env:"DB_HOST"`
	Name           string `env:"DB_NAME"`
	Port           string `env:"DB_PORT"`
	User           string `env:"DB_USER"`
	Password       string `env:"DB_PASSWORD"`
	SSLMode        string `env:"DB_SSL_MODE"`
	ConnectionName string `yaml:"connectionName"`
	DebugLog       bool   `yaml:"debugLog"`
}
