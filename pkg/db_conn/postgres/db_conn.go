package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type DBConnector struct {
	db *sqlx.DB
}

func (c *DBConnector) GetConnect() *sqlx.DB {
	return c.db
}

func NewDBConnector(cfg Config) (*DBConnector, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Database)
	db, err := sqlx.Connect(cfg.Database, dsn)
	if err != nil {
		return nil, err
	}
	return &DBConnector{db: db}, nil
}
