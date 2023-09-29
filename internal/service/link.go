package service

import (
	"errors"
	"url_shortener/internal/dtos"
	errorsApi "url_shortener/internal/errors"
	"url_shortener/internal/models"
	"url_shortener/internal/repository/link"
	"url_shortener/pkg/logger"
)

type IService interface {
	GetUrlForRedirect(suffix string) (string, error)
	GetLinkInfoFromID(id string) (dtos.LinkInfoDto, error)
	SaveLink(link dtos.LinkDTO) (string, error)
}

type Service struct {
	logger logger.ILogger
	repo   link.IRepository
}

func NewService(logger logger.ILogger, repo link.IRepository) *Service {
	return &Service{logger: logger, repo: repo}
}

func (s *Service) SaveLink(link dtos.LinkDTO) (string, error) {
	var linkDB models.Link
	_, err := s.repo.SelectBySuffix(link.ShortSuffix)
	switch {
	case err == nil:
		s.logger.Error(errorsApi.ErrSuffixAlreadyExists.Error())
		return "", errorsApi.ErrSuffixAlreadyExists
	case !errors.Is(err, errorsApi.ErrSuffixNotFound):
		s.logger.Error(err.Error())
		return "", err
	}
	expirationDate, err := _TTLDTOToDate(link.TTLCount, link.TTLUnit)
	if err != nil {
		return "", err
	}
	linkDB.ShortSuffix = link.ShortSuffix
	linkDB.Url = link.Url
	linkDB.ExpirationDate = expirationDate
	return s.repo.Save(linkDB)
}

func (s *Service) GetLinkInfoFromID(id string) (dtos.LinkInfoDto, error) {
	linkDB, err := s.repo.SelectByID(id)
	if err != nil {
		s.logger.Error(err.Error())
		return dtos.LinkInfoDto{}, err
	}
	var linkDTO dtos.LinkInfoDto
	linkDTO.ShortSuffix = linkDB.ShortSuffix
	linkDTO.ID = linkDB.ID
	linkDTO.Url = linkDB.Url
	linkDTO.Clicks = linkDB.Clicks
	linkDTO.ExpirationDate = linkDB.ExpirationDate
	return linkDTO, nil
}

func (s *Service) GetUrlForRedirect(suffix string) (string, error) {
	linkDB, err := s.repo.SelectBySuffix(suffix)
	if err != nil {
		s.logger.Error(err.Error())
		return "", err
	}
	err = s.repo.IncrementClicksBySuffix(linkDB.ShortSuffix)
	if err != nil {
		s.logger.Error(err.Error())
		return "", err
	}
	return linkDB.Url, nil
}
