package link

import (
	"database/sql"
	"time"
	errorsApi "url_shortener/internal/errors_api"
	"url_shortener/internal/models"
	"url_shortener/pkg/db"
	"url_shortener/pkg/logger"
)

type IRepository interface {
	Save(link models.Link) (string, error)
	SelectByLink(link string) (models.Link, error)
	SelectByID(id string) (models.Link, error)
	SelectBySuffix(suffix string) (models.Link, error)
	DeleteByID(id string) error
	IncrementClicksBySuffix(suffix string) error
}

type Repository struct {
	db     db.IDB
	logger logger.ILogger
}

func NewRepository(db db.IDB, logger logger.ILogger) *Repository {
	return &Repository{db: db, logger: logger}
}

func (r *Repository) Save(link models.Link) (string, error) {
	client := r.db.GetConnect()
	res := client.QueryRow(saveRequest, link.ShortSuffix, link.Url, link.ExpirationDate)
	var ID string
	err := res.Scan(&ID)
	if err != nil {
		r.logger.Error(err.Error())
		return "", err
	}
	return ID, err
}

func (r *Repository) SelectBySuffix(shortSuffix string) (models.Link, error) {
	client := r.db.GetConnect()
	row := client.QueryRow(selectBySuffixRequest, shortSuffix, time.Now())
	var link models.Link
	err := row.Scan(&link.ID, &link.ShortSuffix, &link.Url, &link.Clicks, &link.ExpirationDate)
	switch {
	case err == sql.ErrNoRows:
		r.logger.Error(err.Error())
		return models.Link{}, errorsApi.ErrSuffixNotFound
	case err != nil:
		r.logger.Error(err.Error())
		return models.Link{}, err
	}

	return link, nil
}

func (r *Repository) SelectByLink(longLink string) (models.Link, error) {
	client := r.db.GetConnect()
	row := client.QueryRow(selectByLinkRequest, longLink, time.Now())

	var link models.Link
	err := row.Scan(&link.ID, &link.ShortSuffix, &link.Url, &link.Clicks, &link.ExpirationDate)
	switch {
	case err == sql.ErrNoRows:
		r.logger.Error(err.Error())
		return models.Link{}, errorsApi.ErrLinkNotFound
	case err != nil:
		r.logger.Error(err.Error())
		return models.Link{}, err
	}

	return link, nil
}

func (r *Repository) SelectByID(ID string) (models.Link, error) {
	client := r.db.GetConnect()
	row := client.QueryRow(selectByIDRequest, ID)

	var link models.Link
	err := row.Scan(&link.ID, &link.ShortSuffix, &link.Url, &link.Clicks, &link.ExpirationDate)
	switch {
	case err == sql.ErrNoRows:
		r.logger.Error(err.Error())
		return models.Link{}, errorsApi.ErrIDNotFound
	case err != nil:
		r.logger.Error(err.Error())
		return models.Link{}, err
	}

	return link, nil
}

func (r *Repository) DeleteByID(ID string) error {
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

func (r *Repository) IncrementClicksBySuffix(shortSuffix string) error {
	client := r.db.GetConnect()
	_, err := client.Exec(incrementClicksBySuffixRequest, shortSuffix, time.Now())
	if err != nil {
		r.logger.Error(err.Error())
	}
	return err
}
