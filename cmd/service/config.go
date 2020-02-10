package main

import (
	"github.com/breiting/bookless/pkg/database"
)

// Config to configure the app
type Config struct {
	DataBase database.Config
}

// GetDefaultConfig gets the local configuration
func GetDefaultConfig() *Config {
	return &Config{
		DataBase: database.Config{
			Hostname: "neo",
			Port:     5432,
			Dialect:  "postgres",
			Username: "bookless",
			Password: "bookless",
			Name:     "bookless",
			Charset:  "utf8",
			Migrate:  true,
		},
	}
}
