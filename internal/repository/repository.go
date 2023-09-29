package repository

import (
	"url_shortener/internal/repository/link"
	"url_shortener/pkg/db"
	"url_shortener/pkg/logger"
)

type IRepository interface {
	GetLinkStorage() link.IRepository
}

type Repository struct {
	db     db.IDB
	logger logger.ILogger
}

func NewRepository(db db.IDB, logger logger.ILogger) *Repository {
	return &Repository{db: db, logger: logger}
}

func (r *Repository) GetLinkRepository() *link.Repository {
	return link.NewRepository(r.db, r.logger)
}
