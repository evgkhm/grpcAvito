package config

import (
	"github.com/spf13/viper"
)

type PostgresConfig struct {
	Host    string
	Port    string
	User    string
	Pass    string
	Name    string
	SSLMode string
}

var Postgres *PostgresConfig

func (p PostgresConfig) Init() {
	Postgres = &PostgresConfig{
		Host:    viper.GetString("postgres.host"),
		Port:    viper.GetString("postgres.port"),
		User:    viper.GetString("postgres.user"),
		Name:    viper.GetString("postgres.name"),
		Pass:    viper.GetString("postgres.pass"),
		SSLMode: viper.GetString("postgres.ssl_mode"),
	}
}
