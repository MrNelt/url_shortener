package config

import "url_shortener/pkg/db/postgres"

type Config struct {
	Host     string `config:"HOST" yaml:"host"`
	Port     string `config:"PORT" yaml:"port"`
	Database postgres.Config
}
