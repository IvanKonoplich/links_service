package service

import (
	"context"
	"github.com/IvanKonoplich/shortened_links_service/pkg/shortened_links"
	"github.com/sirupsen/logrus"
)

//go:generate mockery --name Storage
type Storage interface {
	SaveLink(ctx context.Context, originalLink string) (shortenedLink string, err error)
	GetLink(ctx context.Context, shortenedLink string) (originalLink string, err error)
}

type Service struct {
	storage Storage
	shortened_links.UnimplementedShortenedLinksServer
}

func New(storage Storage) *Service {
	return &Service{
		storage: storage,
	}
}

func (s *Service) SaveLink(ctx context.Context, link *shortened_links.OriginalLink) (*shortened_links.ShortenedLink, error) {

	logrus.Infof("Новый запрос SaveLink ссылка: %s", link.GetMessage())

	shortLink, err := s.storage.SaveLink(ctx, link.GetMessage())
	if err != nil {
		logrus.Errorf("Ошибка во время выполнения запроса SaveLink: %s", err.Error())
		return nil, err
	}

	logrus.Infof("Запрос SaveLink выполнен: оригинальная ссылка: %s, сокращенная ссылка %s", link.GetMessage(), shortLink)
	return &shortened_links.ShortenedLink{Message: shortLink}, err
}

func (s *Service) GetLink(ctx context.Context, link *shortened_links.ShortenedLink) (*shortened_links.OriginalLink, error) {
	logrus.Infof("Новый запрос GetLink ссылка: %s", link.GetMessage())

	originalLink, err := s.storage.GetLink(ctx, link.GetMessage())
	if err != nil {
		logrus.Errorf("Ошибка во время выполнения запроса GetLink: %s", err.Error())
	}

	logrus.Infof("Запрос GetLink выполнен: сокращенная ссылка: %s, оригинальная ссылка %s", link.GetMessage(), originalLink)
	return &shortened_links.OriginalLink{Message: originalLink}, err
}
