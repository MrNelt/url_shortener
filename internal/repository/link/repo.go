package link

import (
	"context"
	"database/sql"
	errorsApi "url_shortener/internal/errors_api"
	"url_shortener/internal/models"
	dbConn "url_shortener/pkg/db_conn"
	"url_shortener/pkg/logger"
)

type IRepository interface {
	Save(link models.Link) (uint, error)
	SelectByLink(link string) (models.Link, error)
	SelectByID(id uint) (models.Link, error)
	SelectBySuffix(suffix string) (models.Link, error)
	DeleteByID(id uint) error
	IncrementClicksBySuffix(suffix string) error
}

type Repository struct {
	db     dbConn.IDBConnector
	logger logger.ILogger
}

func NewRepository(db dbConn.IDBConnector, logger logger.ILogger) *Repository {
	return &Repository{db: db, logger: logger}
}

func (r *Repository) Save(link models.Link) (uint, error) {
	client := r.db.GetConnect()
	res, err := client.Query(saveRequest, link.ShortSuffix, link.Url)
	if err != nil {
		r.logger.Error(err.Error())
		return 0, err
	}
	var ID uint
	err = res.Scan(&ID)
	if err != nil {
		r.logger.Error(err.Error())
		return 0, err
	}
	return ID, err
}

func (r *Repository) SelectBySuffix(shortSuffix string) (*models.Link, error) {
	client := r.db.GetConnect()
	row, err := client.Query(selectBySuffixRequest, shortSuffix)
	if err != nil {
		r.logger.Error(err.Error())
		return nil, err
	}
	var link models.Link
	err = row.Scan(&link.ID, &link.ShortSuffix, &link.Url, &link.ID, &link.Clicks, &link.ExpirationDate)
	switch {
	case err == sql.ErrNoRows:
		r.logger.Error(err.Error())
		return nil, errorsApi.ErrSuffixNotFound
	case err != nil:
		r.logger.Error(err.Error())
		return nil, err
	}

	return &link, nil
}

func (r *Repository) SelectByLink(longLink string) (*models.Link, error) {
	client := r.db.GetConnect()
	row := client.QueryRow(selectByLinkRequest, longLink)

	var link models.Link
	err := row.Scan(&link.ShortSuffix, &link.Url, &link.Clicks, &link.ExpirationDate)
	switch {
	case err == sql.ErrNoRows:
		r.logger.Error(err.Error())
		return nil, errorsApi.ErrLinkNotFound
	case err != nil:
		r.logger.Error(err.Error())
		return nil, err
	}

	return &link, nil
}

func (r *Repository) SelectByID(ID uint) (*models.Link, error) {
	client := r.db.GetConnect()
	row := client.QueryRow(selectByIDRequest, ID)

	var link models.Link
	err := row.Scan(&link.ID, &link.ShortSuffix, &link.Url, &link.Clicks, &link.ExpirationDate)
	switch {
	case err == sql.ErrNoRows:
		r.logger.Error(err.Error())
		return nil, errorsApi.ErrIDNotFound
	case err != nil:
		r.logger.Error(err.Error())
		return nil, err
	}

	return &link, nil
}

func (r *Repository) DeleteID(ID uint) error {
	client := r.db.GetConnect()
	res, err := client.Exec(deleteByIDRequest, ID)
	if err != nil {
		r.logger.Error(err.Error())
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		r.logger.Error(err.Error())
		return err
	}
	if rowsAffected == 0 {
		return errorsApi.ErrIDNotFound
	}

	return nil
}

func (r *Repository) IncrementClicksBySuffix(ctx context.Context, shortSuffix string) error {
	client := r.db.GetConnect()
	_, err := client.ExecContext(ctx, incrementClicksBySuffixRequest, shortSuffix)
	if err != nil {
		r.logger.Error(err.Error())
	}
	return err
}
