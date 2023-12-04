package in_memory

import (
	"context"
	"errors"
	"fmt"
	"github.com/IvanKonoplich/shortened_links_service/internal/infrastructure/repository"
)

const mapStartSize int64 = 100

type InMemoryStorage struct {
	storage map[string]string
}

func NewInMemoryStorage() *InMemoryStorage {
	return &InMemoryStorage{
		storage: make(map[string]string, mapStartSize),
	}
}

func (d *InMemoryStorage) SaveLink(_ context.Context, originalLink string) (shortenedLink string, err error) {
	if len(originalLink) == 0 {
		return "", errors.New("передана пустая оригинальная ссылка")
	}

	for {
		shortLink := repository.GenerateRandomLink()
		if _, ok := d.storage[shortLink]; ok {
			continue
		}
		d.storage[shortLink] = originalLink

		return shortLink, nil
	}
}

func (d *InMemoryStorage) GetLink(_ context.Context, shortenedLink string) (originalLink string, err error) {
	if len(shortenedLink) == 0 {
		return "", errors.New("передана пустая сокращенная ссылка")
	}

	originalLink, ok := d.storage[shortenedLink]
	if !ok {
		return "", fmt.Errorf("для сокращенной ссылки %s нет оригинальной ссылки", shortenedLink)
	}
	return originalLink, err
}
