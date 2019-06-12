package auth

// Config --
type Config struct {
	Audience     string `yaml:"audience" env:"AUTH_AUDIENCE"`
	Issuer       string `yaml:"issuer" env:"AUTH_ISSUER"`
	Domain       string `yaml:"domain" env:"AUTH_DOMAIN"`
	ClientID     string `yaml:"clientID" env:"AUTH_CLIENT_ID"`
	ClientSecret string `env:"AUTH_CLIENT_SECRET"`
	RedirectURL  string `yaml:"redirectURL" env:"AUTH_REDIRECT_URL"`
	LoginURL     string `yaml:"loginURL" env:"AUTH_LOGIN_URL"`
	HomeURL      string `yaml:"homeURL" env:"AUTH_HOME_URL"`
}
