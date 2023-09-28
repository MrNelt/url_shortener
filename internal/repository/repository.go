package repository

import (
	"url_shortener/internal/repository/link"
	dbConn "url_shortener/pkg/db_conn"
	"url_shortener/pkg/logger"
)

type IRepository interface {
	GetLinkStorage() link.IRepository
}

type Repository struct {
	db     dbConn.IDBConnector
	logger logger.ILogger
}

func NewRepository(db dbConn.IDBConnector, logger logger.ILogger) *Repository {
	return &Repository{db: db, logger: logger}
}

func (r *Repository) GetLinkRepository() *link.Repository {
	return link.NewRepository(r.db, r.logger)
}
