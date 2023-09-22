package postgres

import (
	"fmt"
	"url_shortener/pkg/logger"
	"url_shortener/pkg/logger/zerolog"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DBConnector struct {
	db *sqlx.DB
}

func (c *DBConnector) GetConnect() *sqlx.DB {
	return c.db
}

func (c *DBConnector) CloseConnect() error {
	return c.db.Close()
}

func NewDBConnector(cfg Config) (*DBConnector, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Database)
	var log logger.ILogger = zerolog.NewLogger()
	log.Debug(dsn)
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, err
	}
	return &DBConnector{db: db}, nil
}
