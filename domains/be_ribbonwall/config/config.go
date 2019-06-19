package config

import (
	commonAuth "github.com/jfinnson/ribbonwall/common/auth"
	commonMySQL "github.com/jfinnson/ribbonwall/common/mysql"
)

// Config --
type Config struct {
	Env string

	// DB config
	DB commonMySQL.Config

	// Auth config
	Auth commonAuth.Config
}
