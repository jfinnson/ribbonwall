package config

import (
	commonAuth "github.com/ribbonwall/common/auth"
	commonMySQL "github.com/ribbonwall/common/mysql"
)

// Config --
type Config struct {
	Env string

	// DB config
	DB commonMySQL.Config

	// Auth config
	Auth commonAuth.Config
}
