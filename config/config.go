package config

import "url_shortener/pkg/db_conn/postgres"

type Config struct {
	Host string `config:"HOST" yaml:"host"`
	Port string `config:"PORT" yaml:"port"`
	postgres.Config
}
