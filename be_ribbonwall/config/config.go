package config

import (
	commonMySQL "github.com/ribbonwall/common/mysql"
)

// Config --
type Config struct {
	Env string

	DB commonMySQL.Config
}
